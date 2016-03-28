/*
扫红包管理器
*/

package logic

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"app-server/define"
	"app-server/models"
	"app-server/pkg/mongodb"
	"app-server/pkg/redis"
	"app-server/pkg/utils"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var UpdateTime int64
var CloseTime chan struct{}

func InitScanning() {
	UpdateTime = time.Now().Unix()
	CloseTime = make(chan struct{})

	go scanningTimer()
}

/////////////////////////////////////////////////////////
// private
// 定时器，每天凌晨生成红包列表，放到redis中
func scanningTimer() {
	ticker := time.NewTicker(time.Duration(3) * time.Second) // 3s定时器
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			now := time.Now()
			last := time.Unix(UpdateTime, 0)
			if now.Format("20060102") > last.Format("20060102") { //新的一天到来
				go genRedpacket() //生成当天红包列表
				go clearing()     //结算
			}
			UpdateTime = now.Unix()

		case <-CloseTime:
			fmt.Println("Closing ScanningMgr onTimer ……")
			return
		}
	}
}

// 新的一天到来的时候，生成当天的红包列表
func genRedpacket() {
	var list []*models.ScanningRedpkt
	nowDate, _ := strconv.Atoi(time.Now().Format("20060102"))

	// 正在进行的扫红包活动
	query := func(c *mgo.Collection) error {
		// stopdate >=now>=startdate
		// $lte -> <= 小于等于
		find := bson.M{
			"start_date": bson.M{"$gte": nowDate},
			"stop_date":  bson.M{"$lte": nowDate},
			"gen_date":   bson.M{"$lte": nowDate},
		}
		return c.Find(find).All(&list)
	}

	mongodb.M(define.ScanningCollection, query)

	// 生成红包队列
	for _, v := range list {
		totalMoney := v.DailyMoney
		if v.StopDate == uint32(nowDate) && v.LastdayMoney > 0 { //如果是最后一天
			totalMoney = v.LastdayMoney
		}

		number := totalMoney / 20 // 计算要发几个红包（按平均2毛钱一个红包计算）
		//this.PushToGenList(v.Id_.Hex(), totalMoney, number)
		go GenAlgorithms(v.Id_.Hex(), totalMoney, number)
	}

	//go this.GenTimer()

	// 触发红包生成定时器
	//go this.GenAlgorithms(v.Id_.Hex(), totalMoney, number)
}

/*暂时不采用这种方式
// 防止要生成红包的个数太多，使用redis list 做缓存，定时生成
func (this *ScanningMgr) PushToGenList(redpktId string, total, number uint32) {
	c := map[string]interface{}{
		"id":    redpktId,
		"total": total,
		"num":   number,
	}

	j, err := json.Marshal(c)
	if err == nil {
		redis.Do("RPUSH", "genlist", string(j))
	} else {
		fmt.Println("PushToGenList Error")
	}
}

// 生成红包的定时器
func (this *ScanningMgr) GenTimer() {
	tick := time.NewTicker(time.Duration(30) * time.Second) // 30s的定时器
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			r, err := redis.Do("EXISTS", "genlist")
			if err == nil && r.(int64) == 1 {
				v, _ := redis.Do("BLPOP", "genlist", 10) // 弹出
				vList := v.([]interface{})
				if len(vList) == 2 {
					fmt.Println("genlist 弹出 = ", v, reflect.TypeOf(v))

					var v map[string]interface{}
					json.Unmarshal(vList[1].([]byte), &v)

					go this.GenAlgorithms(string(v["id"].([]byte)), v["total"].(uint32), v["num"].(uint32))
				}
			} else { //关闭定时器
				return
			}
		}
	}
}
*/

// 生成红包列表的算法
// @total 红包总金额
// @number 红包个数
func GenAlgorithms(redpktId string, total, number uint32) bool {
	if total == 0 || number == 0 {
		return false
	}

	rand.Seed(time.Now().UnixNano()) // 随机种子
	key := "scannig_unused_" + redpktId
	reply, err := redis.Do("LLEN", key)
	if err != nil {
		return false
	}

	var genOk bool = true
	var i uint32 = 1
	for ; i < number; i++ {
		safe := (total - (number-i)*1) / (number - i) //随机安全上限
		v := utils.RandomInterval(1, int(safe))
		total = total - uint32(v)

		if _, err := redis.Do("RPUSH", key, v); err != nil { // 生成红包队列失败
			genOk = false
			break
		}
		fmt.Printf("红包[%d] = %d \n", i, float32(v)/100.0)
	}

	fmt.Println("最后一个红包 = ", float32(total)/100.0)

	if genOk {
		// 生成列表成功后，扣钱
		selector := bson.M{"_id": bson.ObjectIdHex(redpktId)}
		update := bson.M{"$inc": bson.M{"balance": -int(total)}}

		if mongodb.Update(define.ScanningCollection, selector, update) != nil {
			genOk = false
		}
	}

	if !genOk { // 扣钱错误，删除生成的队列
		if reply.(int64) > 0 {
			redis.Do("LTRIM", key, 0, reply.(int64))
		} else {
			redis.Do("DEL", key)
		}
	}

	return true
}

// 活动结束后，结算
func clearing() {
	nowDate := time.Now().Format("20060102")

	var result []*models.ScanningRedpkt
	// 结算没有发完的
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{
			"stop_date": bson.M{"$gt": nowDate},
			"balance":   bson.M{"$gt": 0},
			"isend":     0,
		}).All(&result)
	}
	err := mongodb.M(define.ScanningCollection, query)
	if err != nil {
		return
	}

	var surplus uint32 = 0
	for _, v := range result {
		key := "scannig_unused_" + v.Id_.Hex()
		reply, _ := redis.Do("LRANGE", key, 0, -1)

		var unused int = 0
		// 回收未发完的红包
		for _, vv := range reply.([]interface{}) {
			money, _ := strconv.Atoi(string(vv.([]byte)))
			unused += money
		}

		surplus += (v.Balance + uint32(unused))
	}

	fmt.Println("退钱 = ", surplus)
	// 退钱逻辑
}

/////////////////////////////////////////////////////////
// public
// 创建一个扫红包
func CreateScanning(form *models.SendScannigBinding) uint32 {
	if form.TotalMoney == 0 {
		return 10351
	}

	days := utils.IntervalDays(form.Startdate, form.Stopdate) // 计算活动天数
	if days == 0 {
		return 10352
	}

	m := new(models.ScanningRedpkt)
	m.Id_ = bson.NewObjectId()
	m.BarCode = form.Barcode
	m.ItemName = form.Itemname //商品名称
	m.Tag = form.Tag
	m.Description = form.Desc
	m.SendTime = time.Now().Unix()
	m.StartDate = form.Startdate
	m.StopDate = form.Stopdate
	m.Balance = uint32(float32(form.TotalMoney) * (1 - 0.1))
	m.TotalMoney = form.TotalMoney
	m.DailyMoney = m.Balance / days

	if form.TotalMoney%days != 0 { // 除不尽,最后一天特殊处理
		m.LastdayMoney = m.Balance - (days-1)*m.DailyMoney
	}

	err := mongodb.Insert(define.ScanningCollection, m)
	if err != nil {
		fmt.Println("插入错误 = ", err)
		return 10353
	}

	return 0
}

// 根据商品类型获取扫红包列表
func GetScanListByTag(startIdx uint32, tag string) *models.ScanningList {
	var list []*models.ScanningInfo

	// 正在进行的扫红包活动
	query := func(c *mgo.Collection) error {
		// stopdate >=now>=startdate
		nowDate, _ := strconv.Atoi(time.Now().Format("20060102"))
		find := bson.M{
			"stop_date":  bson.M{"$gte": nowDate},
			"start_date": bson.M{"$lte": nowDate},
		}
		feilds := bson.M{
			"item_name": 1,
			"balance":   1,
			"pic_urls":  1,
		}
		skip := int(startIdx) * 10
		iter := c.Find(find).Select(feilds).Sort("send_time").Skip(skip).Limit(10).Iter() // 按发布时间先后排序

		return iter.All(&list)
	}

	if mongodb.M(define.ScanningCollection, query) == nil {
		retList := new(models.ScanningList)
		retList.Count = uint32(len(list))
		retList.List = make([]*models.ScanningItem, 0, 10)
		for _, info := range list {
			var picUrl string
			if len(info.Pic) > 0 {
				picUrl = info.Pic[0]
			}

			retList.List = append(retList.List, &models.ScanningItem{
				Id:      info.Id.Hex(),
				Name:    info.Name,
				Pic:     picUrl,
				Balance: info.Balance,
			})
			fmt.Println("扫红包商品列表 = ", info)
		}

		return retList
	}

	return nil
}

func GetScanListByTag1(startIdx uint32, tag string) *models.ScanningList {
	var list []*models.ScanningInfo

	// 正在进行的扫红包活动
	// stopdate >=now>=startdate
	nowDate, _ := strconv.Atoi(time.Now().Format("20060102"))
	var find bson.M
	if len(tag) == 0 {
		find = bson.M{
			"stop_date":  bson.M{"$gte": nowDate},
			"start_date": bson.M{"$lte": nowDate},
		}
	} else {
		find = bson.M{
			"stop_date":  bson.M{"$gte": nowDate},
			"start_date": bson.M{"$lte": nowDate},
			"tag":        tag,
		}
	}

	feilds := bson.M{
		"item_name": 1,
		"balance":   1,
		"pic_urls":  1,
	}
	skip := int(startIdx) * 10

	if mongodb.SelectAllWithParam(define.ScanningCollection, find, "", feilds, skip, 10, &list) == nil {
		retList := new(models.ScanningList)
		retList.Count = uint32(len(list))
		retList.List = make([]*models.ScanningItem, 0, 10)
		for _, info := range list {
			var picUrl string
			if len(info.Pic) > 0 {
				picUrl = info.Pic[0]
			}

			retList.List = append(retList.List, &models.ScanningItem{
				Id:      info.Id.Hex(),
				Name:    info.Name,
				Pic:     picUrl,
				Balance: info.Balance,
			})
			fmt.Println("扫红包商品列表 = ", info)
		}

		return retList
	}

	return nil
}

// 扫描获得红包
func GetScanningRedpkt(redpktId, userId string) uint32 {
	// 是否有该活动
	keyUsed := "scannig_used_" + redpktId

	isExist, errUsed := redis.Do("HEXISTS", keyUsed, userId)
	if errUsed != nil || isExist.(int64) == 1 { //已经抢过红包
		return 10354
	}

	keyUnused := "scannig_unused_" + redpktId
	v, err := redis.Do("BLPOP", keyUnused, 3) // 阻塞pop，超时5s
	if err != nil || v == nil {               // 没抢到
		return 10355
	}

	redis.Do("HSET", keyUsed, userId)

	// 扫红包记录插入到mongodb
	money, _ := strconv.ParseUint(string(v.([]byte)), 10, 32)
	m := new(models.ScanRecord)
	m.Id_ = bson.NewObjectId()
	m.RedpktId = redpktId
	m.UserId = userId
	m.UserName = "用户名称"
	m.Money = uint32(money)
	m.Time = time.Now().Unix()
	mongodb.Insert("scanrecord", m)

	// 给玩家钱
	// TODO

	return 0
}
