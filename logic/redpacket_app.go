// app/游戏刷量类红包

package logic

import (
	"fmt"
	"time"

	"app-server/define"
	"app-server/models"
	"app-server/pkg/redis"
)

type AppRedpacket struct {
	RedpacketBase
}

func NewAppRedpacket(redpkt *models.RedPacket) IRedpacket {
	con := new(AppRedpacket)
	con.Id = redpkt.Id_.Hex()
	con.ColdData = redpkt

	con.GrabStatusKey = define.GrabStatusPrefix + con.Id
	con.GrabKey = define.RedpktGrabRrefix + con.Id
	con.RecordKey = define.RedpktRecordPrefix + con.Id
	con.DeviceKey = define.RedpktDevicePrefix + con.Id

	con.CloseTimer = make(chan struct{})
	go con.onTimer()

	return con
}

/////////////////////////////////////////////////////////
// private
func (this *AppRedpacket) onTimer() {
	tick := time.NewTicker(time.Duration(1) * time.Second)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			now := time.Now().Unix()
			redis.GetInt64("ZREMRANGEBYSCORE", this.GrabKey, "-inf", now)
		case <-this.CloseTimer:
			return
		}
	}
}

/////////////////////////////////////////////////////////
// public
// 抢红包
func (this *AppRedpacket) Grab(userId, deviceId string) uint32 {
	switch {
	case userId == this.ColdData.SenderId: // 不能领取自己发的app红包
		return 10263
	case len(deviceId) == 0: // 检查设备
		return 10265
	case this.getRemainder() == 0: // 红包剩余个数
		return 10254
	case this.isReceive(userId) == 1: // 是否完成了红包
		return 10252
	}

	// 30分钟分享倒计时
	now := time.Now().Unix()
	exp := time.Now().Add(time.Duration(30) * time.Minute).Unix()
	args := []interface{}{
		this.GrabKey, this.GrabStatusKey, this.DeviceKey, //KEYS
		exp, userId, deviceId, now, 1, // ARGV
	}
	ecode, err := redis.DoLuaUint(define.GRedpktLua_Grab, 3, args...)
	fmt.Println("[App]Grab Redis Lua = ", ecode, err)
	if err != nil {
		return 10253
	}

	return ecode
}

// app和游戏下载类不需要
func (this *AppRedpacket) Doing(userId, deviceId string) uint32 {
	return 0
}

func (this *AppRedpacket) Finish(userId, deviceId string) uint32 {
	now := time.Now().Unix()
	args := []interface{}{
		this.GrabKey, this.RecordKey, this.DeviceKey, //KEYS
		userId, deviceId, now, this.ColdData.Total, // ARGV
	}
	ecode, err := redis.DoLuaUint(define.GRedpktLua_Finish, 3, args...)
	fmt.Println("[App]Finish Redis Lua = ", ecode, err)
	if err != nil {
		return 10269
	}

	if ecode == 0 {
		this.insertRecord(userId)
		go this.updateRemainder()             // 更新红包剩余个数到db
		go this.delDeviceAndGrabRedis(userId) // 删除用户的设备记录和抢夺状态记录
		go this.updateDistribution(userId)    // 红包下载次数+1
		go this.updateDistribution(userId)    // 红包统计数据，性别分布和地狱分布
	}

	return ecode
}

func (this *AppRedpacket) GiveUp(userId string) {
	redis.GetInt64("ZREM", this.GrabKey, userId) // 放弃红包
	go this.delDeviceAndGrabRedis(userId)
}

// 获取用户抢到该红包的到期时间
func (this *AppRedpacket) GetExpires(userId string) *models.RedpacketExpire {
	gs := this.getGrabStatus(userId)
	if gs == nil { // 从未操作过该红包
		return nil
	}

	recieved := this.isReceive(userId) // 是否领了该红包（收到钱了）
	if recieved == 0 && time.Now().Unix() >= gs.Expire {
		return nil
	}

	exp := models.RedpacketExpire{}
	exp.Id = this.Id
	exp.Type = this.ColdData.Type
	exp.IsGrab = recieved
	exp.Download = gs.Expire

	return &exp
}
