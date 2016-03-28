// 文章/广告分享类红包

package logic

import (
	"fmt"
	"time"

	"app-server/config"
	"app-server/define"
	"app-server/models"
	"app-server/pkg/redis"
)

const (
	//screenshotDiff = 18 * 3600 // 截图确认需要6小时后
	screenshotDiff = 23*3600 + 60*55 // 测试环境，5分钟后可以上传截图
)

type ShareRedpacket struct {
	RedpacketBase
}

func NewShareRedpacket(redpkt *models.RedPacket) IRedpacket {
	con := new(ShareRedpacket)
	con.Id = redpkt.Id_.Hex()
	con.ColdData = redpkt

	con.GrabKey = define.RedpktGrabRrefix + con.Id
	con.GrabStatusKey = define.GrabStatusPrefix + con.Id
	con.RecordKey = define.RedpktRecordPrefix + con.Id
	con.DeviceKey = define.RedpktDevicePrefix + con.Id

	con.CloseTimer = make(chan struct{})
	go con.onTimer()

	return con
}

/////////////////////////////////////////////////////////
// private
func (this *ShareRedpacket) onTimer() {
	tick := time.NewTicker(time.Duration(1) * time.Second)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			if !this.IsStart() {
				continue
			}
			now := time.Now().Unix()
			redis.Do("ZREMRANGEBYSCORE", this.GrabKey, "0", now)
		case <-this.CloseTimer:
			fmt.Println("收到关闭红包的通知")
			return
		}
	}
}

// 检查每种红包每天可抢次数
func (this *ShareRedpacket) checkGrabLimit(userId string) bool {
	used := getGrabLimitByType(userId, this.ColdData.Type)
	if used < 0 {
		return false
	}

	if used >= config.GetLimitByType(this.ColdData.Type) { // 次数已满
		return false
	}

	return true
}

/////////////////////////////////////////////////////////
// public
// 抢红包
func (this *ShareRedpacket) Grab(userId, deviceId string) uint32 {
	switch {
	case !this.checkGrabLimit(userId):
		return 10266
	case len(deviceId) == 0: // 检查设备是否为空
		return 10265
	case this.getRemainder() <= 0: //剩余个数
		return 10254
	case this.isReceive(userId) == 1: //是否已经完成
		return 10252
	}

	// 5分钟分享倒计时
	now := time.Now().Unix()
	exp := time.Now().Add(time.Duration(5) * time.Minute).Unix()
	args := []interface{}{
		this.GrabKey, this.GrabStatusKey, this.DeviceKey, //KEYS
		exp, userId, deviceId, now, 1, // ARGV
	}
	ecode, err := redis.DoLuaUint(define.GRedpktLua_Grab, 3, args...)
	fmt.Println("[Share]Grab Redis Lua = ", ecode, err)
	if err != nil {
		return 10253
	}

	return ecode
}

func (this *ShareRedpacket) Doing(userId, deviceId string) uint32 {
	exp := time.Now().Add(time.Duration(24) * time.Hour).Unix() //24小时内截图
	now := time.Now().Unix()
	args := []interface{}{
		this.GrabStatusKey, this.GrabKey, this.DeviceKey, //KEYS
		userId, exp, now, deviceId, 2, // ARGV
	}
	ecode, err := redis.DoLuaUint(define.GRedpktLua_Share, 3, args...)
	fmt.Println("[Share]Share Redis Lua = ", ecode, err)
	if err != nil {
		return 10268
	}

	if ecode == 0 { // 红包任务积分
		cfg := config.GetPointCfg(define.EPoint_Redpacket)
		if cfg != nil {
			updateUserPoint(userId, cfg.Point)
		}

		// 红包的转发次数+1
		go this.updateStatisticsCount()
	}

	return ecode
}

func (this *ShareRedpacket) Finish(userId, deviceId string) uint32 {
	now := time.Now().Unix()
	tm := this.getExpire(userId)
	fmt.Println("[分享类红包截图开始时间] = ", now, tm, tm-screenshotDiff)
	if tm == 0 || now < tm-screenshotDiff { //截图未满6小时
		return 10260
	}

	// redis脚本
	args := []interface{}{
		this.GrabKey, this.RecordKey, this.DeviceKey, //KEYS
		userId, deviceId, now, this.ColdData.Total, // ARGV
	}
	ecode, err := redis.DoLuaUint(define.GRedpktLua_Finish, 3, args...)
	fmt.Println("[Share]Finish Redis Lua = ", ecode, err)
	if err != nil {
		return 10269
	}

	if ecode == 0 {
		this.insertRecord(userId)
		go updateGrabLimit(userId, this.ColdData.Type) // 更新领取次数
		go this.updateRemainder()                      // 更新红包剩余个数到db
		go this.delDeviceAndGrabRedis(userId)          // 删除用户的设备记录和抢夺状态记录
		go this.updateDistribution(userId)             // 红包统计数据，性别分布和地狱分布
	}

	return ecode
}

func (this *ShareRedpacket) GiveUp(userId string) {
	redis.GetInt("ZREM", this.GrabKey, userId) // 放弃红包
	go this.delDeviceAndGrabRedis(userId)
}

// 获取用户抢到该红包的到期时间
func (this *ShareRedpacket) GetExpires(userId string) *models.RedpacketExpire {
	gs := this.getGrabStatus(userId)
	if gs == nil { // 从未操作过该红包
		return nil
	}

	fmt.Println(gs)
	recieved := this.isReceive(userId) // 是否领该红包（收到钱了）
	if recieved == 0 && time.Now().Unix() >= gs.Expire {
		return nil
	}

	exp := models.RedpacketExpire{}
	exp.Id = this.Id
	exp.Type = this.ColdData.Type
	exp.IsGrab = recieved
	if gs.Status == 1 {
		exp.Share = gs.Expire
	} else if gs.Status == 2 {
		exp.Screenshot = gs.Expire
	}

	return &exp
}
