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
// 玩家提现记录
type BackpayRecord struct {
	Id_      bson.ObjectId `bson:"_id"`       // 记录id
	UserId   string        `bson:"user_id"`   // 用户id
	Type     int           `bson:"type"`      // 记录类型，1=汇款，2=银联，3=支付宝，4=微信
	Account  string        `bson:"account"`   // 汇款/银联：卡号，支付宝：账号，微信：没用
	Name     string        `bson:"name"`      // 汇款/银联：户名，支付宝：实名，微信：实名
	BankName string        `bson:"bank_name"` // 汇款/银联：开户银行，支付宝：没用
	Fee      int64         `bson:"fee"`       // 金额
	Time     int64         `bson:"time"`      // 提现时间
	Status   int           `bson:"status"`    // 状态，0=未处理，1=处理成功，处理失败
}
