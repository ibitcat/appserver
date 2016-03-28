package token

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"net/http"
	"time"

	"app-server/pkg/redis"
	"app-server/pkg/utils"

	"github.com/dgrijalva/jwt-go"
)

/*
访问token，有效期暂定1天 = 86400s
刷新token，有效期暂定1个月 = 30*86400s
*/

// token类型
const (
	AccessToken = iota
	RefreshToken
)

const (
	expireOffset = 3600 // token在redis中存储的有效期
	redisDBNum   = 1    // 使用redis第2个db来处理token
)

// jwt
type JwtBackend struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
	secret     []byte
}

// 生成token(使用hmac256,token比较短)
func (this *JwtBackend) GenerateToken(userUUID string, tokenType uint8) (string, uint32) {
	token := jwt.New(jwt.SigningMethodHS256)

	exp := time.Hour * time.Duration(24) // 1天后到期
	if tokenType == RefreshToken {
		exp = time.Hour * time.Duration(24*30*6) // 6个月后到期
	}

	jwtId := fmt.Sprintf("jwt:%s-%s-%d", userUUID, utils.RandUUID(10), tokenType)
	fmt.Println(jwtId)
	token.Claims["exp"] = time.Now().Add(exp).Unix() // jwt到期时间
	token.Claims["iat"] = time.Now().UnixNano()      // jwt创建时间
	token.Claims["jti"] = jwtId                      // jwt的唯一标示
	token.Claims["aud"] = userUUID                   // jwt的接收者
	token.Claims["sub"] = tokenType                  // jwt主题（这里赋值为token类型）

	tokenString, err := token.SignedString(this.secret)
	if err != nil {
		return "", 0
	}

	return tokenString, uint32(exp.Seconds())
}

// 解析jwt（传入string或者*http.Request）
func (this *JwtBackend) ParseToken(in interface{}) (*jwt.Token, error) {
	var jwtObj *jwt.Token
	var err error

	keyFunc := func(obj *jwt.Token) (interface{}, error) {
		switch obj.Method.(type) {
		case *jwt.SigningMethodRSA:
			return this.PublicKey, nil
		case *jwt.SigningMethodHMAC:
			return this.secret, nil
		default:
			return nil, fmt.Errorf("Unexpected signing method: %v", obj.Header["alg"])
		}
	}

	switch v := in.(type) {
	case *http.Request:
		jwtObj, err = jwt.ParseFromRequest(v, keyFunc)
	case string:
		jwtObj, err = jwt.Parse(v, keyFunc)
	default:
		jwtObj, err = nil, errors.New("parse jwt fail!") // 参数错误
	}

	return jwtObj, err
}

// 解析token并认证
func (this *JwtBackend) TokenAuthentication(req *http.Request) (error, *jwt.Token) {
	tokenObj, err := this.ParseToken(req)
	if err != nil {
		return err, nil
	}

	if tokenObj.Valid {
		if int(tokenObj.Claims["sub"].(float64)) == RefreshToken { // 不能用refreshtoken来访问api
			return errors.New("it's refresh token"), nil
		}

		if !this.isInBlacklist(tokenObj) { // 没有在黑名单中
			return nil, tokenObj
		}
	}

	return errors.New("access token not valid"), nil
}

// 登出时，把未过期的token入库到redis中
// accessToken和refreshToken一起失效（安全起见，要失效都要一起失效）
func (this *JwtBackend) OnLogout(req *http.Request, refreshTokenStr string) error {
	var accessToken, refreshToken *jwt.Token
	var err error

	accessToken, err = this.ParseToken(req)
	if err == nil && accessToken.Valid {
		key := accessToken.Claims["jti"].(string)
		expire := this.GetRemainingValidity(accessToken) + expireOffset
		redis.Do("SETEX", key, expire, accessToken.Raw)
	}

	refreshToken, err = this.ParseToken(refreshTokenStr)
	if err == nil && refreshToken.Valid {
		key := refreshToken.Claims["jti"].(string)
		expire := this.GetRemainingValidity(refreshToken) + expireOffset
		redis.Do("SETEX", key, expire, refreshTokenStr)
	}

	return err
}

// 从token中获取user uuid
func (this *JwtBackend) GetUserIdFromToken(tokenObj interface{}) string {
	if value, ok := tokenObj.(*jwt.Token); ok {
		userId, commaOk := value.Claims["aud"].(string)
		if commaOk {
			return userId
		}
	}

	return ""
}

// 从token string获取 userid
func (this *JwtBackend) GetUserIdFromTokenStr(tokenStr string) string {
	tokenObj, err := this.ParseToken(tokenStr)
	if err != nil || !tokenObj.Valid {
		return ""
	}

	userId := this.GetUserIdFromToken(tokenObj)
	if len(userId) == 0 || this.isInBlacklist(tokenObj) {
		return ""
	}

	return userId
}

// 丢弃token到redis中
func (this *JwtBackend) DropToken(tokenStr string) error {
	tokenObj, parseErr := this.ParseToken(tokenStr)
	if parseErr == nil {
		key := tokenObj.Claims["jti"].(string)
		expire := this.GetRemainingValidity(tokenObj) + expireOffset
		_, err := redis.Do("SETEX", key, expire, tokenStr)
		return err
	}

	return nil
}

// token的有效期
func (this *JwtBackend) GetRemainingValidity(tokenObj *jwt.Token) int64 {
	timestamp := tokenObj.Claims["exp"] // 到期时间
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int64(remainer.Seconds())
		}
	}

	return 0
}

/////////////////////////////////////////////////////////
// 是否在redis中存在未失效的token
func (this *JwtBackend) isInBlacklist(tokenObj *jwt.Token) bool {
	jwtId, commaOk := tokenObj.Claims["jti"].(string)
	if commaOk {
		redisToken, _ := redis.Do("GET", jwtId)
		if redisToken == nil {
			return false
		}
	}

	return true
}
