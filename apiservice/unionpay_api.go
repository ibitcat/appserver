package apiservice

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"app-server/define"
	"app-server/models"
	"app-server/pkg/mongodb"
	"app-server/pkg/sdk/unionpay"

	"gopkg.in/mgo.v2/bson"
)

// 发送银联订单推送请求
func GetUnionpayTradeNo(orderId string) (string, error) {
	return unionpay.GetTradeNo(orderId, 100)
}

// 查询交易状态
// orderid即为红包id
func QueryUnionpayStatus(orderId string) error {
	fmt.Println("查询交易状态,id = ", orderId)
	return unionpay.QueryTradeStatus(orderId)
}

// 银联异步通知
func UnionpayNotify(request *http.Request) error {
	err := request.ParseForm()
	if err != nil {
		return err
	}

	params := make(map[string]string)
	for k, s := range request.Form {
		if len(s) > 0 {
			params[k] = s[0]
		}
	}

	verifyErr := unionpay.Verify(params)
	if verifyErr != nil {
		return verifyErr
	}

	redpacketId := params["orderId"] // 订单号即红包id
	txnAmt := params["txnAmt"]       // 交易金额，分为单位
	fmt.Println("orderid = ", redpacketId, "txnAmt = ", txnAmt)
	totalMoney, pErr := strconv.Atoi(txnAmt)
	if pErr != nil {
		return pErr
	}

	// 处理红包
	PayRedpacketBy3rdParty(redpacketId, totalMoney)

	// 加日志
	// TODO

	return nil
}

// 请求提现
func UnionBackpay() error {
	m := new(models.TradeInfo)
	m.Id_ = bson.NewObjectId()
	m.Status = 0
	m.CreatTime = time.Now().Unix()

	err := unionpay.WithdrawCash(m.Id_.Hex(), 1)
	if err != nil {
		return err
	}

	insertErr := mongodb.Insert(define.TradeCollection, m)
	if insertErr != nil {
		return insertErr
	}
	return nil
}

// 银联取现异步通知
func UnionBackpayNotify(request *http.Request) error {
	fmt.Println("银联取现异步通知")
	err := request.ParseForm()
	if err != nil {
		return err
	}

	params := make(map[string]string)
	for k, s := range request.Form {
		if len(s) > 0 {
			params[k] = s[0]
		}
	}

	verifyErr := unionpay.Verify(params)
	if verifyErr != nil {
		return verifyErr
	}

	fmt.Println("orderid = ", params["orderId"])

	// 处理红包
	// TODO

	// 加日志

	return nil
}
