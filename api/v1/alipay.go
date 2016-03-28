// 支付宝相关api

package v1

import (
	"app-server/apiservice"
	"app-server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// @SubApi 支付宝相关 [/alipay]
func Alipay(parentRoute *gin.RouterGroup) {
	router := parentRoute.Group("/alipay")

	router.POST("/notify", alipayNotify)

	router.Use(middleware.JwtAuthMiddleware())
	router.GET("/tradeno", getOutTradeNo)
	router.GET("/params", getAlipayParams)
}

// @Title alipay_notify
// @Description 支付宝异步回调api
// @Accept  json
// @Success 200 {string} string "回调成功"
// @Failure 400 {string} string "回调失败"
// @Resource /alipay
// @Router /alipay/notify [post]
func alipayNotify(c *gin.Context) {
	ok := apiservice.AlipayNotify(c.Request)
	if ok {
		c.String(200, "success")
	} else {
		c.String(400, "fail")
	}
}

// @Title outTradeNo
// @Description 生成订单号
// @Accept  json
// @Success 200 {string} string "生成订单成功"
// @Failure 400 {string} string "生成订单失败"
// @Resource /alipay
// @Router /alipay/tradeno [get]
func getOutTradeNo(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	apiservice.MakeOutTradeNo(userId)

	c.String(200, "success")
}

// @Title alipayParams
// @Description 获取支付宝支付参数
// @Accept  json
// @Success 200 {object} models.AlipayParams "成功"
// @Failure 400 {string} string "失败"
// @Resource /alipay
// @Router /alipay/params [get]
func getAlipayParams(c *gin.Context) {
	params := apiservice.GetAlipayParams()

	c.JSON(200, params)
}
