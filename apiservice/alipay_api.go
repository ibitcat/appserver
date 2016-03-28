package apiservice

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"app-server/define"
	"app-server/models"
	"app-server/pkg/mongodb"
	"app-server/pkg/redis"
	"app-server/pkg/sdk/alipay"

	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// 支付宝异步回调处理
func AlipayNotify(req *http.Request) bool {
	ok, tradeNo, totalFee := alipay.VerifyNotify(req)
	if !ok {
		return false
	}

	feefl32, pErr := strconv.ParseFloat(totalFee, 32)
	if pErr != nil {
		return false
	}

	// 处理充值成功的逻辑
	if len(tradeNo) >= 0 {
		fmt.Println("支付宝订单号 = ", tradeNo)
		m := new(models.TradeInfo)
		m.Id_ = bson.NewObjectId()
		m.Status = 0
		m.CreatTime = time.Now().Unix()

		insertErr := mongodb.Insert(define.TradeCollection, m)
		if insertErr != nil {
			return false
		}

		// 上架红包
		totalMoney := int(feefl32 * 100.0) //支付宝要乘100(1元RMB-->100分RMB)
		PayRedpacketBy3rdParty(tradeNo, totalMoney)
	}
	return true
}

// 生成订单号
func MakeOutTradeNo(userId string) (string, error) {
	curDate := time.Now().Format("20060102")

	key := "OutTradeNo_" + curDate
	reply, err := redis.Do("INCR", key)
	if err != nil {
		return "", err
	}

	noStr := strconv.FormatInt(reply.(int64), 10)
	if len(noStr) > 0 {
		m := new(models.TradeInfo)
		m.Id_ = bson.NewObjectId()
		m.TradeNo = noStr
		m.UserId = userId
		m.Status = 0
		m.CreatTime = time.Now().Unix()

		insertErr := mongodb.Insert(define.TradeCollection, m)
		if insertErr != nil {
			return "", insertErr
		}
	}

	return noStr, nil
}

func GetAlipayParams() *models.AlipayParams {
	return &models.AlipayParams{
		AlipayPid: alipay.AlipayPid,
		AlipayAcc: alipay.AlipayEmail,
		NotifyUrl: alipay.NotifyUrl,
	}
}
