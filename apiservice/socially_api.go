//社交服务

package apiservice

import (
	"fmt"
	"time"

	"app-server/define"
	"app-server/logic"
	"app-server/models"
	"app-server/pkg/mongodb"
	"app-server/pkg/redis"

	"gopkg.in/mgo.v2/bson"
)

// 从数据库中查询融云token，有就不需要去融云serve获取token
// flag = true表示重新去融云服务器获取token
func FindRcToken(userId, flag string) (string, error) {
	user, err := logic.FindUserDataById(userId)
	if err != nil {
		return "", err
	}

	if flag == "0" && len(user.RcToken) > 0 {
		return user.RcToken, nil
	}

	rcToken, getErr := logic.GetRcTokenFromRcSrv(userId, user.NickName, user.Portrait)
	if getErr != nil {
		return "", getErr
	}

	fmt.Println("融云token= ", rcToken)
	return rcToken.Token, nil
}

// 用户获取好友列表
func GetFriendList(userId string) (*models.FriendList, error) {
	var friends models.FriendList
	id := bson.ObjectIdHex(userId)
	selector := bson.M{"friends": 1, "_id": 0}
	err := mongodb.SelectById(define.AccountCollection, id, selector, &friends)

	return &friends, err
}

// 用户获取好友列表
func GetFriendInfo(userId, targetId string, updateTime int64) (*models.S2C_UserData, error) {
	// 是否是黑名单
	// TODO

	// 更新时间判断

	info, err := GetUserInfo(targetId, updateTime)
	if info != nil {
		info.Money = 0 //隐藏好用的账户余额和积分
		info.Point = 0
	}

	return info, err
}

// 通过手机号或者红包号添加好友
func AddFriendByAccount(userId, targetAccount string) uint32 {
	targetId := logic.FindUserIdByAccount(targetAccount)

	if len(targetId) == 0 {
		return 10002
	}

	return AddFriendById(userId, targetId)
}

// 通过用户id添加好友（适用于手机通讯录）
func AddFriendById(fromId, targetId string) uint32 {
	if logic.Isfriends(fromId, targetId) {
		return 10406
	}

	addKey := define.AddFriendPrefix + fromId + "+" + targetId
	r, _ := redis.Do("GET", addKey) //先检查是否有添加记录或者请求是否过期
	if r != nil {                   // 有记录
		return 10402
	}

	// 调用融云sdk
	err := logic.SendAddFriendReqToRcServer(fromId, targetId)
	if err != nil {
		return 10401
	}

	expire := time.Now().Add(time.Duration(72) * time.Hour).Unix() // 3天有效期
	redis.Do("SETEX", addKey, expire, 1)

	return 0
}

// 用户B同意A的加好友请求
func AgreeFriend(userIdB, userIdA string) uint32 {
	if logic.Isfriends(userIdA, userIdB) { // 已经是好友
		return 10406
	}

	// 从添加请求列表找到记录
	addKey := define.AddFriendPrefix + userIdA + "+" + userIdB
	r, err := redis.Do("DEL", addKey) //先检查是否有添加记录或者请求是否过期
	if err != nil || r.(int64) == 0 {
		return 10404
	}

	userA, userAErr := logic.GetUserData(userIdA)
	userB, userBErr := logic.GetUserData(userIdB)
	if userAErr != nil || userBErr != nil {
		return 10405
	}

	logic.AddFriendToDB(userIdA, userIdB, userB.NickName, userB.Portrait) // A 添加 B 为好友
	logic.AddFriendToDB(userIdB, userIdA, userA.NickName, userA.Portrait) // B 添加 A 为好友
	logic.SendAgreeNotifyToRcServer(userIdA, userIdB, userB.NickName)     // 通知A，好友添加成功
	logic.SendAgreeNotifyToRcServer(userIdB, userIdA, userA.NickName)     // 通知B，好友添加成功

	return 0
}

// B拒绝A的加好友请求
func RefuseFriend(userIdB, userIdA string) {
	addKey := define.AddFriendPrefix + userIdA + "+" + userIdB
	redis.Do("DEL", addKey)
}

// 删除好友
func RemoveFriend(userId, friendId string) {
	logic.UpdateUserById(userId, bson.M{"$pull": bson.M{"friends": bson.M{"userid": friendId}}})
	logic.UpdateUserById(friendId, bson.M{"$pull": bson.M{"friends": bson.M{"userid": userId}}})
}

// 加入到黑名单中
func AddToBlackList(userIdB, targetId string) {
	// 数据库处理
	// 本地服务器维护黑名单

	// 发消息个融云
	logic.SendAddBlacklistReqToRcServer(userIdB, targetId)
}
