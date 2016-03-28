package apiservice

import (
	"app-server/define"
	"fmt"

	"app-server/logic"
	"app-server/models"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func SubmitEnterpriseCertMaterial(c *gin.Context) uint32 {
	userId := c.MustGet("userId").(string)
	var material models.EnterpriseCertMaterialBinding
	err := c.BindJSON(&material)
	if err != nil {
		fmt.Println("绑定json错误", err.Error())
		return 10551
	}
	err = logic.CreateEnterpriseCert(&models.EnterpriseCertMaterial{
		Id_:             bson.ObjectIdHex(userId),
		UserId:          userId,
		Name:            material.Name,
		EnterpriseName:  material.EnterpriseName,
		OperateName:     material.OperateName,
		OperatePlace:    material.OperatePlace,
		OperateId:       material.OperateId,
		OfficialWebsite: material.OfficialWebsite,
		Weibo:           material.Weibo,
		Weixin:          material.Weixin,
		BusinessLicense: material.BusinessLicense,
		TrademarkCert:   material.TrademarkCert,
		OperateIdPhoto:  material.OperateIdPhoto,
		Status:          0,
	})
	if err != nil {
		fmt.Println("创建认证材料错误", err.Error())
		return 10551
	}

	err = UpdatePersonalInfo(userId, define.EUser_Cert, 1)
	if err != nil {
		fmt.Println("修改用户信息错误", err.Error())
		return 10551
	}
	return 0
}

func EditEnterpriseCertInfo(c *gin.Context) uint32 {
	userId := c.MustGet("userId").(string)
	var info models.EnterpriseCertInfoBinding
	c.BindJSON(&info)
	err := logic.EditEnterpriseCertInfo(userId, &info)
	if err != nil {
		fmt.Println("修改失败", err.Error())
		return 10552
	}

	return 0
}

func GetEnterpriseCertInfo(userId string) (*models.S2C_EnterpriseCertInfo, uint32) {
	material, err := logic.GetEnterpriseCertMaterial(userId)
	if err != nil {
		fmt.Println("获取material失败", err.Error())
		return nil, 10553
	}

	ret := &models.S2C_EnterpriseCertInfo{
		CertName:        material.Name,
		OfficialWebsite: material.OfficialWebsite,
		Weibo:           material.Weibo,
		Weixin:          material.Weixin,
	}
	return ret, 0
}
