package logic

import (
	"encoding/json"
	"errors"
	"fmt"

	"app-server/pkg/sdk/rongcloud"

	"gopkg.in/mgo.v2/bson"
)

type RcTokenResponse struct {
	Code   int    `json:"code"`   // 返回码
	Token  string `json:"token"`  // 融云token
	UserId string `json:"userId"` // 用户id
}

// 从融云server获取新的token
func GetRcTokenFromRcSrv(userId, name, portrait string) (*RcTokenResponse, error) {
	data, rcErr := rongcloud.RcServer.UserGetToken(userId, name, portrait)
	if rcErr != nil {
		return nil, rcErr
	}

	var rcToken RcTokenResponse
	jsonErr := json.Unmarshal(data, &rcToken)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if rcToken.Code != 200 {
		return nil, errors.New("rcToken.Code error")
	}

	if rcToken.UserId != userId {
		return nil, errors.New("用户名不一致")
	}

	//  存储token到mongodb
	go UpdateUserById(userId, bson.M{"$set": bson.M{"rctoken": rcToken.Token}})
	return &rcToken, nil
}

// 发送加好友请求到融云服务器
func SendAddFriendReqToRcServer(fromId, targetId string) error {
	var response struct {
		Code int `json:"code"`
	}

	content := fmt.Sprintf(`{"operation":"op1","sourceUserId":"%s","targetUserId":"%s","message":"%s","extra":"附加信息"}`, fromId, targetId, "我要加你好友")
	returnData, returnError := rongcloud.RcServer.MessageSystemPublish(fromId, []string{targetId}, "RC:ContactNtf", content, "请求加为好友", "")

	if returnError != nil || len(returnData) == 0 {
		fmt.Println("发送单聊消息：测试失败。returnError:", returnError)
		return returnError
	} else {
		fmt.Println("发送单聊消息：测试通过。returnData:", string(returnData))
		json.Unmarshal(returnData, &response)
		if response.Code == 200 {
			return nil
		} else {
			return errors.New("response code error")
		}
	}

	return nil
}

// 发送成为好友的通知
func SendAgreeNotifyToRcServer(fromId, targetId, targetName string) {
	content := fmt.Sprintf(`{"message":"你与%s已经成为好友，开始聊天吧","extra":""}`, targetName)
	returnData, returnError := rongcloud.RcServer.MessageSystemPublish(fromId, []string{targetId}, "RC:InfoNtf", content, "添加好友成功", "")

	if returnError != nil || len(returnData) == 0 {
		fmt.Println("发送成为好友的通知：测试失败。returnError:", returnError)
	} else {
		fmt.Println("发送成为好友的通知：测试成功。returnError:", string(returnData))
	}
}

// 刷新融云用户信息
func RefreshRcUser(userId, name, portraitUri string) {
	returnData, returnError := rongcloud.RcServer.UserRefresh(userId, name, portraitUri)

	if returnError != nil || len(returnData) == 0 {
		fmt.Println("刷新融云用户信息失败 。returnError:", returnError)
	} else {
		fmt.Println("刷新融云用户信息成功。returnError:", string(returnData))
	}
}

// 向融云服务器发送拉黑请求
func SendAddBlacklistReqToRcServer(userId, blackUserId string) {
	returnData, returnError := rongcloud.RcServer.UserBlackAdd(userId, blackUserId)

	if returnError != nil || len(returnData) == 0 {
		fmt.Println("拉黑失败 。returnError:", returnError)
	} else {
		fmt.Println("拉黑成功。returnError:", string(returnData))
	}
}
