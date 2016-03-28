// 红包基类

package logic

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"app-server/config"
	"app-server/define"
	"app-server/models"
	"app-server/pkg/mongodb"
	"app-server/pkg/redis"

	"gopkg.in/mgo.v2/bson"
)

type RedpacketBase struct {
	Id            string            // 红包id
	CloseTimer    chan struct{}     // 关闭定时器
	ColdData      *models.RedPacket // 红包冷数据
	GrabStatusKey string            // 用户抢红包的状态
	GrabKey       string            // 抢红包的用户队列
	RecordKey     string            // 已领取红包的用户队列
	DeviceKey     string            // 用户设备映射
}

/////////////////////////////////////////////////////////
// private
// 获取红包剩余个数
func (this *RedpacketBase) getRemainder() int {
	zcard, _ := redis.GetInt("ZCARD", this.GrabKey)
	hlen, _ := redis.GetInt("HLEN", this.RecordKey)

	r := this.ColdData.Total - (zcard + hlen)
	if r >= 0 {
		return r
	}
	return 0
}

// 红包真实剩余个数
func (this *RedpacketBase) getRealRemainder() int {
	count, err := redis.GetInt64("ZCARD", this.RecordKey)
	if err != nil {
		return 0
	}

	var r int = 0
	if int(count) < this.ColdData.Total {
		r = this.ColdData.Total - int(count)
	}

	return r
}

func (this *RedpacketBase) getExpire(userId string) int64 {
	exp, zErr := redis.GetInt64("ZSCORE", this.GrabKey, userId)
	if zErr != nil {
		return 0
	}

	return exp
}

// 获取抢红包状态
func (this *RedpacketBase) getGrabStatus(userId string) *models.GrabStatus {
	data, _ := redis.GetBytes("HGET", this.GrabStatusKey, userId)
	if len(data) > 0 {
		var gs models.GrabStatus
		err := json.Unmarshal(data, &gs)
		if err != nil {
			return nil
		}

		return &gs
	}

	return nil
}

func (this *RedpacketBase) insertRecord(userId string) {
	now := time.Now()
	m := new(models.GrabRecord)
	m.Id_ = bson.NewObjectId()
	m.UserId = userId
	m.RedpacketId = this.Id
	m.RedpacketName = "红包标题"
	m.RedpacketType = this.ColdData.Type
	m.GrabMoney = uint32(float32(this.ColdData.PerMoney) * (float32(100-this.ColdData.Rebate) / 100.0))
	m.GrabTime = now.Unix()
	m.GrabDate = now.Format("20060103")

	mongodb.Insert(define.GrabRecordCollection, m)
}

// 是否领取了红包
func (this *RedpacketBase) isReceive(userId string) int {
	_, err := redis.GetString("ZSCORE", this.RecordKey, userId)
	if err == nil {
		return 1
	}

	// 查询数据库
	find := bson.M{"userid": userId, "redpacket_id": this.Id}
	if mongodb.Exists(define.GrabRecordCollection, find) {
		return 1
	}

	return 0
}

// 删除设备状态和抢红包状态的redis记录
func (this *RedpacketBase) delDeviceAndGrabRedis(userId string) {
	gs := this.getGrabStatus(userId)
	if gs != nil {
		redis.Do("HDEL", this.DeviceKey, gs.Device)
		redis.Do("HDEL", this.GrabStatusKey, userId)
	}
}

// 更新剩余红包
func (this *RedpacketBase) updateRemainder() error {
	remainder := this.getRealRemainder()
	selector := bson.M{"_id": bson.ObjectIdHex(this.Id)}
	if remainder > 0 {
		update := bson.M{"$set": bson.M{"remainder": remainder}}
		err := mongodb.Update(define.RedpacketCollection, selector, update)
		return err
	} else if remainder == 0 { // 红包抢完，结束
		now := time.Now().Unix()
		update := bson.M{"$set": bson.M{"remainder": remainder, "end_time": now}}
		err := mongodb.Update(define.RedpacketCollection, selector, update)
		if err == nil {
			this.closeRedpacket()
		}
		return err
	}

	return nil
}

// 下载或转发次数+1
func (this *RedpacketBase) updateStatisticsCount() {
	selector := bson.M{"_id": bson.ObjectIdHex(this.Id)}
	update := bson.M{"$inc": bson.M{"statistics.count": 1}}
	mongodb.Update(define.RedpacketCollection, selector, update)
}

// 更新红包的分布情况，性别和区域
func (this *RedpacketBase) updateDistribution(userId string) {
	values, err := getUserFields(userId, define.EUser_Sex, define.EUser_Area)
	if err != nil {
		return
	}

	sex, _ := strconv.ParseUint(values[0], 10, 8)
	var area models.AreaInfo
	json.Unmarshal([]byte(values[1]), &area)

	selector := bson.M{"_id": bson.ObjectIdHex(this.Id)}
	sexKey := fmt.Sprintf("statistics.sex.%d", uint8(sex))
	areaKey := fmt.Sprintf("statistics.area.%s", config.GetProvinceName(area.Province))
	update := bson.M{"$inc": bson.M{sexKey: 1, areaKey: 1}}
	mongodb.Update(define.RedpacketCollection, selector, update)
}

// 关闭红包
func (this *RedpacketBase) closeRedpacket() {
	fmt.Println("---------- 关闭红包 ---------- id = ", this.Id)
	// 关闭定时器，注意：不要在close后立刻将chan赋值为nil
	close(this.CloseTimer)

	// 清除redis缓存数据
	redis.Do("DEL", this.RecordKey) // 删除record
	redis.Do("DEL", this.GrabKey)   // 删除grab
	redis.Do("DEL", this.DeviceKey) // 删除device

	deleteCon(this.Id)
}

/////////////////////////////////////////////////////////
// public
func (this *RedpacketBase) IsStart() bool {
	return time.Now().Unix() >= this.ColdData.BeginTime
}

func (this *RedpacketBase) GetRedpacketData() *models.RedPacket {
	this.ColdData.Remainder = this.getRemainder()
	return this.ColdData
}

func (this *RedpacketBase) GetRedpktType() int {
	return this.ColdData.Type
}

func (this *RedpacketBase) ZRevRange(key string, start, stop int, record map[string]int64) {
	reply, _ := redis.GetStrings("ZREVRANGE", key, start, stop, "WITHSCORES")
	zrangeLen := len(reply)
	if zrangeLen == 0 || zrangeLen%2 != 0 { // 返回值为空
		return
	}

	for i := 0; i < len(reply); i += 2 {
		userId := reply[i]
		tmStr := reply[i+1]

		tm, e := strconv.ParseInt(tmStr, 10, 64)
		if e != nil {
			continue
		}
		record[userId] = tm
	}
}

func (this *RedpacketBase) ScanRecord(cursor int) map[string]int64 {
	start := cursor*10 - 1
	stop := start + 10
	if start < 0 {
		start = 0
	}

	record := make(map[string]int64, 10)
	zCardLen, _ := redis.GetInt("ZCARD", this.RecordKey)
	if zCardLen > stop { //全部在完成列表中
		this.ZRevRange(this.RecordKey, start, stop, record)
	} else if zCardLen > start && zCardLen < stop { // 一部分在完成列表，一部分在正在进行中
		this.ZRevRange(this.RecordKey, start, zCardLen-1, record)
		temp := stop - zCardLen
		this.ZRevRange(this.GrabKey, 0, temp, record)
	} else { // 全部在进行中列表
		temp := start - zCardLen
		this.ZRevRange(this.GrabKey, temp, temp+10-1, record)
	}

	return record
}

// 已经拥有的次数
func (this *RedpacketBase) IsGrabed(userId string) bool {
	gs := this.getGrabStatus(userId)
	if gs == nil { // 从未操作过该红包
		return false
	}

	return time.Now().Unix() < gs.Expire
}
