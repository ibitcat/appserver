//抢红包系统

package logic

import (
	"fmt"
	"sort"
	"time"

	"app-server/config"
	"app-server/define"
	"app-server/models"
	"app-server/pkg/mongodb"
	"app-server/pkg/redis"
	"app-server/pkg/utils"

	"gopkg.in/mgo.v2/bson"
)

var RedpacketContainers map[string]IRedpacket

func InitRedpacket() {
	RedpacketContainers = make(map[string]IRedpacket)
	loadRedpacket()
}

/////////////////////////////////////////////////////////
// private
// 从mongodb中读取红包信息
func loadRedpacket() {
	if _, err := redis.Do("DEL", define.RedpktListKey); err != nil { // 清空列表
		return
	}

	var redpkts []*models.RedPacket
	find := bson.M{
		"remainder":    bson.M{"$gt": 0},
		"trade_status": 1,
		"verify":       1,
	}
	if mongodb.SelectAll(define.RedpacketCollection, find, nil, &redpkts) != nil {
		panic("读取红包列表出错……")
		return
	}

	fmt.Println("今天要发的红包个数 = ", len(redpkts))

	for _, v := range redpkts {
		insertCon(v)
	}
}

func insertCon(v *models.RedPacket) {
	id := v.Id_.Hex()
	if _, ok := RedpacketContainers[id]; ok {
		return
	}

	var con IRedpacket
	switch v.Type {
	case define.ERedpkt_Share, define.ERedpkt_OA: //需要分享截图
		con = NewShareRedpacket(v)
	case define.ERedpkt_App, define.ERedpkt_Game: //导量
		con = NewAppRedpacket(v)
	default:
		return
	}

	if con != nil {
		redis.Do("ZADD", define.RedpktListKey, v.BeginTime, id)
		fmt.Println("红包池中的红包id = ", id)
		RedpacketContainers[id] = con
	}
}

func deleteCon(id string) {
	if _, ok := RedpacketContainers[id]; ok {
		delete(RedpacketContainers, id)
	}
}

// 根据红包id查询红包
func findRedpacketById(redpktId string) *models.RedPacket {
	redpkt := new(models.RedPacket)
	id := bson.ObjectIdHex(redpktId)
	err := mongodb.SelectById(define.RedpacketCollection, id, nil, &redpkt)
	if err != nil {
		fmt.Println("该红包不存在……")
		return nil
	}

	return redpkt
}

// 检查能否用账户余额支付
func checkPay(redpkt *models.RedPacket, money int64) bool {
	if redpkt.TradeStatus == 1 { // 是否已经支付完成
		fmt.Println("红包已经支付完成……")
		return false
	}

	totalMoney := int(redpkt.PerMoney) * redpkt.Total
	if money < int64(totalMoney) {
		fmt.Println("支付金额不足……")
		return false
	}

	return true
}

// 红包上架
func redpacketOnsale(redpkt *models.RedPacket) {
	if redpkt.TradeStatus == 1 && redpkt.Verify == 1 { //审核通过并且付款了
		totalMoney := int(redpkt.PerMoney) * redpkt.Total
		insertCon(redpkt)                          // 上架
		updateOutcome(redpkt.SenderId, totalMoney) // 更新用户发出的红包
	}
}

/////////////////////////////////////////////////////////
// public
// 获取红包的实际单价
func GetRedpacketPerMoney(redpacketId string) int {
	if con, ok := RedpacketContainers[redpacketId]; ok {
		redpkt := con.GetRedpacketData()
		if redpkt != nil {
			percent := float32(100-redpkt.Rebate) / 100.0
			return int(float32(redpkt.PerMoney) * percent)
		}
	}
	return 0
}

func GetRedpacketCon(redpktId string) IRedpacket {
	if con, ok := RedpacketContainers[redpktId]; ok {
		return con
	}
	return nil
}

// 获取红包列表
func GetRedpktIdList(startIdx uint32) []string {
	//每次获取10个
	start := startIdx*10 - 1
	if startIdx == 0 {
		start = 0
	}
	stop := start + 10
	idList, _ := redis.GetStrings("ZRANGE", define.RedpktListKey, start, stop)
	return idList
}

// 生成红包，等待付款后上架
func CreateRedpacket(userId string, binding *models.SendRedpacketBinding) (string, error) {
	now := time.Now().Unix()
	m := models.RedPacket{}
	m.Id_ = bson.NewObjectId()
	m.SenderId = userId
	m.CreateTime = now
	m.Verify = 0
	m.BeginTime = binding.BeginTime
	m.EndTime = 0
	m.Rebate = 25
	m.PerMoney = binding.PerMoney
	m.TradeStatus = 0 // 还没付款
	m.Total = binding.TotalNum
	m.Remainder = binding.TotalNum
	m.Invoice = binding.Invoice
	m.Type = binding.Type
	m.Area = binding.Area
	m.Address = binding.Address
	switch m.Type {
	case define.ERedpkt_Share:
		m.Verify = 1 // 审核时间
		m.Share = binding.Share
	case define.ERedpkt_App, define.ERedpkt_Game:
		m.App = binding.App
	case define.ERedpkt_OA:
		m.Verify = 1
		m.OfficialAcc = binding.OfficialAcc
	}

	// 入库
	err := mongodb.Insert(define.RedpacketCollection, &m)
	if err != nil {
		return "", err
	}

	return m.Id_.Hex(), nil
}

// 测试第三方支付
func PayRedpacketBy3rdParty(redpktId string, fee int) bool {
	redpkt := findRedpacketById(redpktId)
	if redpkt == nil { // 红包不存在
		return false
	}
	fmt.Println("需要支付的金额 = ", int(redpkt.PerMoney)*redpkt.Total)

	if !checkPay(redpkt, int64(fee)) {
		return false
	}

	update := bson.M{"$set": bson.M{"trade_status": 1}}
	err := mongodb.UpdateById(define.RedpacketCollection, redpkt.Id_, update)
	if err != nil { //更新付款状态失败，钱回到红包账户
		totalMoney := int(redpkt.PerMoney) * redpkt.Total
		UpdateMoney(redpkt.SenderId, int64(totalMoney))
		return false
	}

	redpkt.TradeStatus = 1
	redpacketOnsale(redpkt) // 尝试上架

	return true
}

// 余额支付红包
func PayRedpacketByBalance(userId, redpktId string) bool {
	user, getErr := GetUserData(userId)
	if getErr != nil {
		fmt.Println(1)
		return false
	}

	redpkt := findRedpacketById(redpktId)
	if redpkt == nil { // 红包不存在
		fmt.Println(2)
		return false
	}

	if !checkPay(redpkt, user.Money) {
		fmt.Println(3)
		return false
	}

	update := bson.M{"$set": bson.M{"trade_status": 1}}
	err := mongodb.UpdateById(define.RedpacketCollection, redpkt.Id_, update)
	if err == nil { //扣钱
		redpkt.TradeStatus = 1
		totalMoney := int(redpkt.PerMoney) * redpkt.Total
		UpdateMoney(userId, -int64(totalMoney))
	}

	redpacketOnsale(redpkt) // 尝试上架
	return true
}

// 更新红包的审核状态
func UpdateVerify(redpktId string, status int) bool {
	redpkt := findRedpacketById(redpktId)
	if redpkt == nil { // 红包不存在
		return false
	}

	if redpkt.Verify == 1 { // 已经审核通过
		return false
	}

	redpkt.Verify = status
	update := bson.M{"$set": bson.M{"verify": status}}
	err := mongodb.UpdateById(define.RedpacketCollection, redpkt.Id_, update)
	if err == nil {
		if status == 2 { //审核失败,
			if redpkt.TradeStatus == 1 { //退钱
				totalMoney := int(redpkt.PerMoney) * redpkt.Total
				UpdateMoney(redpkt.SenderId, int64(totalMoney))
			}
		} else if status == 1 {
			redpacketOnsale(redpkt) // 尝试上架
		}
	}

	return true
}

// 抢红包
func GrabRedpacket(userId, redpacketId, deviceId string) uint32 {
	con, ok := RedpacketContainers[redpacketId]
	if ok {
		// 次数判断
		grab := GetDailyGrab(userId)
		lv := GetUserLevel(userId)
		if grab < 0 || lv < 0 {
			return 10253
		}

		if grab >= config.GetGrabLimit(lv) {
			return 10255
		}

		tp := con.GetRedpktType()
		own := config.GetOwnLimitByType(tp)
		fmt.Println("每类红包最多拥有的次数配置 = ", own)
		if GetUserOwnedCount(userId, tp) >= own {
			return 10267
		}

		ecode := con.Grab(userId, deviceId)
		if ecode == 0 { // 更新次数
			UpdateDailyGrab(userId)
		}
		return ecode
	}

	return 10251
}

// 完成分享等待截图确认（红包任务进行中，分享类需要）
func DoingRedpacket(userId, redpacketId, deviceId string) uint32 {
	con, ok := RedpacketContainers[redpacketId]
	if ok {
		return con.Doing(userId, deviceId)
	}

	return 10251
}

// 完成红包任务
func FinishRedpacket(userId, redpacketId, deviceId string) uint32 {
	con, ok := RedpacketContainers[redpacketId]
	if ok {
		return con.Finish(userId, deviceId)
	}

	return 10251
}

// 用户主动放弃红包
func GiveupRedpacket(userId, redpacketId string) {
	con, ok := RedpacketContainers[redpacketId]
	if ok {
		con.GiveUp(userId)
	}
}

// 删选器
// 筛选出用户以及操作的红包
func Filter(userId string) *models.UserExpireList {
	now := time.Now().Unix()
	idList, err := redis.GetStrings("ZRANGEBYSCORE", define.RedpktListKey, "-inf", now)
	if err != nil {
		fmt.Println("删选器 ", err)
		return nil
	}

	res := new(models.UserExpireList)
	res.List = make([]*models.RedpacketExpire, 0, 10)
	for _, id := range idList {
		con, ok := RedpacketContainers[id]
		if !ok {
			continue
		}

		if !con.IsStart() {
			continue
		}

		exp := con.GetExpires(userId)
		if exp != nil {
			res.List = append(res.List, exp)
		}
	}

	return res
}

// 红包抢夺记录
func ScanRecordList(redpacketId string, cursor int) map[string]int64 {
	con, ok := RedpacketContainers[redpacketId]
	if !ok { // 查询数据库
		var list []*models.GrabRecord

		selector := bson.M{"redpacket_id": redpacketId}
		sort := "-grab_time" // 按时间倒序
		fields := bson.M{"_id": 0, "grab_money": 0}
		err := mongodb.SelectAllWithParam(define.GrabRecordCollection, selector, sort, fields, 0, 30, &list)
		if err == nil {
			record := make(map[string]int64, 30)
			for _, r := range list {
				record[r.UserId] = r.GrabTime
			}
			return record
		}

		return nil
	}

	if !con.IsStart() {
		return nil
	}

	return con.ScanRecord(cursor)
}

// 用户当前拥有该类红包的次数
func GetUserOwnedCount(userId string, redpktType int) int {
	now := time.Now().Unix()
	idList, err := redis.GetStrings("ZRANGEBYSCORE", define.RedpktListKey, "-inf", now)
	if err != nil {
		return -1
	}

	var count int
	for _, id := range idList {
		con, ok := RedpacketContainers[id]
		if !ok || !con.IsStart() || !con.IsGrabed(userId) {
			continue
		}
		if con.GetRedpktType() == redpktType {
			count++
		}
	}
	fmt.Println("[该类红包已经拥有的次数(正在进行中)] = ", count, redpktType)
	return count
}

func GetStatistics(redpacketId string) *models.RedpktStatistics {
	var dbdata models.RedpktStatistics
	id := bson.ObjectIdHex(redpacketId)
	fields := bson.M{"statistics": 1}
	err := mongodb.SelectById(define.RedpacketCollection, id, fields, &dbdata)
	if err != nil {
		//return nil
	}

	dbdata.Area = map[string]int{"beijing": 11, "gd": 2, "hn": 8, "hb": 3}

	sortImp := make([][]interface{}, 0, len(dbdata.Area))
	for k, v := range dbdata.Area {
		sortImp = append(sortImp, []interface{}{k, v})
	}
	sort.Sort(utils.SortSlice{sortImp, 1})
	fmt.Println(sortImp)

	// 统计前十地区
	temp := make(map[string]int)
	for k, v := range sortImp {
		if k < 9 {
			temp[v[0].(string)] = v[1].(int)
		} else {
			temp[v[0].(string)] += v[1].(int)
		}
	}

	// 新的统计
	var statistics models.RedpktStatistics
	statistics.Sex = dbdata.Sex
	statistics.Area = temp
	statistics.Count = dbdata.Count

	return &statistics
}
