package logic

import (
	"errors"
	//"fmt"

	"app-server/define"
	"app-server/models"
	"app-server/pkg/mongodb"
	"app-server/pkg/token"

	"gopkg.in/mgo.v2/bson"
)

const (
	RefreshTokenDiff = 10 * 86400 // 如果refresh token的有效期低于10天，则生成新的refresh token，旧的token也有效
)

// 查询用户的融云Token
func FindUserRcToken(userId string) string {
	var rcToken struct {
		RcToken string `bson:"rctoken"`
	}

	fields := bson.M{"rctoken": 1}
	id := bson.ObjectIdHex(userId)
	mongodb.SelectById(define.AccountCollection, id, fields, &rcToken)

	return rcToken.RcToken
}

// 根据用户id查询refresh token
func FindRefreshToken(userId string) string {
	var refreshToken struct {
		Refresh string `bson:"refreshtoken"`
	}

	fields := bson.M{"refreshtoken": 1}
	id := bson.ObjectIdHex(userId)
	mongodb.SelectById(define.AccountCollection, id, fields, &refreshToken)
	return refreshToken.Refresh
}

// 根据用户id查询access token
func FindAccessToken(userId string) string {
	var token struct {
		Access string `bson:"accesstoken"`
	}

	fields := bson.M{"accesstoken": 1}
	id := bson.ObjectIdHex(userId)
	mongodb.SelectById(define.AccountCollection, id, fields, &token)
	return token.Access
}

func genTokens(userId string) (accessToken, refreshToken *models.TokenInfo) {
	authBackend := token.GetJwtBackend()
	accesstokenStr, accessExp := authBackend.GenerateToken(userId, token.AccessToken)
	refreshtokenStr, refreshExp := authBackend.GenerateToken(userId, token.RefreshToken)
	if len(accesstokenStr) == 0 || len(refreshtokenStr) == 0 {
		return
	}

	// access token
	accessToken = new(models.TokenInfo)
	accessToken.Token = accesstokenStr
	accessToken.Expiresin = accessExp

	// refresh token
	refreshToken = new(models.TokenInfo)
	refreshToken.Token = refreshtokenStr
	refreshToken.Expiresin = refreshExp

	return
}

// 登陆时，刷新token
func RefreshTokens(userId, accessStr, refreshStr string, oper int) (error, *models.TokenInfo, *models.TokenInfo) {
	jwtBackend := token.GetJwtBackend()
	refresh, err := jwtBackend.ParseToken(refreshStr)
	if err != nil {
		return err, nil, nil
	}

	expire := jwtBackend.GetRemainingValidity(refresh) // refresh token的剩余有效期
	var accessToken, refreshToken models.TokenInfo
	var needrefresh bool = false
	switch oper {
	case define.ERefreshToken_Refresh: // 用refresh token刷新
		if !refresh.Valid {
			return errors.New("[Error]refresh token invalid"), nil, nil
		}
		if expire <= RefreshTokenDiff {
			needrefresh = true
		}
	case define.ERefreshToken_Login: // 登陆
		if !refresh.Valid || expire <= RefreshTokenDiff {
			needrefresh = true
		}
	case define.ERefreshToken_Forcibly: // 强制全部更新
		needrefresh = true
	default:
		return errors.New("[Error] invaild operate"), nil, nil
	}

	newAccessToken, newRefreshToken := genTokens(userId)
	if newAccessToken == nil || newRefreshToken == nil {
		return errors.New("[Error] refresh tokens fail"), nil, nil
	}
	jwtBackend.DropToken(accessStr) // 丢弃旧的access token
	accessToken = *newAccessToken
	if needrefresh {
		jwtBackend.DropToken(refreshStr) // 丢弃旧的refresh token
		refreshToken = *newRefreshToken
	} else {
		refreshToken.Token = refreshStr
		refreshToken.Expiresin = uint32(expire)
	}

	//fmt.Println("[Refresh token]---------->operate ", oper, accessToken, refreshToken)

	// 保存最新的token
	update := bson.M{
		"$set": bson.M{
			"accesstoken":  accessToken.Token,
			"refreshtoken": refreshToken.Token,
		},
	}
	err = UpdateUserById(userId, update)
	return err, &accessToken, &refreshToken
}
