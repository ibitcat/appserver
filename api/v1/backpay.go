package v1

import (
	"app-server/apiservice"
	"app-server/models"
	"app-server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func BackPay(parentRoute *gin.RouterGroup) {
	router := parentRoute.Group("/backpay")

	router.Use(middleware.JwtAuthMiddleware())
	router.POST("/common", backpayCommon)
}

// @Title backpayCommon
// @Description 提现（汇款/银联/支付宝）
// @Accept  string
// @Param   money    form   string  true   "要提现的金额"
// @Param   passwd   form   string  true   "密码"
// @Param   type     form   string  true   "提现类型，1=汇款，2=银联，3=支付宝"
// @Param   account  form   string  true   "汇款/银联：卡号，支付宝：账号"
// @Param   name     form   string  true   "汇款/银联：户名，支付宝：实名"
// @Param   bankname form   string  true   "汇款/银联：开户银行，支付宝：空"
// @Success 200 {string} string "成功"
// @Failure 400 {object} models.APIError "错误"
// @Resource /backpay
// @Router /backpay/common [post]
func backpayCommon(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	errCode := apiservice.BackPayCommon(c.Request, userId)
	if errCode == 0 {
		c.String(200, "成功")
	} else {
		c.JSON(400, models.APIError{errCode, "backpay failed"})
	}
}
