/*
抢红包api
*/

package v1

import (
	//"fmt"

	"app-server/apiservice"
	"app-server/models"
	"app-server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// @SubApi 抢/发红包 [/redpacket]
func Redpacket(parentRoute *gin.RouterGroup) {
	router := parentRoute.Group("/redpacket")

	// 不需要登陆的操作
	router.GET("/list", getRedpacketList)
	router.GET("/paytest", testPayRedpacket)
	router.GET("/verify", verifyAppRedpacket)
	router.GET("/record", getRecordList)

	// 需要登陆的操作
	router.Use(middleware.JwtAuthMiddleware())
	router.POST("/send", sendRedpacket)
	router.GET("/grab", grabRedpacket)
	router.GET("/share", finishRedpacketShare)
	router.GET("/finish", finishRedpacket)
	router.GET("/giveup", giveupRedpacket)
	router.POST("/pay", paymentByBalance)
	router.GET("/filter", filterByUser)
	router.GET("/statistics", getRedpacketStatistics)
	router.GET("/redpkt_recieve", getRedpacketRecieve)
	router.GET("/recievelist", getRedpktRecieveList)
	router.GET("/redpkt_send", getRedpacketSend)
	router.GET("/sendlist", getRedpktSendList)
	router.GET("/ready", getTobeReleased)
}

// @Title paytest
// @Description 红包付款测试
// @Accept  json
// @Param   id  query  string  true   "红包id"
// @Param   money  query  string  true   "支付金额"
// @Success 200 {string} string "付款成功"
// @Failure 400 {string} string "付款失败"
// @Resource /redpacket
// @Router /redpacket/paytest [get]
func testPayRedpacket(c *gin.Context) {
	id := c.Query("id")
	//money := c.Query("money")
	ret := apiservice.PayRedpacketBy3rdParty(id, 200000000)
	if ret {
		c.JSON(200, "testPayRedpacket ok")
	} else {
		c.JSON(400, "testPayRedpacket fail")
	}
}

// @Title redpacket
// @Description 获取红包列表
// @Accept  json
// @Param   startidx  query  uint32  true   "查询开始索引"
// @Success 200 {object} models.S2C_RedpacketList "获取红包列表成功"
// @Failure 400 {object} models.APIError "获取红包列表失败"
// @Resource /redpacket
// @Router /redpacket/list [get]
func getRedpacketList(c *gin.Context) {
	index := c.Query("startidx")
	retList := apiservice.GetRedpacketList(index, c.Request)
	if retList != nil {
		c.JSON(200, retList)
	} else {
		c.JSON(400, models.APIError{10301, "getRedpacketList fail"})
	}
}

// @Title record
// @Description 抢红包的记录
// @Accept  json
// @Param   redpacketid  query   string  true   "红包id"
// @Param   cursor  query   int  true   "游标"
// @Success 200 {object} models.S2C_RedpktRecord "领取红包记录"
// @Resource /redpacket
// @Router /redpacket/record [get]
func getRecordList(c *gin.Context) {
	redpacketId := c.Query("redpacketid")
	cursor := c.Query("cursor")
	list := apiservice.GetRedpktRecordList(redpacketId, cursor)
	c.JSON(200, list)
}

// @Title verify
// @Description 审核红包[app和游戏红包]
// @Accept  json
// @Param   redpacketid  query   string  true   "红包id"
// @Param   status  query   int  true   "审核结果"
// @Success 200 {string} string "操作成功"
// @Resource /redpacket
// @Router /redpacket/verify [get]
func verifyAppRedpacket(c *gin.Context) {
	redpacketId := c.Query("redpacketid")
	status := c.Query("status")
	apiservice.VerifyAppRedpacket(redpacketId, status)
	c.JSON(200, "verifyAppRedpacket ok")
}

// @Title sendredpacket
// @Description 发红包
// @Accept  json
// @Param   data  body   models.SendRedpacketBinding  true   "红包信息"
// @Success 201 {string} string "发红包成功"
// @Failure 400 {object} models.APIError "发红包失败"
// @Resource /redpacket
// @Router /redpacket/send [post]
func sendRedpacket(c *gin.Context) {
	redpktId, err := apiservice.CreateRedpacket(c)

	if err == nil {
		c.JSON(201, gin.H{"redpacketId": redpktId})
	} else {
		c.JSON(400, models.APIError{10201, err.Error()})
	}
}

// @Title grabredpacket
// @Description 抢红包
// @Accept  json
// @Param   redpacketid  query   string  true   "红包id"
// @Param   deviceid  query   string  true   "设备id"
// @Success 200 {string} string "抢红包成功"
// @Failure 400 {object} models.APIError "抢红包失败"
// @Resource /redpacket
// @Router /redpacket/grab [get]
func grabRedpacket(c *gin.Context) {
	redpacketId := c.Query("redpacketid")
	userId := c.MustGet("userId").(string)
	deviceid := c.Query("deviceid")

	ecode := apiservice.GrabRedpacket(userId, redpacketId, deviceid)
	if ecode == 0 {
		c.JSON(200, "grabRedpacket success")
	} else {
		c.JSON(400, models.APIError{ecode, "grabRedpacket fail"})
	}
}

// @Title doingtask
// @Description 完成红包分享任务，等待6小时后截图
// @Accept  json
// @Param   redpacketid  query   string  true   "红包id"
// @Param   deviceid  query   string  true   "设备id"
// @Success 200 {string} string "完成分享"
// @Failure 400 {object} models.APIError "失败"
// @Resource /redpacket
// @Router /redpacket/share [get]
func finishRedpacketShare(c *gin.Context) {
	redpacketId := c.Query("redpacketid")
	deviceid := c.Query("deviceid")
	userId := c.MustGet("userId").(string)

	ecode := apiservice.FinishShare(userId, redpacketId, deviceid)
	if ecode == 0 {
		c.JSON(200, "finishRedpacketShare success")
	} else {
		c.JSON(400, models.APIError{ecode, "grabRedpacket fail"})
	}
}

// @Title finishredpacket
// @Description 完成红包任务
// @Accept  json
// @Param   redpacketid  query   string  true   "红包id"
// @Param   deviceid  query   string  true   "设备id"
// @Success 200 {string} string "完成红包任务，并获得红包"
// @Failure 400 {object} models.APIError "完成失败"
// @Resource /redpacket
// @Router /redpacket/finish [get]
func finishRedpacket(c *gin.Context) {
	redpacketId := c.Query("redpacketid")
	deviceid := c.Query("deviceid")
	userId := c.MustGet("userId").(string)

	ecode := apiservice.FinishRedpacket(userId, redpacketId, deviceid)
	if ecode == 0 {
		c.JSON(200, "finishRedpacket success")
	} else {
		c.JSON(400, models.APIError{ecode, "finishRedpacket fail"})
	}
}

// @Title giveupredpacket
// @Description 放弃红包
// @Accept  json
// @Param   redpacketid  query   string  true   "红包id"
// @Success 200 {string} string "放弃红包"
// @Resource /redpacket
// @Router /redpacket/giveup [get]
func giveupRedpacket(c *gin.Context) {
	redpacketId := c.Query("redpacketid")
	userId := c.MustGet("userId").(string)

	apiservice.GiveupRedpacket(userId, redpacketId)
	c.JSON(200, "giveupRedpacket success")
}

// @Title paybybalance
// @Description 通过余额支付红包
// @Accept  json
// @Param   redpacketid  form   string  true   "红包id"
// @Param   password  form   string  true   "支付密码"
// @Success 200 {string} string "支付成功"
// @Failure 400 {object} models.APIError "支付失败"
// @Resource /redpacket
// @Router /redpacket/pay [post]
func paymentByBalance(c *gin.Context) {
	redpacketId := c.PostForm("redpacketid")
	//password := c.PostForm("password")
	userId := c.MustGet("userId").(string)

	ok := apiservice.PayRedpacketByBalance(userId, redpacketId)
	if ok {
		c.JSON(200, "paymentByBalance success")
	} else {
		c.JSON(400, models.APIError{10258, "paymentByBalance fail"})
	}
}

// @Title filter
// @Description 根据用户过滤红包
// @Accept  json
// @Success 200 {object} models.UserExpireList "根据玩家过滤红包池中的红包"
// @Failure 400 {object} models.APIError "过滤失败"
// @Resource /redpacket
// @Router /redpacket/filter [get]
func filterByUser(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	list := apiservice.FilterRedpacket(userId)
	if list != nil {
		c.JSON(200, list)
	} else {
		c.JSON(400, models.APIError{10004, "filter error"})
	}
}

// @Title Statistics
// @Description 获取红包的统计数据
// @Accept  json
// @Param   redpacketid  form   string  true   "红包id"
// @Success 200 {object} models.S2C_RedpktStatistics "红包统计数据"
// @Failure 400 {object} models.APIError "获取失败"
// @Resource /redpacket
// @Router /redpacket/statistics [get]
func getRedpacketStatistics(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	redpktId := c.Query("redpacketid")

	list := apiservice.GetRedpacketStatistics(userId, redpktId)
	if list != nil {
		c.JSON(200, list)
	} else {
		c.JSON(400, models.APIError{10004, "getRedpacketStatistics error"})
	}
}

// @Title redpacketRecieve
// @Description 收到的红包
// @Accept  json
// @Success 200 {object} models.S2C_RedpketRecieveInfo "成功"
// @Resource /redpacket
// @Router /redpacket/redpkt_recieve [get]
func getRedpacketRecieve(c *gin.Context) {
	userId := c.MustGet("userId").(string)

	ret := apiservice.GetRedpketRecieveInfo(userId)
	c.JSON(200, ret)
}

// @Title RecieveList
// @Description 收到的红包记录
// @Accept  json
// @Param   date  query   string  true   "日期,（格式 20060102）"
// @Success 200 {object} models.S2C_ReceivedList "成功"
// @Failure 400 {object} models.APIError "失败"
// @Resource /redpacket
// @Router /redpacket/recievelist [get]
func getRedpktRecieveList(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	date := c.Query("date")
	ret := apiservice.GetRedpketRecieveList(userId, date)
	if ret == nil {
		c.JSON(200, ret)
	} else {
		c.JSON(400, models.APIError{10004, "getRedpktRecieveList fail"})
	}
}

// @Title redpacketSend
// @Description 发出的红包
// @Accept  json
// @Param   year  query   string  true   "年份"
// @Success 200 {object} models.S2C_RedpketSendInfo "成功"
// @Failure 400 {string} string "失败"
// @Resource /redpacket
// @Router /redpacket/redpkt_send [get]
func getRedpacketSend(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	year := c.Query("year")

	if len(year) == 0 {
		c.JSON(200, "[Error]getRedpacketSend fail")
	} else {
		ret := apiservice.GetRedpketSendInfo(userId, year)
		c.JSON(200, ret)
	}
}

// @Title redpacketSend
// @Description 发出的红包记录
// @Accept  json
// @Param   year  query   string  true   "年份"
// @Param   cursor  query   string  true   "分页游标(0开始)"
// @Success 200 {object} models.S2C_RedpktSendList "成功"
// @Failure 400 {object} models.APIError "失败"
// @Resource /redpacket
// @Router /redpacket/sendlist [get]
func getRedpktSendList(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	year := c.Query("year")

	ret := apiservice.GetRedpketSendList(userId, year)
	if ret != nil {
		c.JSON(200, ret)
	} else {
		c.JSON(400, models.APIError{10004, "getRedpktSendList fail"})
	}
}

// @Title tobereleased
// @Description 待发出的红包
// @Accept  json
// @Success 200 {object} models.S2C_ToBeReleasedList "成功"
// @Failure 400 {object} models.APIError "失败"
// @Resource /redpacket
// @Router /redpacket/ready [get]
func getTobeReleased(c *gin.Context) {
	userId := c.MustGet("userId").(string)

	ret := apiservice.GetTobeReleasedRedpktList(userId)
	if ret != nil {
		c.JSON(200, ret)
	} else {
		c.JSON(400, models.APIError{10004, "getRedpktSendList fail"})
	}
}
