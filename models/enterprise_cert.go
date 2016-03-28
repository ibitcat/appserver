package models

import (
	"gopkg.in/mgo.v2/bson"
)

// 企业认证材料
type EnterpriseCertMaterial struct {
	Id_             bson.ObjectId `bson:"_id"`
	UserId          string        `bson:"user_id"`          // 用户id
	Name            string        `bson:"name"`             // 认证名称
	EnterpriseName  string        `bson:"enterprise_name"`  // 企业名称
	OperateName     string        `bson:"operate_name"`     // 运营者身份证姓名
	OperatePlace    string        `bson:"operate_place"`    // 运营者职务
	OperateId       string        `bson:"operate_id"`       // 运营者身份证
	OfficialWebsite string        `bson:"official_website"` // 官网
	Weibo           string        `bson:"weibo"`            // 微博
	Weixin          string        `bson:"weixin"`           // 微信公众号
	BusinessLicense string        `bson:"business_license"` // 企业法人营业执照
	TrademarkCert   string        `bson:"trademark_cert"`   // 商标注册证
	OperateIdPhoto  string        `bson:"operate_id_photo"` // 手持身份证照片
	Status          uint8         `bson:"status"`           // 认证状态（0=未认证，1=认证通过，2=认证失败）
	Reason          string        `bson:"reason"`           // 审核结果（失败原因）
}

//
// S2C
//

type S2C_EnterpriseCertInfo struct {
	CertName        string `json:"cert_name"`        // 认证名称
	OfficialWebsite string `json:"official_website"` // 官网
	Weibo           string `json:"weibo"`            // 微博
	Weixin          string `json:"weixin"`           // 微信公众号
}

//
// binding
//

/*
注意：
gin 的binding 需要注意required tag，如果字段未赋值或者赋值zero value，转换成对应的struct之后，该字段被忽略了，
所以需要注意零值的问题。类似golang内置的 omitempty，例如：`json:"middle_name,omitempty"`
*/

type EnterpriseCertMaterialBinding struct {
	Name            string `json:"name" binding:"required"`             // 认证昵称
	EnterpriseName  string `json:"enterprise_name" binding:"required"`  // 企业名称
	OperateName     string `json:"operate_name" binding:"required"`     // 运营者身份证姓名
	OperatePlace    string `json:"operate_place" binding:"required"`    // 运营者职务
	OperateId       string `json:"operate_id" binding:"required"`       // 运营者身份证
	OfficialWebsite string `json:"official_website,omitempty"`          // 官网
	Weibo           string `json:"weibo,omitempty"`                     // 微博
	Weixin          string `json:"weixin,omitempty"`                    // 微信公众号
	BusinessLicense string `json:"business_license" binding:"required"` // 企业法人营业执照
	TrademarkCert   string `json:"trademark_cert,omitempty"`            // 商标注册证
	OperateIdPhoto  string `json:"operate_id_photo" binding:"required"` // 手持身份证照片
}

type EnterpriseCertInfoBinding struct {
	OfficialWebsite string `json:"official_website,omitempty"` // 官网
	Weibo           string `json:"weibo,omitempty"`            // 微博
	Weixin          string `json:"weixin,omitempty"`           // 微信公众号
}
