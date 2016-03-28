// 银联支付api

package v1

import (
	"app-server/apiservice"
	"app-server/models"
	"app-server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// @SubApi 银联支付相关 [/unionpay]
func Unionpay(parentRoute *gin.RouterGroup) {
	router := parentRoute.Group("/unionpay")

	router.POST("/notify", unionpayNotify)
	router.POST("/backpaynotify", unionpayBackpayNotify)

	router.Use(middleware.JwtAuthMiddleware())
	router.GET("/tradeno", getUnionpayTradeNo)
	router.GET("/tradestatus", queryUnionpayTradeStatus)
	router.GET("/backpay", unionBackpay)
}

// @Title unionpay_notify
// @Description 银联异步回调api
// @Accept  string
// @Success 200 {string} string "回调成功"
// @Failure 400 {string} string "回调失败"
// @Resource /unionpay
// @Router /unionpay/notify [post]
func unionpayNotify(c *gin.Context) {
	err := apiservice.UnionpayNotify(c.Request)
	if err == nil {
		c.String(200, "success")
	} else {
		c.String(400, err.Error())
	}
}

// @Title unionpay_backpay
// @Description 银联提现回调api
// @Accept  string
// @Success 200 {string} string "回调成功"
// @Failure 400 {string} string "回调失败"
// @Resource /unionpay
// @Router /unionpay/backpaynotify [post]
func unionpayBackpayNotify(c *gin.Context) {
	err := apiservice.UnionBackpayNotify(c.Request)
	if err == nil {
		c.String(200, "success")
	} else {
		c.String(400, err.Error())
	}
}

// @Title TradeNo
// @Description 获取银联交易流水号
// @Accept  json
// @Param   redpktid  query   string  true   "要付款的红包id"
// @Success 200 {string} string "获取成功"
// @Failure 400 {object} models.APIError "获取失败"
// @Resource /unionpay
// @Router /unionpay/tradeno [get]
func getUnionpayTradeNo(c *gin.Context) {
	//userId := c.MustGet("userId").(string)
	redpktId := c.Query("redpktid")
	tn, err := apiservice.GetUnionpayTradeNo(redpktId)

	if err == nil {
		c.JSON(200, gin.H{"tn": tn})
	} else {
		c.JSON(400, models.APIError{1111, err.Error()})
	}
}

// @Title TradeStatus
// @Description 查询订单交易状态
// @Accept  json
// @Param   redpktid  query   string  true   "要付款的红包id"
// @Success 200 {string} string "成功"
// @Failure 400 {string} string "失败"
// @Resource /unionpay
// @Router /unionpay/tradestatus [get]
func queryUnionpayTradeStatus(c *gin.Context) {
	//userId := c.MustGet("userId").(string)
	redpktId := c.Query("redpktid")
	err := apiservice.QueryUnionpayStatus(redpktId)

	if err == nil {
		c.JSON(200, "ok")
	} else {
		c.JSON(400, "fail")
	}
}

// @Title TradeStatus
// @Description 请求提现
// @Accept  json
// @Param   money  query   string  true   "要提现的金额"
// @Success 200 {string} string "成功"
// @Failure 400 {string} string "失败"
// @Resource /unionpay
// @Router /unionpay/backpay [get]
func unionBackpay(c *gin.Context) {
	//userId := c.MustGet("userId").(string)
	//money := c.Query("money")
	err := apiservice.UnionBackpay()

	if err == nil {
		c.JSON(200, "ok")
	} else {
		c.JSON(400, "fail")
	}
}
