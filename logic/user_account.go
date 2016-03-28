// 用户账号相关

package logic

import (
	"errors"
	"time"

	"app-server/define"
	"app-server/models"
	"app-server/pkg/mongodb"
	"app-server/pkg/utils"

	"gopkg.in/mgo.v2/bson"
)

func insertUserToDb(newUser *models.User) (error, *models.TokenInfo, *models.TokenInfo) {
	userId := newUser.Id_.Hex()
	accessToken, refreshToken := genTokens(userId)
	if accessToken == nil && refreshToken == nil { // 保存最新的token
		return errors.New("[Error] gen tokens fail"), nil, nil
	}

	newUser.AccessToken = accessToken.Token
	newUser.RefreshToken = refreshToken.Token
	err := mongodb.Insert(define.AccountCollection, newUser)
	if err != nil {
		return err, nil, nil
	}

	// 加入到缓存中
	updateUserCache(userId, newUser)
	return nil, accessToken, refreshToken
}

// 创建新用户
func CreateNewUser(phonenum, password string) (error, *models.TokenInfo, *models.TokenInfo) {
	salt := utils.RandUUID(32)
	hashpw := utils.Md5(password + salt)

	m := new(models.User)
	m.Id_ = bson.NewObjectId()
	m.Phone = phonenum
	m.Password = hashpw // hash(hasn(明文密码)+盐)
	m.Salt = salt
	m.NickName = utils.GenNickname() // 随机生成昵称
	m.UpdateTime = time.Now().Unix()
	m.Outcome = map[string][]int{"2015": []int{1, 2}}

	return insertUserToDb(m)
}

// 通过微信注册用户
func CreateUserByWeixin(openId string, weixinUser *models.WeixinUserInfo) (error, *models.TokenInfo, *models.TokenInfo) {
	m := new(models.User) // 注册账号
	m.Id_ = bson.NewObjectId()
	m.WeixinOpenId = openId
	m.NickName = weixinUser.NickName
	m.UpdateTime = time.Now().Unix()
	m.Portrait = weixinUser.Portrait
	m.Sex = uint8(weixinUser.Sex)
	if m.Sex == 2 { // 纠正微信性别定义
		m.Sex = 0
	}

	return insertUserToDb(m)
}

// 通过QQ注册用户
func CreateUserByQQ(openId string, qqUerInfo *models.QQUserInfo) (error, *models.TokenInfo, *models.TokenInfo) {
	m := new(models.User) // 注册账号
	m.Id_ = bson.NewObjectId()
	m.QQOpenId = openId
	m.NickName = qqUerInfo.NickName
	m.UpdateTime = time.Now().Unix()
	m.Portrait = qqUerInfo.Portrait
	m.Sex = 0
	if qqUerInfo.Sex == "男" {
		m.Sex = 1
	}

	return insertUserToDb(m)
}

func CreateUserByWeibo(uid string, weiboUerInfo *models.WeiboUserInfo) (error, *models.TokenInfo, *models.TokenInfo) {
	m := new(models.User) // 注册账号
	m.Id_ = bson.NewObjectId()
	m.WeiboOpenId = uid
	m.NickName = weiboUerInfo.NickName
	m.UpdateTime = time.Now().Unix()
	m.Portrait = weiboUerInfo.Portrait
	m.Sex = 0
	if weiboUerInfo.Sex == "m" {
		m.Sex = 1
	}

	return insertUserToDb(m)
}

// 检查第三方是否注册过
func CheckOauth(openId string, platform int) bool {
	var query bson.M
	switch platform {
	case define.E_Weixin:
		query = bson.M{"weixin_openid": openId}
	case define.E_Weibo:
		query = bson.M{"weibo_openid": openId}
	case define.E_QQ:
		query = bson.M{"qq_openid": openId}
	default:
		return true
	}

	return mongodb.Exists(define.AccountCollection, query)
}

// 设置第三方账户
func BindOauth(userId, openId string, platform int) error {
	var update bson.M
	var key string
	switch platform {
	case define.E_Weixin:
		key = define.EUser_WeixinOpenId
		update = bson.M{"$set": bson.M{"weixin_openid": openId}}
	case define.E_Weibo:
		key = define.EUser_WeiboOpenId
		update = bson.M{"$set": bson.M{"weibo_openid": openId}}
	case define.E_QQ:
		key = define.EUser_QQOpenId
		update = bson.M{"$set": bson.M{"qq_openid": openId}}
	default:
		return errors.New("[Error] invaild oauth2")
	}

	err := UpdateUserById(userId, update)
	if err == nil { // 更新缓存内的字段
		go UpdateUserCacheField(userId, key, openId)
	}

	return err
}

func BindPhone(userId, phonenum, password string) error {
	salt := utils.RandUUID(32)
	hashpwd := utils.Md5(password + salt)
	update := bson.M{"$set": bson.M{"phone": phonenum, "password": hashpwd, "salt": salt}}
	err := UpdateUserById(userId, update)
	if err == nil {
		go UpdateUserCacheField(userId, define.EUser_Phone, phonenum)
	}
	return err
}

func ResetPhone(userId, phonenum string) error {
	update := bson.M{"$set": bson.M{"phone": phonenum}}
	err := UpdateUserById(userId, update)
	if err == nil {
		go UpdateUserCacheField(userId, define.EUser_Phone, phonenum)
	}
	return err
}

// 解绑第三方账号
func UnBindOauth(userId string, platform int) error {
	var update bson.M
	var key string

	switch platform {
	case define.E_Weixin:
		key = define.EUser_WeixinOpenId
		update = bson.M{"$set": bson.M{"weixin_openid": ""}}
	case define.E_Weibo:
		key = define.EUser_WeiboOpenId
		update = bson.M{"$set": bson.M{"weibo_openid": ""}}
	case define.E_QQ:
		key = define.EUser_QQOpenId
		update = bson.M{"$set": bson.M{"qq_openid": ""}}
	default:
		return errors.New("[Error] invaild oauth2")
	}

	err := UpdateUserById(userId, update)
	if err == nil { // 更新缓存内的字段
		go UpdateUserCacheField(userId, key, "")
	}

	return err
}

// 重置密码并且重新生成盐
func ResetPassword(userData *models.User, pwd string) (error, *models.TokenInfo, *models.TokenInfo) {
	userId := userData.Id_.Hex()
	tokenErr, accessToken, refreshToken := RefreshTokens(userId, userData.AccessToken, userData.RefreshToken, define.ERefreshToken_Forcibly)
	if tokenErr == nil { // 保存最新的token
		salt := utils.RandUUID(32)
		hashpw := utils.Md5(pwd + salt)

		field := bson.M{"salt": salt,
			"password": hashpw,
		}
		update := bson.M{"$set": field}
		err := UpdateUserById(userId, update)
		if err != nil {
			return err, nil, nil
		}
	}
	return tokenErr, accessToken, refreshToken
}
