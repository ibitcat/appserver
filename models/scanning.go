/*
	扫红包相关models
*/

package models

import (
	"gopkg.in/mgo.v2/bson"
)

/////////////////////////////////////////////////////////
// 提成以及结算
// 提成百分比为 10%，总金额*10%=实际发给用户红包的总额
// 如果红包有剩余，结算时返回剩余的红包金额+剩余红包扣除的提成，即 剩余红包/(1-提成百分比)
// 发布的扫描红包
type ScanningRedpkt struct {
	Id_          bson.ObjectId `bson:"_id"`           // 扫红包的id
	ItemName     string        `bson:"item_name"`     // 商品名称
	BarCode      string        `bson:"bar_code"`      // 商品条形码
	Tag          string        `bson:"tag"`           // 商品的分类标签
	PicUrls      [3]string     `bson:"pic_urls"`      // 商品图片
	Description  string        `bson:"description"`   // 促销描述信息
	SendTime     int64         `bson:"send_time"`     // 发扫红包的时间
	StartDate    uint32        `bson:"start_date"`    // 开始日期
	StopDate     uint32        `bson:"stop_date"`     // 结束日期
	DailyMoney   uint32        `bson:"daily_money"`   // 每日发的钱
	TotalMoney   uint32        `bson:"total_money"`   // 总金额，没有扣提成的
	Balance      uint32        `bson:"balance"`       // 余额，这里为总金额扣除了提成后的金额
	LastdayMoney uint32        `bson:"lastday_money"` // 最后一天要发的金额（除不尽的情况）
	Isend        bool          `bson:"isend"`         // 该扫红包活动已经结束
	GenDate      uint32        `bson:"gen_date"`      // 最后一次生成红包的日期
}

// 扫描红包 记录
type ScanRecord struct {
	Id_      bson.ObjectId `bson:"_id"`       // 扫红包记录id
	RedpktId string        `bson:"redpkt_id"` // 扫红包的id
	UserId   string        `bson:"user_id"`   // 用户id
	UserName string        `bson:"user_name"` // 用户名，为了不再去查询用户表
	Money    uint32        `bson:"money"`     // 抢到的钱，分分钱为单位
	Time     int64         `bson:"time"`      // 扫描红包时间戳
}

// 扫红包列表查询结果
type ScanningInfo struct {
	Id      bson.ObjectId `bson:"_id"`
	Name    string        `bson:"item_name"`
	Pic     []string      `bson:"pic_urls"`
	Balance uint32        `bson:"balance"`
}

/////////////////////////////////////////////////////////
//
type ScanningItem struct {
	Id      string `json:"id" description:"扫红包商品id"`
	Name    string `json:"name" description:"商品名称"`
	Pic     string `json:"pic" description:"商品图片"`
	Balance uint32 `json:"balance" description:"红包余额"`
}

// 扫红包列表
type ScanningList struct {
	Count uint32          `json:"count" description:"列表长度"`
	List  []*ScanningItem `json:"list" description:"扫红包列表"`
}
