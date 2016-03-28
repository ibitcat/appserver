// General API info, API 通用信息，在项目入口函数所在文件写一份即可

// @APIVersion 1.0.0
// @Apititle AppServer API
// @Apidescription appserver swagger 文档测试
// @Basepath http://192.168.1.106:8080/v1

package api

import (
	"fmt"

	"app-server/api/v1"
	"app-server/pkg/middleware"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// 注册路由表
func SetupRouters(router *gin.Engine) {
	setupSwagger(router)

	api_v1 := router.Group("/v1")
	{
		api_v1.Use(middleware.Graylog(false))
		//api_v1.Use(middleware.InspectMiddleware()) // 注意：中间件有顺序限制，即中间件只能作用于use之后的router

		v1.Users(api_v1)
		v1.Oauth(api_v1)
		v1.Alipay(api_v1)
		v1.Unionpay(api_v1)
		v1.QiniuCloud(api_v1)
		v1.Redpacket(api_v1)
		v1.ScanningRedpkt(api_v1)
		v1.Socially(api_v1)
		v1.Public(api_v1)
		v1.Wechatpay(api_v1)
		v1.BackPay(api_v1)
		v1.PayPhone(api_v1)
		v1.EnterpriseCert(api_v1)
		v1.Gm(api_v1)
	}

	// catch no router
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, "Oops,这个页面被狗吃了!")
	})
	//printRouterList(router)
}

// 集成swaggerUI
// 注意：暂时需要手动修改包名为 “api” 以及修改listenjson中的"basePath": "http://192.168.1.xxx:8080/docs",
func setupSwagger(router *gin.Engine) {
	if gin.Mode() == gin.DebugMode { // 开发模式
		// 设置静态文件路径
		//router.Static("/css", "swagger/css") //参数1=url文件路径 参数2=静态文件的物理路径
		router.Use(static.Serve("/", static.LocalFile("swagger", false)))
		router.LoadHTMLGlob("swagger/index.html")

		router.GET("/", func(c *gin.Context) {
			c.HTML(200, "index.html", nil)
		})

		router.GET("/docs", func(c *gin.Context) {
			c.String(200, resourceListingJson)
		})

		router.GET("/docs/:apiKey", func(c *gin.Context) {
			apiKey := c.Param("apiKey")
			if json, ok := apiDescriptionsJson[apiKey]; ok {
				c.String(200, json)
			} else {
				c.String(404, "404！page not found！")
			}
		})
	} else if gin.Mode() == gin.ReleaseMode { // 生产模式
		router.GET("/", func(c *gin.Context) {
			c.String(200, "welcome my home!")
		})
	}
}

// router 列表
func printRouterList(router *gin.Engine) {
	fmt.Println("\n【注册的路由列表】：")
	routerList := router.Routes()
	for _, v := range routerList {
		fmt.Println("router handle = ", v)
	}
}
