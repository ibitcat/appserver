/*
红包相关
*/

package models

import (
	"gopkg.in/mgo.v2/bson"
)

/////////////////////////////////////////////////////////
// 发票地址
type InvoiceAddress struct {
	Title     string `bson:"title" json:"title"`         // 发票抬头
	Addressee string `bson:"addressee" json:"addressee"` // 收件人
	Tel       string `bson:"tel" json:"tel"`             // 联系电话
	Address   string `bson:"address" json:"address"`     // 联系地址
}

// 转发分享类红包的信息
type ShareInfo struct {
	Title    string `bson:"title" json:"title"`         // 广告/文章标题
	ImageUri string `bson:"image_uri" json:"image_uri"` // 图片地址
}

// 刷榜类红包的信息
type AppInfo struct {
	Size      string `bson:"size" json:"size"`             // 游戏或app大小
	Name      string `bson:"name" json:"name"`             // 游戏或app名字
	IconUrl   string `bson:"icon_url" json:"icon_url"`     // app图标链接
	Keyword   string `bson:"keyword" json:"keyword"`       // 关键字
	UrlScheme string `bson:"url_scheme" json:"url_scheme"` // ios调用url
	BundleId  string `bson:"bundle_id" json:"bundle_id"`   // 游戏唯一标识符
	Price     int    `bson:"price" json:"price"`           // appstore 价格 0=免费
	Platform  int    `bson:"platform" json:"platform"`     // 投放平台 0=全部 1=iphone 2=ipad 3=Android
}

// 公众号类红包的信息
type OfficialAccInfo struct {
	Name  string `bson:"name" json:"name"`   // 公众号名字
	Url   string `bson:"url" json:"url"`     // 文章链接
	Title string `bson:"title" json:"title"` // 文章标题
}

type RedpktStatistics struct {
	Area  map[string]int `bson:"area" json:"area"`   // 区域分布统计
	Sex   []int          `bson:"sex" json:"sex"`     // 性别分布统计 [女，男，未知]
	Count int            `bson:"count" json:"count"` // 下载或转发的总次数
}

// 红包
// 红包类型：1=分享;2=游戏;3=app刷量;4=公众号
type RedPacket struct {
	Id_         bson.ObjectId    `bson:"_id" redis:"-"`          //红包id
	SenderId    string           `bson:"sender_id"`              //红包发送者id
	CreateTime  int64            `bson:"create_time"`            //红包创建时间
	BeginTime   int64            `bson:"begin_time"`             //红包开领时间
	EndTime     int64            `bson:"end_time"`               //红包下架时间
	Verify      int              `bson:"verify"`                 //红包审核状态 0=审核中，1=已通过，2=未通过
	Year        string           `bson:"year"`                   //红包创建年份
	TradeStatus int              `bson:"trade_status"`           //红包的付款状态 1=已付款，0=等待付款
	Rebate      int              `bson:"rebate"`                 //红包提成 10%
	PerMoney    uint32           `bson:"per_money"`              //红包单个金额
	Total       int              `bson:"total"`                  //红包总个数
	Remainder   int              `bson:"remainder"`              //红包剩余个数
	Area        AreaInfo         `bson:"area,omitempty"`         //区域
	Invoice     int              `bson:"invoice"`                //是否需要发票
	Address     InvoiceAddress   `bson:"address,omitempty"`      //发票邮寄地址
	Type        int              `bson:"type"`                   //红包类型 [看注释]
	Share       ShareInfo        `bson:"share,omitempty"`        //文章、广告分享类信息
	App         AppInfo          `bson:"app,omitempty"`          //app、游戏刷量类信息
	OfficialAcc OfficialAccInfo  `bson:"official_acc,omitempty"` //公众号关注分享类信息
	Statistics  RedpktStatistics `bson:"statistics"`             //分布统计
}

// 领取红包记录
type GrabRecord struct {
	Id_           bson.ObjectId `bson:"_id" json:"-"`                         //抢红包记录id
	UserId        string        `bson:"userid" json:"-"`                      //用户id
	UserName      string        `bson:"user_name" json:"-"`                   //用户昵称
	UserProvince  string        `bson:"user_province" json:"-"`               //用户所在省份
	UserSex       int           `bson:"user_sex" json:"-"`                    //用户性别 0=女 1=男 2=未知
	RedpacketId   string        `bson:"redpacket_id" json:"redpacket_id"`     //红包id
	RedpacketType int           `bson:"redpacket_type" json:"redpacket_type"` //红包类型
	RedpacketName string        `bson:"redpacket_name" json:"redpacket_name"` //红包名称
	GrabMoney     uint32        `bson:"grab_money" json:"grab_money"`         //抢到的金额 分分钱
	GrabTime      int64         `bson:"grab_time" json:"grab_time"`           //抢红包的时间
	GrabDate      string        `bson:"grab_date" json:"-"`                   //抢红包的日期
}

/////////////////////////////////////////////////////////
// redis
// 用户抢红包的状态
type GrabStatus struct {
	Status int    `json:"status"` // 抢红包的状态 1=分享或下载
	Expire int64  `json:"expire"` // 到期时间
	Device string `json:"device"` // 设备id
}

/////////////////////////////////////////////////////////
// 返回给客户端
// 红包信息
type RedpacketInfo struct {
	Id          string           `json:"id"`                     //红包id
	BeginTime   int64            `json:"begin_time"`             //红包开始时间
	EndTime     int64            `json:"end_time"`               //红包结束时间
	PerMoney    uint32           `json:"per_money"`              //每个红包的金额
	Number      int              `json:"number"`                 //红包剩余个数
	UserId      string           `json:"user_id"`                //红包发送者id
	UserName    string           `json:"user_name"`              //发送者名字
	Portrait    string           `json:"portrait"`               //发送者图像
	IsAuth      uint8            `json:"is_auth"`                //是否是认证商户
	Type        int              `json:"type"`                   //红包类型
	Area        AreaInfo         `json:"area,omitempty"`         //区域
	Share       *ShareInfo       `json:"share,omitempty"`        //文章、广告分享类信息
	App         *AppInfo         `json:"app,omitempty"`          //app、游戏刷量类信息
	OfficialAcc *OfficialAccInfo `json:"official_acc,omitempty"` //公众号关注分享类信息
}

// 红包列表
type S2C_RedpacketList struct {
	Count uint32           `description:"该时段红包个数"`
	List  []*RedpacketInfo `description:"红包列表"`
}

// 红包到期时间列表
type RedpacketExpire struct {
	Id         string `json:"id"`         // 红包id
	Type       int    `json:"type"`       // 红包类型
	Share      int64  `json:"share"`      // 5分钟分享到期时间戳
	Screenshot int64  `json:"screenshot"` // 6小时截图到期时间戳
	Download   int64  `json:"download"`   // 30分钟下载到期时间
	IsGrab     int    `json:"is_grab"`    // 是否已经抢到红包
}

// 用户的红包到期列表
type UserExpireList struct {
	List []*RedpacketExpire `json:"list"` // 到期时间戳列表
}

type RedpktRecord struct {
	UserId   string `json:"userid"`   // 用户id
	UserName string `json:"nickname"` // 用户昵称
	Time     int64  `json:"time"`     // 领取时间
}

// 红包领取记录
type S2C_RedpktRecord struct {
	Id   string          `json:"id"`             // 红包id
	List []*RedpktRecord `json:"list,omitempty"` // 领取记录
}

// 收到的红包记录
type S2C_ReceivedList struct {
	List []*GrabRecord `json:"list"`
}

type SendRedpacket struct {
	Id        string `json:"id"`         //红包id
	BeginTime int64  `json:"begin_time"` //红包开领时间
	EndTime   int64  `json:"end_time"`   //红包下架时间
	Title     string `json:"title"`      //红包标题
	PerMoney  uint32 `json:"per_money"`  //红包单个金额
	Total     int    `json:"total"`      //红包总个数
	Remainder int    `json:"remainder"`  //红包剩余个数
}

// 用户发送的红包记录
type S2C_RedpktSendList struct {
	List []*SendRedpacket `json:"list,omitempty"`
}

// 待发布红包
type ToBeReleasedRedpkt struct {
	Id          string `json:"id"`           //红包id
	CreateTime  int64  `json:"create_time"`  //红包创建时间
	BeginTime   int64  `json:"begin_time"`   //红包开领时间
	Verify      int    `json:"verify"`       //审核状态
	TradeStatus int    `json:"trade_status"` //付款状态
	Title       string `json:"title"`        //红包标题
}

// 用户待发布的红包列表
type S2C_ToBeReleasedList struct {
	List []*ToBeReleasedRedpkt `json:"list"`
}

type S2C_RedpktStatistics struct {
	Id               string                        `json:"id"`
	RedpktStatistics `json:"statistics,omitempty"` // 统计数据
}
