package models

/*
注意：
gin 的binding 需要注意required tag，如果字段未赋值或者赋值zero value，转换成对应的struct之后，该字段被忽略了，
所以需要注意零值的问题。类似golang内置的 omitempty，例如：`json:"middle_name,omitempty"`
*/

// 发红包
type SendRedpacketBinding struct {
	Type        int             `json:"type" binding:"required" `     //红包类型 1-4
	PerMoney    uint32          `json:"permoney" binding:"required"`  //每个红包的金额金额
	TotalNum    int             `json:"totalnum" binding:"required"`  //红包总个数
	BeginTime   int64           `json:"begintime" binding:"required"` //红包开领时间
	Area        AreaInfo        `json:"area,omitempty"`               //红包发送区域
	Invoice     int             `json:"invoice"`                      //是否需要发票
	Address     InvoiceAddress  `json:"address"`                      //发票快递地址
	Share       ShareInfo       `json:"share,omitempty"`              //文章、广告分享类信息
	App         AppInfo         `json:"app,omitempty"`                //app、游戏刷量类信息
	OfficialAcc OfficialAccInfo `json:"official_acc,omitempty"`       //公众号关注分享类信息
}
