/*
玩家相关
分为三块：
1、和数据库交互的model；
2、和逻辑交互的model；
3、和客户端交互的model；
*/

package models

import (
	"gopkg.in/mgo.v2/bson"
)

/* *
* struct tag 说明：
* bson  tag 不为"-"，表示对mongodb的有效字段
* json  tag 不为"-"，表示对客户端的有效字段（对客户端开放）
* redis tag 不为"-"，表示对redis的有效字段（该字段缓存到redis中）
* */
/////////////////////////////////////////////////////////
// 用户的私密信息
type UserPrivate struct {
	Password     string `bson:"password" redis:"-"`                  //密码
	Salt         string `bson:"salt" redis:"-"`                      //盐
	AccessToken  string `bson:"accesstoken" redis:"-"`               //当前使用的access token
	RefreshToken string `bson:"refreshtoken" redis:"-"`              //当前使用的refresh token
	RcToken      string `bson:"rctoken" redis:"-"`                   //用户的融云token
	WeixinOpenId string `bson:"weixin_openid" redis:"weixin_openid"` //微信openid(对客户端隐藏)
	WeiboOpenId  string `bson:"weibo_openid" redis:"weibo_openid"`   //新浪微博openid(对客户端隐藏)
	QQOpenId     string `bson:"qq_openid" redis:"qq_openid"`         //腾讯QQ openid(对客户端隐藏)
	LoginCount   int    `bson:"login_count" redis:"login_count"`     //用户的登陆次数(累加)
	IsGm         bool   `bson:"is_gm" redis:"is_gm"`                 // 是否为gm
}

// 用户个人信息
type UserPublic struct {
	Phone      string   `bson:"phone"  json:"phone" redis:"phone"`                  //手机号码,暂只支持国内手机号码 "13877778888"
	Account    string   `bson:"account" json:"account" redis:"account"`             //红包账号,只能以英文字母开头且只能包含英文字母和数字
	UpdateTime int64    `bson:"update_time" json:"update_time" redis:"update_time"` //更新信息时间
	Portrait   string   `bson:"portrait" json:"portrait" redis:"portrait"`          //头像url（七牛云图片url）
	NickName   string   `bson:"nickname" json:"nickname" redis:"nickname"`          //昵称，可重复
	Sex        uint8    `bson:"sex" json:"sex" redis:"sex"`                         //性别（0=女，1=男）
	Area       AreaInfo `bson:"area" json:"area" redis:"area"`                      //地区
	Cert       uint8    `bson:"cert" json:"cert" redis:"cert"`                      //商户认证（0=未认证，1=审核中，2=已认证）
	Signature  string   `bson:"signature" json:"signature" redis:"signature"`       //个性签名，30个字
	Money      int64    `bson:"money" json:"money" redis:"money"`                   //当前余额（分为单位，需要除以100）
	TempMoney  int64    `bson:"temp_money" json:"temp_money" redis:"temp_money"`    //待确认金额
	Point      int      `bson:"point" json:"point" redis:"point"`                   //总积分
	Level      int      `bson:"level" json:"level" redis:"level"`                   //等级
}

// 好友简介
type FriendBrief struct {
	UserId   string `bson:"userid" json:"userid"`               // 用户id
	Name     string `bson:"name,omitempty" json:"name"`         // 用户名
	Portrait string `bson:"portrait,omitempty" json:"portrait"` // 头像url
	Star     int    `bson:"star,omitempty" json:"star"`         // 是否是特别关注
	Black    int    `bson:"black,omitempty" json:"black"`       // 是否被拉黑
	Time     int64  `bson:"time,omitempty" json:"-"`            // 成为好友的时间
}

// 用户红包数据
type UserRedpacket struct {
	Income     int              `bson:"income" redis:"income"`           //总收益(对客户端隐藏)
	Outcome    map[string][]int `bson:"outcome" redis:"outcome"`         //总支出(对客户端隐藏)
	DailyGrab  []int64          `bson:"daily_grab" redis:"daily_grab"`   //每日抢红包次数和时间戳[次数,时间戳]
	ShareLimit []int64          `bson:"share_limit" redis:"share_limit"` //分享类红包抢的次数限制，[次数,时间戳](每类红包分开，方便更新缓存)
	OAlimit    []int64          `bson:"oa_limit" redis:"oa_limit"`       //公众号类红包抢的次数限制
}

// 用户完整信息（mongodb）
type User struct {
	Id_           bson.ObjectId `bson:"_id" redis:"-"` //id，唯一账号id
	UserPrivate   `bson:",inline" redis:",inline"`
	UserPublic    `bson:",inline" redis:",inline"`
	UserRedpacket `bson:",inline" redis:",inline"`
	Friends       []*FriendBrief `bson:"friends" redis:"friends"` //好友列表
}

/////////////////////////////////////////////////////////
// 发给客户端个人数据
type S2C_UserData struct {
	UserPublic
	UserId string `json:"id"`              // 用户id
	Oauth  [5]int `json:"oauth,omitempty"` // 第三方账号绑定标示
}

// 好友列表
type FriendList struct {
	Friends []*FriendBrief `bson:"friends" json:"friends"` // 好友列表
}

// 收到的红包信息
type S2C_RedpketRecieveInfo struct {
	Total int `json:"total"` // 总收入
	Rank  int `json:"rank"`  // 排名
}

// 发出的红包信息
type S2C_RedpketSendInfo struct {
	Total  int `json:"total"`  // 总支出
	Amount int `json:"amount"` // 红包的个数
}
