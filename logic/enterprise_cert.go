package logic

import (
	"fmt"

	"app-server/define"
	"app-server/models"
	"app-server/pkg/mongodb"

	"gopkg.in/mgo.v2/bson"
)

// 获取认证状态
// 返回值 0=无认证，1=认证中，2=认证通过
func GetEnterpriseCertStatus(userId string) int {
	var material models.EnterpriseCertMaterial
	err := mongodb.SelectOne(define.EnterpriseCertCollection, bson.M{"_id": bson.ObjectIdHex(userId)}, nil, &material)
	if err != nil {
		return 0
	}
	if material.Status == 0 {
		// 有数据但是还没有审核
		return 1
	} else if material.Status == 1 {
		// 审核通过
		return 2
	} else if material.Status == 2 {
		// 审核不通过，所以无认证
		return 0
	}
	fmt.Println("无效状态", material.Status)
	return 0
}

// 获取企业认证材料
func GetEnterpriseCertMaterial(userId string) (*models.EnterpriseCertMaterial, error) {
	var material models.EnterpriseCertMaterial
	err := mongodb.SelectOne(define.EnterpriseCertCollection, bson.M{"_id": bson.ObjectIdHex(userId)}, nil, &material)
	if err != nil {
		return nil, err
	}
	return &material, nil
}

// 创建企业认证
func CreateEnterpriseCert(material *models.EnterpriseCertMaterial) error {
	if material == nil {
		return fmt.Errorf("material nil")
	}
	if _, err := GetEnterpriseCertMaterial(material.UserId); err == nil {
		return mongodb.UpdateById(define.EnterpriseCertCollection, material.Id_, material)
	} else {
		return mongodb.Insert(define.EnterpriseCertCollection, material)
	}
}

// 修改信息
func EditEnterpriseCertInfo(userId string, info *models.EnterpriseCertInfoBinding) error {
	if info == nil {
		return fmt.Errorf("info nil")
	}
	update := bson.M{"$set": bson.M{
		"official_website": info.OfficialWebsite,
		"weibo":            info.Weibo,
		"weixin":           info.Weixin,
	}}
	return mongodb.Update(define.EnterpriseCertCollection, bson.M{"_id": bson.ObjectIdHex(userId)}, update)
}

// 审核结果
func AuditEnterpriseCert(userId string, pass bool, reason string) {
	if pass {
		mongodb.Update(define.EnterpriseCertCollection, bson.M{"_id": bson.ObjectIdHex(userId)}, bson.M{"$set": bson.M{"status": 1, "reason": reason}})
	} else {
		mongodb.Update(define.EnterpriseCertCollection, bson.M{"_id": bson.ObjectIdHex(userId)}, bson.M{"$set": bson.M{"status": 2, "reason": reason}})
	}
}
