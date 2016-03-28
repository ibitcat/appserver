package pkg

import (
	"app-server/pkg/mongodb"
	"app-server/pkg/redis"
	"app-server/pkg/sdk/alipay"
	"app-server/pkg/sdk/jpush"
	"app-server/pkg/sdk/qiniu"
	"app-server/pkg/sdk/rongcloud"
	"app-server/pkg/sdk/unionpay"
	"app-server/pkg/token"
)

// 初始化服务器组件和sdk
func Init() {
	mongodb.InitMasterSession()
	token.InitJwtBackend()
	redis.NewRedisCache()

	// sdk
	if !rongcloud.InitRcSDK() {
		panic("融云sdk初始化失败……")
	}

	if !qiniu.InitQiniuSdk() {
		panic("七牛sdk初始化失败……")
	}

	alipay.InitAlipaySDK()
	unionpay.InitUnionPaySdk()
	jpush.InitJpushSDK()
}
