package logic

import (
	"app-server/models"
)

// 红包的接口
type IRedpacket interface {
	GetRedpktType() int                               // 获取红包的类型
	GetRedpacketData() *models.RedPacket              // 获取红包信息
	IsStart() bool                                    // 红包是否开始
	Grab(userId, deviceId string) uint32              // 抢红包
	Doing(userId, deviceId string) uint32             // 红包任务进行中
	Finish(userId, deviceId string) uint32            // 完成红包任务
	GiveUp(userId string)                             // 放弃红包
	GetExpires(userId string) *models.RedpacketExpire // 获取红包的到期时间信息
	ScanRecord(cursor int) map[string]int64           // 红包领取记录
	IsGrabed(userId string) bool                      // 用户是否抢到了红包
}
