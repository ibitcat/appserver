// gm操作

package v1

import (
	"app-server/apiservice"
	"app-server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// @SubApi gm操作 [/gm]
func Gm(parentRoute *gin.RouterGroup) {
	router := parentRoute.Group("/gm")

	router.Use(middleware.JwtAuthMiddleware())
	router.POST("/audit_redpacket", auditRedpacket)
	router.POST("/manual_charge", manualCharge)
}

// @Title auditRedpacket
// @Description 红包审核
// @Accept  json
// @Param   redpktId form   string  true   "红包id"
// @Param   status   form   string  true   "审核状态，1=成功，2=失败"
// @Success 201 {object} string "成功"
// @Failure 400 {object} string "失败信息"
// @Resource /gm
// @Router /gm/audit_redpacket [post]
func auditRedpacket(c *gin.Context) {
	err := apiservice.AuditRedpacket(c)
	if err == nil {
		c.JSON(200, "success")
	} else {
		c.JSON(400, err.Error())
	}
}

// @Title manualCharge
// @Description 手动充值
// @Accept  json
// @Param   target_phone     form   string  true   "充值用户手机号"
// @Param   target_nickname  form   string  true   "充值用户昵称"
// @Param   target_account   form   string  true   "充值用户红包好"
// @Param   amount           form   string  true   "充值用户金额（分）"
// @Success 201 {object} string "成功"
// @Failure 400 {object} string "失败信息"
// @Resource /gm
// @Router /gm/manual_charge [post]
func manualCharge(c *gin.Context) {
	err := apiservice.ManualCharge(c)
	if err == nil {
		c.JSON(200, "success")
	} else {
		c.JSON(400, err.Error())
	}
}
