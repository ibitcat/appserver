package v1

import (
	"fmt"

	"app-server/apiservice"
	"app-server/models"
	"app-server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// @SubApi 微信支付相关 [/wechatpay]
func Wechatpay(parentRoute *gin.RouterGroup) {
	router := parentRoute.Group("/wechatpay")

	router.POST("/notify", wechatpayNotify)

	router.Use(middleware.JwtAuthMiddleware())
	router.GET("/params", getWechatpayParams)
	router.POST("/backpay", wechatBackPay)
}

// @Title wechatpayNotify
// @Description 微信付款异步回调api
// @Accept  string
// @Success 200 {object} models.WechatPayResult "微信支付回调成功"
// @Failure 400 {object} models.WechatPayResult "微信支付回调失败"
// @Resource /wechatpay
// @Router /wechatpay/notify [post]
func wechatpayNotify(c *gin.Context) {
	var ret models.WechatPayResult
	err := apiservice.WechatpayNoify(c.Request)
	if err == nil {
		ret.ReturnCode = "SUCCESS"
		ret.ReturnMsg = "OK"
		c.XML(200, &ret)
	} else {
		fmt.Println("微信充值错误 ", err.Error())
		ret.ReturnCode = "FAIL"
		ret.ReturnMsg = err.Error()
		c.XML(400, &ret)
	}
}

// @Title getWechatpayParams
// @Description 获取微信支付参数
// @Accept  json
// @Success 200 {object} models.WechatPayParams "成功"
// @Resource /wechatpay
// @Router /wechatpay/params [get]
func getWechatpayParams(c *gin.Context) {
	c.JSON(200, apiservice.GetWechatPayParams())
}

// @Title wechatBackPay
// @Description 微信提现
// @Accept  json
// @Param   money    query   string  true   "要提现的金额"
// @Param   realname query   string  true   "收款用户真实姓名"
// @Param   passwd   query   string  true   "密码"
// @Success 200 {string} string "成功"
// @Failure 400 {object} models.APIError "失败"
// @Resource /wechatpay
// @Router /wechatpay/backpay [post]
func wechatBackPay(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	errCode := apiservice.WechatBackPay(userId, c.Request)
	if errCode == 0 {
		c.String(200, "success")
	} else {
		c.JSON(400, &models.APIError{errCode, "wechat backpay failed"})
	}
}
