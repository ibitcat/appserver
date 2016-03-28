// 手机充值
package v1

import (
	"fmt"

	"app-server/apiservice"
	"app-server/models"
	"app-server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// @SubApi 手机充话费/流量 [/payphone]
func PayPhone(parentRoute *gin.RouterGroup) {
	router := parentRoute.Group("/payphone")

	router.GET("/phone_recharge_notify", phoneRechargeNotify)
	router.GET("/data_recharge_notify", dataRechargeNotify)

	router.GET("/phone_recharge_balance", phoneRechargeBalance)
	router.GET("/phone_recharge_query", phoneRechargeQuery)
	router.GET("/data_recharge_query", dataRechargeQuery)

	router.Use(middleware.JwtAuthMiddleware())
	router.POST("/phone_recharge", phoneRecharge)
	router.POST("/data_recharge", dataRecharge)
}

// @Title phoneRechargeBalance
// @Description 查询余额
// @Accept  string
// @Success 200 {object} models.PhoneRechargeBalance "成功"
// @Failure 400 {object} models.APIError "错误"
// @Resource /payphone
// @Router /payphone/phone_recharge_balance [get]
func phoneRechargeBalance(c *gin.Context) {
	ret, err := apiservice.PayPhoneBalance()
	if err == nil {
		c.JSON(200, &models.PhoneRechargeBalance{ret})
	} else {
		fmt.Println("查询余额错误", err.Error())
		c.JSON(400, &models.APIError{10501, "query balance failed"})
	}
}

// @Title phoneRechargeQuery
// @Description 根据手机号和充值额度查询商品信息
// @Accept  string
// @Param   phone  query  string  true  "手机号"
// @Param   price  query  string  true  "充值金额（有效值参考apix）"
// @Success 200 {object} models.PhoneRechargeQueryResp "apix数据直接返回"
// @Failure 400 {object} models.APIError "错误"
// @Resource /payphone
// @Router /payphone/phone_recharge_query [get]
func phoneRechargeQuery(c *gin.Context) {
	c.Request.ParseForm()
	var (
		phone = c.Request.FormValue("phone")
		price = c.Request.FormValue("price")
	)
	ret, err := apiservice.PhoneRechargeQuery(phone, price)
	if err == nil {
		c.JSON(200, ret)
	} else {
		fmt.Println("查询充值商品信息错误", err.Error())
		c.JSON(400, &models.APIError{10502, "query phone recharge product failed"})
	}
}

// @Title phoneRecharge
// @Description 充值话费
// @Accept  string
// @Param   phone  form  string  true  "手机号"
// @Param   price  form  string  true  "充值金额（有效值参考apix）"
// @Success 200 {object} models.PhoneRechargeResp "apix数据直接返回"
// @Failure 400 {object} models.APIError "错误"
// @Resource /payphone
// @Router /payphone/phone_recharge [post]
func phoneRecharge(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	c.Request.ParseForm()
	var (
		phone = c.Request.FormValue("phone")
		price = c.Request.FormValue("price")
	)
	ret, errCode := apiservice.PhoneRecharge(userId, phone, price)
	if errCode == 0 {
		c.JSON(200, ret)
	} else {
		c.JSON(400, &models.APIError{errCode, "phone recharge failed"})
	}
}

// @Title dataRechargeQuery
// @Description 查询号码支持的流量套餐
// @Accept  string
// @Param   phone  query  string  true  "手机号"
// @Success 200 {object} models.DataRechargeQueryResp "apix数据直接返回"
// @Failure 400 {object} models.APIError "错误"
// @Resource /payphone
// @Router /payphone/data_recharge_query [get]
func dataRechargeQuery(c *gin.Context) {
	c.Request.ParseForm()
	var phone = c.Request.FormValue("phone")
	ret, err := apiservice.DataRechargeQuery(phone)
	if err == nil {
		c.JSON(200, ret)
	} else {
		fmt.Println("查询流量套餐信息错误", err.Error())
		c.JSON(400, &models.APIError{10505, "query data recharge product failed"})
	}
}

// @Title dataRecharge
// @Description 充值流量
// @Accept  string
// @Param   phone  form  string  true  "手机号"
// @Param   pkgid  form  string  true  "套餐ID（根据查询号码支持的流量套餐接口得到）"
// @Success 200 {object} models.DataRechargeResp "apix数据直接返回"
// @Failure 400 {object} models.APIError "错误"
// @Resource /payphone
// @Router /payphone/data_recharge [post]
func dataRecharge(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	c.Request.ParseForm()
	var (
		phone = c.Request.FormValue("phone")
		pkgid = c.Request.FormValue("pkgid")
	)

	ret, errCode := apiservice.DataRecharge(userId, phone, pkgid)
	if errCode == 0 {
		c.JSON(200, ret)
	} else {
		c.JSON(400, &models.APIError{errCode, "data recharge failed"})
	}
}

// @Title phoneRechargeNotify
// @Description 充值话费回调
// @Accept  string
// @Param   state      query  string  true  "充值状态"
// @Param   orderid    query  string  true  "商家订单号"
// @Param   ordertime  query  string  true  "订单处理时间"
// @Param   sign       query  string  true  "签名"
// @Param   err_msg    query  string  true  "失败信息"
// @Success 200 {string} string "SUCCESS"
// @Failure 400 {string} string "错误信息"
// @Resource /payphone
// @Router /payphone/phone_recharge_notify [get]
func phoneRechargeNotify(c *gin.Context) {
	c.Request.ParseForm()
	var (
		state     = c.Request.FormValue("state")     // 充值状态（0为充值中 1为成功 其他为失败）
		orderid   = c.Request.FormValue("orderid")   // 商家订单号
		ordertime = c.Request.FormValue("ordertime") // 订单处理时间 (格式为：yyyyMMddHHmmss  如：20150323140214）)
		sign      = c.Request.FormValue("sign")      // 32位小写md5签名：md5(apix-key + orderid+ ordertime)
		errMsg    = c.Request.FormValue("err_msg")   // 充值失败时候返回失败信息。成功时为空。
	)
	fmt.Println(state, orderid, ordertime, sign, errMsg)

	err := apiservice.PhoneRechargeNotify(state, orderid, ordertime, sign, errMsg)
	if err != nil {
		fmt.Println("充值话费回调失败", err.Error())
		c.String(400, err.Error())
	} else {
		c.String(200, "SUCCESS")
	}
}

// @Title dataRechargeNotify
// @Description 充值流量回调
// @Accept  string
// @Param   state      query  string  true  "充值状态"
// @Param   orderid    query  string  true  "商家订单号"
// @Param   ordertime  query  string  true  "订单处理时间"
// @Param   sign       query  string  true  "签名"
// @Param   err_msg    query  string  true  "失败信息"
// @Success 200 {string} string "SUCCESS"
// @Failure 400 {string} string "错误信息"
// @Resource /payphone
// @Router /payphone/data_recharge_notify [get]
func dataRechargeNotify(c *gin.Context) {
	c.Request.ParseForm()
	var (
		state     = c.Request.FormValue("state")     // 充值状态（0为充值中 1为成功 其他为失败）
		orderid   = c.Request.FormValue("orderid")   // 商家订单号
		ordertime = c.Request.FormValue("ordertime") // 订单处理时间 (格式为：yyyyMMddHHmmss  如：20150323140214）)
		sign      = c.Request.FormValue("sign")      // 32位小写md5签名：md5(apix-key + orderid+ ordertime)
		errMsg    = c.Request.FormValue("err_msg")   // 充值失败时候返回失败信息。成功时为空。
	)
	fmt.Println(state, orderid, ordertime, sign, errMsg)

	err := apiservice.DataRechargeNotify(state, orderid, ordertime, sign, errMsg)
	if err != nil {
		fmt.Println("充值流量回调失败", err.Error())
		c.String(400, err.Error())
	} else {
		c.String(200, "SUCCESS")
	}
}
