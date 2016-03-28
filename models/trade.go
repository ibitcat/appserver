// 支付相关models

package models

import (
	"gopkg.in/mgo.v2/bson"
)

// 订单记录
type TradeInfo struct {
	Id_       bson.ObjectId `bson:"_id"`         // 索引号
	TradeNo   string        `bosn:"trade_no"`    // 订单号(红包号)
	UserId    string        `bson:"userid"`      // 用户id
	Fee       int           `bson:"fee"`         // 支付费用
	Status    int           `bson:"status"`      // 订单状态 1=已经完成，0=未完成
	CreatTime int64         `bson:"create_time"` // 订单生成时间
}

// server-->client
// 支付宝支付参数
type AlipayParams struct {
	AlipayPid string `json:"alipaypid"` // 支付宝pid
	AlipayAcc string `json:"alipayacc"` // 支付宝账号
	NotifyUrl string `json:"notifyurl"` // 支付宝回调
}

// server-->api server
// 微信支付回调结果
type WechatPayResult struct {
	ReturnCode string `xml:"return_code"` // 错误码，SUCCESS/FAIL
	ReturnMsg  string `xml:"return_msg"`  // 错误信息
}

// server-->client
// 微信支付参数
type WechatPayParams struct {
	AppId     string `json:"appid"`     // appid
	MchId     string `json:"mchid"`     // mchid
	NotifyUrl string `json:"notifyurl"` // 微信支付回调
	PayKey    string `json:"paykey"`    // 支付key
}

// api server-->server
type WechatBackPayResult struct {
	ReturnCode     string  `xml:"return_code"`
	ReturnMsg      *string `xml:"return_msg,omitempty"`
	MchAppid       *string `xml:"mch_appid,omitempty"`
	Mchid          *string `xml:"mchid,omitempty"`
	DeviceInfo     *string `xml:"device_info,omitempty"`
	NonceStr       *string `xml:"nonce_str,omitempty"`
	ResultCode     *string `xml:"result_code,omitempty"`
	ErrCode        *string `xml:"err_code,omitempty"`
	ErrCodeDes     *string `xml:"err_code_des,omitempty"`
	PartnerTradeNo *string `xml:"partner_trade_no,omitempty"`
	PaymentNo      *string `xml:"payment_no,omitempty"`
	PaymentTime    *string `xml:"payment_time,omitempty"`
}
