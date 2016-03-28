package apiservice

import (
	"fmt"
	"strconv"
	"time"

	"app-server/define"
	"app-server/logic"
	"app-server/models"
	"app-server/pkg/mongodb"
	"app-server/pkg/sdk/apix"

	"gopkg.in/mgo.v2/bson"
)

// 查询余额
func PayPhoneBalance() (float64, error) {
	return apix.GetPayPhoneBalance()
}

// 根据手机号和充值额度查询商品信息
func PhoneRechargeQuery(phone, price string) (*models.PhoneRechargeQueryResp, error) {
	return apix.GetPhoneRechargeProductInfo(phone, price)
}

// 充值话费
func PhoneRecharge(userId, phone, price string) (*models.PhoneRechargeResp, uint32) {
	user, err := logic.GetUserData(userId)
	if err != nil {
		return nil, 10504
	}

	info, err := apix.GetPhoneRechargeProductInfo(phone, price)
	if err != nil {
		fmt.Println("充值话费：获取产品信息失败", phone, price)
		return nil, 10504
	}
	amount := int64(info.Data.Inprice * 100)

	// 余额不足
	if user.Money < amount {
		fmt.Println("充值话费：超出剩余金额", user.Money, amount)
		return nil, 10504
	}

	// 交易id
	tradeNo := bson.NewObjectId()
	resp, err := apix.ReqPhoneRecharge(phone, price, tradeNo.Hex())
	if err != nil {
		fmt.Println("充值话费请求错误", err.Error())
		return nil, 10504
	}

	if resp.Code != 0 {
		fmt.Println("充值话费错误", resp.Code, resp.Msg)
		return nil, 10504
	}
	if resp.Data.State != "0" && resp.Data.State != "1" {
		fmt.Println("充值失败，state", resp.Data.State)
		return nil, 10504
	}

	// 扣除余额
	logic.UpdateMoney(userId, -amount)

	mongodb.Insert(define.PayPhoneRecordCollection, &models.PayPhoneRecord{
		Id_:       tradeNo,
		TradeNo:   tradeNo.Hex(),
		UserId:    userId,
		Type:      1,
		Fee:       int(amount),
		Status:    0,
		CreatTime: time.Now().Unix(),
	})

	return resp, 0
}

// 查询号码支持的流量套餐
func DataRechargeQuery(phone string) (*models.DataRechargeQueryResp, error) {
	return apix.GetDataRechargeProductInfo(phone)
}

// 充值流量
func DataRecharge(userId, phone, pkgid string) (*models.DataRechargeResp, uint32) {
	pkgidInt, err := strconv.Atoi(pkgid)
	if err != nil {
		fmt.Println("无效pkgid %s", pkgid)
		return nil, 10506
	}

	user, err := logic.GetUserData(userId)
	if err != nil {
		return nil, 10506
	}

	productInfo, err := apix.GetDataRechargeProductInfo(phone)
	if err != nil {
		fmt.Println("获取产品信息失败", phone)
		return nil, 10506
	}

	var amount int64
	for _, v := range productInfo.Data.UserDataPackages {
		if v.PkgId == pkgidInt {
			amount = int64(v.Cost * 100)
			break
		}
	}
	if amount == 0 {
		fmt.Println("无效pkgid", pkgid)
		return nil, 10506
	}

	// 余额不足
	if user.Money < amount {
		fmt.Println("超出剩余金额", user.Money, amount)
		return nil, 10506
	}

	// 交易id
	tradeNo := bson.NewObjectId()
	resp, err := apix.ReqDataRecharge(phone, pkgid, tradeNo.Hex())
	if err != nil {
		fmt.Println("充值流量请求错误", err.Error())
		return nil, 10506
	}

	if resp.Code != 0 {
		fmt.Println("充值流量错误", resp.Code, resp.Msg)
		return nil, 10506
	}
	if resp.Data.State != "0" && resp.Data.State != "1" {
		fmt.Println("充值失败，state", resp.Data.State)
		return nil, 10506
	}

	// 扣除余额
	logic.UpdateMoney(userId, -amount)

	mongodb.Insert(define.PayPhoneRecordCollection, &models.PayPhoneRecord{
		Id_:       tradeNo,
		TradeNo:   tradeNo.Hex(),
		UserId:    userId,
		Type:      2,
		Fee:       int(amount),
		Status:    0,
		CreatTime: time.Now().Unix(),
	})

	return resp, 0
}

// 充值话费回调
func PhoneRechargeNotify(state, orderid, ordertime, sign, errMsg string) error {
	if state == "0" {
		return nil
	}
	if state != "1" {
		return fmt.Errorf("state %s", state)
	}

	err := apix.PhoneRechargeNotify_VerifySign(sign, orderid, ordertime)
	if err != nil {
		return err
	}

	logic.UpdatePayPhoneRecordFinish(orderid)
	return nil
}

// 充值流量回调
func DataRechargeNotify(state, orderid, ordertime, sign, errMsg string) error {
	if state == "0" {
		return nil
	}
	if state != "1" {
		return fmt.Errorf("state %s", state)
	}

	err := apix.DataRechargeNotify_VerifySign(sign, orderid, ordertime)
	if err != nil {
		return err
	}

	logic.UpdatePayPhoneRecordFinish(orderid)
	return nil
}
