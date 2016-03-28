/*
扫红包
*/

package v1

import (
	"app-server/apiservice"
	"app-server/models"
	"app-server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// @SubApi 扫红包 [/scanning]
func ScanningRedpkt(parentRoute *gin.RouterGroup) {
	router := parentRoute.Group("/scanning")

	// 不需要登陆的操作
	router.GET("/list", getScanningRedpktList)

	// 需要登陆的操作
	router.Use(middleware.JwtAuthMiddleware())
	router.POST("/send", sendScanningRedpacket)
	router.GET("/scan", scanRedpacket)
}

// @Title list
// @Description 获取扫红包列表
// @Accept  json
// @Param   tag  query   string  true   "商品标签"
// @Param   startidx  query  uint32  true   "查询开始索引"
// @Success 200 {object} models.ScanningList "获取红包列表成功"
// @Failure 400 {object} models.APIError "获取红包列表失败"
// @Resource /scanning
// @Router /scanning/list [get]
func getScanningRedpktList(c *gin.Context) {
	tag := c.Query("tag")
	idx := c.Query("startidx")

	list := apiservice.GetScanListByTag1(idx, tag)
	if list != nil {
		c.JSON(200, list)
	} else {
		c.JSON(400, models.APIError{10357, "getScanningRedpktList fail"})
	}
}

// @Title send
// @Description 商户发送扫红包
// @Accept  json
// @Param   data  body   models.SendScannigBinding  true   "红包信息"
// @Success 201 {string} string "获取红包列表成功"
// @Failure 400 {object} models.APIError "获取红包列表失败"
// @Resource /scanning
// @Router /scanning/send [post]
func sendScanningRedpacket(c *gin.Context) {
	ecode := apiservice.CreateScanning(c)
	if ecode == 0 {
		c.JSON(201, "getScanningRedpktList success")
	} else {
		c.JSON(400, models.APIError{ecode, "getScanningRedpktList fail"})
	}
}

// @Title scan
// @Description 用户扫红包
// @Accept  json
// @Param   redpacketid  query  string  true   "红包id"
// @Success 200 {string} string "扫红包成功"
// @Failure 400 {object} models.APIError "扫红包失败"
// @Resource /scanning
// @Router /scanning/scan [get]
func scanRedpacket(c *gin.Context) {
	redpktId := c.Query("redpacketid")
	userId := c.MustGet("userId").(string)

	ecode := apiservice.GetScanningRedpkt(redpktId, userId)
	if ecode == 0 {
		c.JSON(200, "scanRedpacket success")
	} else {
		c.JSON(400, models.APIError{ecode, "scanRedpacket fail"})
	}
}
