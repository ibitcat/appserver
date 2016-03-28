// 七牛云存储api

/*
需要存储的图片分为几个部分：
1、用户头像 （前缀"portrait/图片md5.jpg"）
2、红包图片 （前缀"redpacket/图片md5.jpg"）
3、晒图图片 （前缀"blueprint/图片md5.jpg"）
*/

package v1

import (
	"app-server/apiservice"
	"app-server/models"
	"app-server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// @SubApi 七牛云存储 [/qiniu]
func QiniuCloud(parentRoute *gin.RouterGroup) {
	router := parentRoute.Group("/qiniu")

	router.Use(middleware.JwtAuthMiddleware())
	router.GET("/uploadtoken", getQiniuUploadToken)
	router.GET("/privateuploadtoken", getQiniuPrivateUploadToken)
	router.GET("/downloadtoken", getQiniuDownloadToken)
	router.POST("/uptest", qiniuUploadTest)
}

// @Title uploadtoken
// @Description 获取七牛上传凭证
// @Accept  json
// @Param   key  string   string  true   "key"
// @Success 200 {object} models.S2C_QiniuUpToken "获取七牛上传凭证成功"
// @Failure 400 {object} models.APIError "获取七牛上传凭证失败"
// @Resource /qiniu
// @Router /qiniu/uploadtoken [get]
func getQiniuUploadToken(c *gin.Context) {
	key := c.Query("key")
	tokenStr := apiservice.MakeQiniuUploadToken(key)
	if len(tokenStr) > 0 {
		c.JSON(200, models.S2C_QiniuUpToken{tokenStr, 3600})
	} else {
		c.JSON(400, models.APIError{10019, "getQiniuUploadToken fail"})
	}
}

// @Title privateuploadtoken
// @Description 获取七牛私人空间上传凭证
// @Accept  json
// @Param   key  string   string  true   "key"
// @Success 200 {object} models.S2C_QiniuUpToken "获取七牛上传凭证成功"
// @Failure 400 {object} models.APIError "获取七牛上传凭证失败"
// @Resource /qiniu
// @Router /qiniu/privateuploadtoken [get]
func getQiniuPrivateUploadToken(c *gin.Context) {
	key := c.Query("key")
	tokenStr := apiservice.MakeQiniuPrivateUploadToken(key)
	if len(tokenStr) > 0 {
		c.JSON(200, models.S2C_QiniuUpToken{tokenStr, 3600})
	} else {
		c.JSON(400, models.APIError{10019, "getQiniuUploadToken fail"})
	}
}

// @Title downloadtoken
// @Description 获取七牛下载凭证(私有空间才需要)
// @Accept  json
// @Param   key  query   string  true   "图片的key"
// @Success 200 {object} models.S2C_QiniuDlUrl "获取七牛下载凭证成功"
// @Failure 400 {object} models.APIError "获取七牛下载凭证失败"
// @Resource /qiniu
// @Router /qiniu/downloadtoken [get]
func getQiniuDownloadToken(c *gin.Context) {
	key := c.Query("key")

	urlStr := apiservice.MakeQiniuPrivateUrl(key)
	if len(urlStr) > 0 {
		c.JSON(200, models.S2C_QiniuDlUrl{urlStr, 3600})
	} else {
		c.JSON(400, models.APIError{10019, "getQiniuDownloadToken fail"})
	}
}

// @Title uptest
// @Description 上传测试(暂时废弃)
// @Accept  json
// @Param   token  form   string  true   "七牛上传凭证"
// @Success 201 {string} string "上传测试成功"
// @Failure 400 {string} string "上传测试失败"
// @Resource /qiniu
// @Router /qiniu/uptest [post]
func qiniuUploadTest(c *gin.Context) {
	//tokenStr := c.PostForm("token")
	//apiservice.TestQiniuUpload(tokenStr)

	c.JSON(201, "qiniuUploadTest")
}
