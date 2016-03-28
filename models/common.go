package models

// 地理信息
type AreaInfo struct {
	Country  string `bson:"country" json:"country"`   // 国家
	Province string `bson:"province" json:"province"` // 省份
	City     string `bson:"city" json:"city"`         // 市
	District string `bson:"district" json:"district"` // 区
}
