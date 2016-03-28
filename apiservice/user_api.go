package apiservice

import (
	//"errors"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"app-server/define"
	"app-server/logic"
	"app-server/models"
	"app-server/pkg/redis"
	"app-server/pkg/token"
	"app-server/pkg/utils"

	"gopkg.in/mgo.v2/bson"
)

var UserList map[string]*models.User

/////////////////////////////////////////////////////////
// public方法
// 用户登录，验证通过返回token
func Login(account, password string) (uint32, *models.TokenInfo, *models.TokenInfo) {
	user := logic.FindUserByAccount(account)
	if user == nil {
		return 10002, nil, nil
	}

	// 验证密码
	pw := utils.Md5(password + user.Salt)
	if pw != user.Password {
		return 10003, nil, nil
	}

	userId := user.Id_.Hex()
	err, accessToken, refreshToken := logic.RefreshTokens(userId, user.AccessToken, user.RefreshToken, define.ERefreshToken_Login)
	if err != nil {
		return 10001, nil, nil
	}

	return 0, accessToken, refreshToken
}

// 退出
func Logout(req *http.Request, userId string) error {
	refreshToken := logic.FindRefreshToken(userId)
	authBackend := token.GetJwtBackend()

	return authBackend.OnLogout(req, refreshToken)
}

// 注册账号，密码加盐
// hashpass = hash(hash(password)+salt)
func Register(phonenum, password string) (uint32, *models.TokenInfo, *models.TokenInfo) {
	// 检查账号的合法性
	if !utils.IsAllowableAccout(phonenum) {
		return 10007, nil, nil
	}

	// 检查账号是否存在
	userId := logic.FindUserIdByAccount(phonenum)
	if len(userId) != 0 {
		return 10008, nil, nil
	}

	// 非明文密码
	err, accessToken, refreshToken := logic.CreateNewUser(phonenum, password)
	if err != nil {
		fmt.Println("err = ", err)
		return 10009, nil, nil
	}

	return 0, accessToken, refreshToken
}

// 拉取用户信息
func GetUserInfo(userId string, updateTime int64) (*models.S2C_UserData, error) {
	userData, err := logic.GetUserData(userId)
	if err != nil {
		return nil, err
	}

	// 判断是否需要更新
	// if updateTime > userData.UpdateTime {
	// 	return nil, errors.New("userdata no change")
	// }

	oauth := [5]int{}
	switch {
	case len(userData.WeixinOpenId) > 0:
		oauth[0] = 1
	case len(userData.WeiboOpenId) > 0:
		oauth[1] = 1
	case len(userData.QQOpenId) > 0:
		oauth[2] = 1
	}
	s2cUserData := models.S2C_UserData{UserPublic: userData.UserPublic, UserId: userId, Oauth: oauth}
	return &s2cUserData, err
}

// 刷新token，用旧的refresh token获取新的access token 和新的refresh token
func GenerateNewToken(tokenStr string) (uint32, *models.TokenInfo, *models.TokenInfo) {
	userId := token.GetJwtBackend().GetUserIdFromTokenStr(tokenStr) // 从refresh token中获取userid
	curToken := logic.FindAccessToken(userId)
	if len(curToken) == 0 {
		return 10002, nil, nil
	}

	err, accessToken, refreshToken := logic.RefreshTokens(userId, curToken, tokenStr, define.ERefreshToken_Refresh)
	if err != nil {
		return 10004, nil, nil
	}

	return 0, accessToken, refreshToken
}

// 重置密码
func ResetPassword(userId, password, flag, code string) (uint32, *models.TokenInfo, *models.TokenInfo) {
	userData, err := logic.FindUserDataById(userId)
	if err != nil {
		return 10002, nil, nil
	}

	if flag == "1" { // 找回密码
		if len(userData.Phone) == 0 {
			return 10029, nil, nil
		}

		// 检查手机验证码
		err = logic.VerifySmsCode(userData.Phone, "86", code)
		if err != nil {
			fmt.Println("手机验证码 err = ", err)
			return 10030, nil, nil
		}
	}

	err, accessToken, refreshToken := logic.ResetPassword(userData, password)
	if err != nil {
		return 10028, nil, nil
	}
	return 0, accessToken, refreshToken
}

// 重置密码
func VerifyPassword(userId, password string) uint32 {
	userData, err := logic.FindUserDataById(userId)
	if err != nil {
		return 10002
	}

	// 验证密码
	pw := utils.Md5(password + userData.Salt)
	if pw != userData.Password {
		return 10003
	}

	return 0
}

// 设置红包账号
func BindAccount(userId, account string) uint32 {
	if !utils.IsAccount(account) {
		return 10012
	}

	user, findErr := logic.GetUserData(userId)
	if findErr != nil {
		return 10002
	}

	if len(user.Account) != 0 { //已经绑定了红包号
		return 10011
	}

	// 红包号是否被注册
	usedUserId := logic.FindUserIdByAccount(user.Account)
	if len(usedUserId) != 0 {
		return 10013
	}

	err := UpdatePersonalInfo(userId, define.EUser_Account, account)
	if err != nil {
		return 10014
	}

	return 0
}

// 绑定手机号(第一次绑定手机号并设置密码)
func BindPhoneNum(userId, phoneNum, password string) uint32 {
	if !utils.IsPhoneNumber(phoneNum) { // 手机号码格式错误
		return 10007
	}

	user, err := logic.GetUserData(userId)
	if err != nil {
		return 10002
	}

	if len(user.Phone) > 0 { // 账号已经绑定了手机号
		return 10006
	}

	logic.BindPhone(userId, phoneNum, password)
	return 0
}

// 更换手机号码
func ResetPhoneNum(userId, phoneNum, password, code string) uint32 {
	if !utils.IsPhoneNumber(phoneNum) { // 手机号码格式错误
		return 10007
	}

	// 检查该手机号是否存在
	uid := logic.FindUserIdByAccount(phoneNum)
	if len(uid) != 0 {
		return 10008
	}

	user, err := logic.FindUserDataById(userId)
	if err != nil {
		return 10002
	}

	// 验证密码
	pw := utils.Md5(password + user.Salt)
	if pw != user.Password {
		return 10003
	}

	err = logic.VerifySmsCode(phoneNum, "86", code)
	if err != nil {
		fmt.Println("手机验证码 err = ", err)
		return 10030
	}

	logic.ResetPhone(userId, phoneNum)
	return 0
}

// 更新个人资料(个性签名、地理位置、性别)
func UpdatePersonalInfo(userId, key string, value interface{}) error {
	update := bson.M{"$set": bson.M{key: value}}
	err := logic.UpdateUserById(userId, update)
	if err != nil {
		return err
	}

	var v interface{}
	switch value.(type) {
	case models.AreaInfo:
		b, _ := json.Marshal(value)
		v = string(b)
	default:
		v = value
	}

	go logic.UpdateUserCacheField(userId, key, v) // 更新缓存
	if key == define.EUser_Nickname {
		logic.UpdateNameToFriends(userId, v.(string)) // 更新好友头像和昵称
		logic.RefreshRcUser(userId, v.(string), "")   // 更新融云token
	}

	return nil
}

func GetRankListByType(userId, tpStr, page string) *models.S2C_RankList {
	tp, _ := strconv.Atoi(tpStr)
	index, _ := strconv.Atoi(page)

	var rankName string
	switch tp {
	case define.ERank_Income:
		rankName = define.IncomeRank
	case define.ERank_Friend:
		rankName = define.FriendRank
	case define.ERank_Level:
		rankName = define.LevelRank
	case define.ERank_Outcome:
		rankName = define.OutcomeRank
	default:
		return nil
	}

	// 自己的排名
	selfRank, err := redis.GetInt("ZRANK", rankName, userId)
	if err != nil {
		selfRank = 0
	} else {
		selfRank += 1
	}

	rankList := &models.S2C_RankList{}
	rankList.List = make([]models.RankItem, 0, 10)
	rankList.Self = selfRank
	rankList.Type = tp

	// rank列表
	start := index*10 - 1
	if index == 0 {
		start = 0
	}
	idList, _ := redis.GetStrings("ZRANGE", rankName, start, start+10, "WITHSCORES")
	if len(idList) >= 2 && len(idList)%2 == 0 {
		var temp int
		for i := 0; i < len(idList); i += 2 {
			uid := idList[i]
			score, _ := strconv.Atoi(idList[i+1])

			rank := models.RankItem{}
			rank.Rank = start + 1 + temp
			rank.Score = score
			rank.Id = uid
			rank.Name = "用户昵称"
			rankList.List = append(rankList.List, rank)
			temp++
		}
	}
	return rankList
}

func GetSystemNoticeList(userId string) *models.S2C_SysNoticeList {
	noticeList := &models.S2C_SysNoticeList{}
	list1 := logic.GetSystemNoticeList(userId)
	if len(list1) > 0 {
		noticeList.List = append(noticeList.List, list1...)
	}

	list2 := logic.GetRedpacketNotice(userId)
	if len(list2) > 0 {
		noticeList.List = append(noticeList.List, list2...)
	}

	return noticeList
}
