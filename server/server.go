package server

import (
	"fmt"
	"os"
	"os/signal"
	"runtime/pprof"

	"app-server/api"
	"app-server/config"
	"app-server/logic"
	"app-server/pkg"

	"github.com/gin-gonic/gin"
)

// ctlr+c 退出
func handlieSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	pprof.StopCPUProfile()
	fmt.Println("Closing server……")
	os.Exit(1)
}

// 初始化一些东西
func Init() {
	config.Init() // 初始化配置
	pkg.Init()    // 初始化各类组件（pkg）
	logic.Init()  // 初始化服务管理器
}

// 开启pprof
func startPprof() {
	f, err := os.Create("cpu.pprof")
	if err != nil {
		return
	}

	pprof.StartCPUProfile(f)
}

func Run() {
	// 注意：app初始化的panic直接抛出来，不要recover
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("app server init panic:", err)
	// 	}
	// }()

	startPprof()

	fmt.Printf("\n-------------------- 开始初始化组件、sdk、service ------------------\n")
	Init()
	fmt.Printf("----------------- 初始化OK(如果panic,后面的逻辑无法执行) ----------------\n\n")

	// 开发模式设置为debug
	gin.SetMode("release")

	// 注册路由
	fmt.Printf("\n------------------ 开始注册路由 ------------------\n")
	router := gin.Default()
	api.SetupRouters(router)
	fmt.Printf("-------------------- 注册路由OK --------------------\n\n")

	go func() {
		//err := router.RunTLS(":8080", "res/ssl/cert.pem", "res/ssl/key.pem")
		err := router.Run(":8080")
		if err != nil {
			fmt.Println("Cant start server:", err)
			os.Exit(1)
		}
	}()

	handlieSignals()
}
