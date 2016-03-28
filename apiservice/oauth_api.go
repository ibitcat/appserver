package apiservice

import (
	"errors"
	"fmt"
	"strconv"

	"app-server/define"
	"app-server/logic"
	"app-server/models"
	"app-server/oauth2/qq"
	"app-server/oauth2/weibo"
	"app-server/oauth2/weixin"
)

/////////////////////////////////////////////////////////
// 第三方登陆
func LoginByWeixin(openId, token string) (err error, access, refresh *models.TokenInfo) {
	err = weixin.CheckWeixinAuth(openId, token)
	if err != nil {
		return
	}

	user := logic.FindUserByOpenId(openId, define.E_Weixin) // 检查有没有注册账号
	if user == nil {
		// 查询微信用户信息
		weixinUerInfo := weixin.GetWeixinUserInfo(openId, token)
		fmt.Println("微信登陆 = ", weixinUerInfo, err)
		if weixinUerInfo == nil || weixinUerInfo.OpenId != openId {
			err = errors.New("[Error]weinxin user error")
			return
		}

		err, access, refresh = logic.CreateUserByWeixin(openId, weixinUerInfo)
		return
	} else {
		err, access, refresh = logic.RefreshTokens(user.Id_.Hex(), user.AccessToken, user.RefreshToken, define.ERefreshToken_Login)
		return
	}
}

func LoginByWeibo(token string) (err error, access, refresh *models.TokenInfo) {
	uid := weibo.CheckWeiboTokenInfo(token)
	if uid == 0 {
		err = errors.New("[Error] weibo token error")
		return
	}

	weiboUid := strconv.FormatInt(uid, 10)                   // 微博用户uid
	user := logic.FindUserByOpenId(weiboUid, define.E_Weibo) // 检查有没有注册账号
	if user == nil {
		// 查询微博用户信息
		weiboUerInfo := weibo.GetWeiboUserInfo(weiboUid, token)
		if weiboUerInfo == nil || weiboUerInfo.Uid != uid {
			err = errors.New("[Error] weibo userinfo error")
			return
		}

		err, access, refresh = logic.CreateUserByWeibo(weiboUid, weiboUerInfo)
		return
	} else {
		err, access, refresh = logic.RefreshTokens(user.Id_.Hex(), user.AccessToken, user.RefreshToken, define.ERefreshToken_Login)
		return
	}
}

func LoginByQQ(openId, openKey string) (err error, access, refresh *models.TokenInfo) {
	err = qq.IsLogin(openId, openKey)
	if err != nil {
		return
	}

	user := logic.FindUserByOpenId(openId, define.E_QQ) // 检查有没有注册账号
	if user == nil {
		// 查询QQ用户信息
		qqUerInfo := qq.GetQQUserInfo(openId, openKey)
		if qqUerInfo == nil {
			err = errors.New("[Error] qq userinfo error")
			return
		}

		err, access, refresh = logic.CreateUserByQQ(openId, qqUerInfo)
		return
	} else {
		err, access, refresh = logic.RefreshTokens(user.Id_.Hex(), user.AccessToken, user.RefreshToken, define.ERefreshToken_Login)
		return
	}
}

/////////////////////////////////////////////////////////
// 绑定第三方账号
// 绑定微信账号
func BindWeixinAccount(userId, openId, token string) uint32 {
	user, err := logic.GetUserData(userId)
	if err != nil {
		return 10002
	}

	if len(user.WeixinOpenId) > 0 { // 已经绑定了微信号
		return 10020
	}

	if logic.CheckOauth(openId, define.E_Weixin) { //该微信号已经绑定了账号
		return 10021
	}

	err = weixin.CheckWeixinAuth(openId, token)
	if err != nil { // 微信授权失败
		return 10016
	}
	logic.BindOauth(userId, openId, define.E_Weixin)
	return 0
}

// 绑定微博账号
func BindWeiboAccount(userId, token string) uint32 {
	user, err := logic.GetUserData(userId)
	if err != nil {
		return 10002
	}

	if len(user.WeiboOpenId) > 0 { // 已经绑定了微博号
		return 10022
	}

	uid := weibo.CheckWeiboTokenInfo(token)
	if uid == 0 { // 微博授权失败
		return 10023
	}
	weiboUid := strconv.FormatInt(uid, 10)          // 微博用户uid
	if logic.CheckOauth(weiboUid, define.E_Weibo) { //该微博号已经绑定了账号
		return 10018
	}

	logic.BindOauth(userId, weiboUid, define.E_Weibo)
	return 0
}

// 绑定QQ账号
func BindQQAccount(userId, openId, openKey string) uint32 {
	user, err := logic.GetUserData(userId)
	if err != nil {
		return 10002
	}

	if len(user.QQOpenId) > 0 { // 已经绑定了QQ号
		return 10024
	}

	if logic.CheckOauth(openId, define.E_QQ) { //该QQ已经绑定了账号
		return 10025
	}

	err = qq.IsLogin(openId, openKey)
	if err != nil { // qq授权失败
		return 10017
	}

	logic.BindOauth(userId, openId, define.E_QQ)
	return 0
}

/////////////////////////////////////////////////////////
// 解绑第三方登陆
func UnbindOauthAccount(userId, plat string) uint32 {
	userData, err := logic.GetUserData(userId)
	if err != nil {
		return 10004
	}

	acc := make([]int, 0, 5)
	if len(userData.WeixinOpenId) > 0 {
		acc = append(acc, 1)
	}
	if len(userData.WeiboOpenId) > 0 {
		acc = append(acc, 1)
	}
	if len(userData.QQOpenId) > 0 {
		acc = append(acc, 1)
	}
	if len(userData.Account) > 0 {
		acc = append(acc, 1)
	}
	if len(userData.Phone) > 0 {
		acc = append(acc, 1)
	}

	fmt.Println(acc)
	if len(acc) < 2 { //解绑至少需要有两个相关账号
		return 10026
	}

	platform, e := strconv.Atoi(plat)
	if e == nil {
		logic.UnBindOauth(userId, platform)
		return 0
	}
	return 10027
}
