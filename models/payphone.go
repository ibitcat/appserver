// 手机充值相关models

package models

import (
	"gopkg.in/mgo.v2/bson"
)

// 充值记录
type PayPhoneRecord struct {
	Id_       bson.ObjectId `bson:"_id"`         // 索引号
	TradeNo   string        `bosn:"trade_no"`    // 订单号
	UserId    string        `bson:"userid"`      // 用户id
	Type      int           `bson:"type"`        // 充值类型 1=话费，2=流量
	Fee       int           `bson:"fee"`         // 支付费用
	Status    int           `bson:"status"`      // 订单状态 1=已经完成，0=未完成
	CreatTime int64         `bson:"create_time"` // 订单生成时间
}

type PhoneRechargeBalance struct {
	Balance float64 `json:"Balance"` // 余额
}

// server-->client
// 查询话费充值商品信息返回
type PhoneRechargeQueryData struct {
	Cardid   string  `json:"Cardid"`   // 商品编号
	Cardname string  `json:"Cardname"` // 商品名称
	Inprice  float64 `json:"Inprice"`  // 商品价格
	GameArea string  `json:"GameArea"` // 商品归属地
}

type PhoneRechargeQueryResp struct {
	Code int                    `json:"Code"` // 0 表示请求成功 其他为失败
	Msg  string                 `json:"Msg"`  // 表示请求成功或者失败信息
	Data PhoneRechargeQueryData `json:"Data"`
}

// server-->client
// 话费充值返回
type PhoneRechargeData struct {
	Cardid      string  `json:"Cardid"`      // 商品编号
	Cardnum     float64 `json:"Cardnum"`     // 商品面值（充值金额）
	Ordercash   float64 `json:"Ordercash"`   // 商品价格
	Cardname    string  `json:"Cardname"`    // 商品名称
	SporderId   string  `json:"SporderId"`   // APIX订单号
	UserOrderId string  `json:"UserOrderId"` // 商家订单号
	Phone       string  `json:"Phone"`       // 手机号
	State       string  `json:"State"`       // 订单状态（0为充值中，1为成功，其他为失败）
}

type PhoneRechargeResp struct {
	Code int               `json:"Code"` // 0 表示请求成功 其他为失败
	Msg  string            `json:"Msg"`  // 表示请求成功或者失败信息
	Data PhoneRechargeData `json:"Data"`
}

// server-->client
// 查询流量套餐商品信息返回
type DataRechargeUserDataPackages struct {
	PkgId           int     `json:"PkgId"`
	DataValue       string  `json:"DataValue"`
	Price           float64 `json:"Price"`
	Cost            float64 `json:"Cost"`
	Scope           int     `json:"Scope"`
	LimitTimes      string  `json:"LimitTimes"`
	Support4G       int     `json:"Support4G"`
	EffectStartTime int     `json:"EffectStartTime"`
	EffectTime      int     `json:"EffectTime"`
}

type DataRechargeQueryData struct {
	ProviderId       int                            `json:"ProviderId"`
	ProviderName     string                         `json:"ProviderName"`
	UserDataPackages []DataRechargeUserDataPackages `json:"UserDataPackages"`
}

type DataRechargeQueryResp struct {
	Code int                   `json:"Code"`
	Msg  string                `json:"Msg"`
	Data DataRechargeQueryData `json:"Data"`
}

// server-->client
// 流量充值返回
type DataRechargeData struct {
	Ordercash   float64 `json:"Ordercash"`   // 商品价格
	Cardname    string  `json:"Cardname"`    // 商品名称
	SporderId   string  `json:"SporderId"`   // APIX订单号
	UserOrderId string  `json:"UserOrderId"` // 商家订单号
	Phone       string  `json:"Phone"`       // 手机号
	State       string  `json:"State"`       // 订单状态（0为充值中，1为成功，其他为失败）
}

type DataRechargeResp struct {
	Code int              `json:"Code"` // 0 表示请求成功 其他为失败
	Msg  string           `json:"Msg"`  // 表示请求成功或者失败信息
	Data DataRechargeData `json:"Data"`
}
