package v1

import (
	"app-server/apiservice"
	"app-server/models"
	"app-server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// @SubApi 企业认证相关 [/enterprise_cert]
func EnterpriseCert(parentRoute *gin.RouterGroup) {
	router := parentRoute.Group("/enterprise_cert")

	router.Use(middleware.JwtAuthMiddleware())
	router.POST("/submit", submitEnterpriseCertMaterial)
	router.POST("/edit", editEnterpriseCertInfo)
	router.GET("/info", getEnterpriseCertInfo)
}

// @Title submitEnterpriseCertMaterial
// @Description 认证/重新认证
// @Accept  json
// @Param   data  body   models.EnterpriseCertMaterialBinding  true   "提交材料"
// @Success 200 {string} string "成功"
// @Failure 400 {object} models.APIError "失败"
// @Resource /enterprise_cert
// @Router /enterprise_cert/submit [post]
func submitEnterpriseCertMaterial(c *gin.Context) {
	errCode := apiservice.SubmitEnterpriseCertMaterial(c)
	if errCode == 0 {
		c.String(200, "success")
	} else {
		c.JSON(400, &models.APIError{errCode, "submit enterprise cert material failed"})
	}
}

// @Title editEnterpriseCertInfo
// @Description 修改认证信息
// @Accept  json
// @Param   data  body   models.EnterpriseCertInfoBinding  true   "修改后的信息"
// @Success 200 {object} string "成功"
// @Failure 400 {object} models.APIError "失败"
// @Resource /enterprise_cert
// @Router /enterprise_cert/edit [post]
func editEnterpriseCertInfo(c *gin.Context) {
	errCode := apiservice.EditEnterpriseCertInfo(c)
	if errCode == 0 {
		c.String(200, "success")
	} else {
		c.JSON(400, &models.APIError{errCode, "edit enterprise cert info failed"})
	}
}

// @Title getEnterpriseCertInfo
// @Description 获取认证信息
// @Accept  json
// @Success 200 {object} models.S2C_EnterpriseCertInfo "成功"
// @Failure 400 {object} models.APIError "失败"
// @Resource /enterprise_cert
// @Router /enterprise_cert/info [get]
func getEnterpriseCertInfo(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	ret, errCode := apiservice.GetEnterpriseCertInfo(userId)
	if errCode == 0 {
		c.JSON(200, ret)
	} else {
		c.JSON(400, &models.APIError{errCode, "get enterprise cert info failed"})
	}
}
