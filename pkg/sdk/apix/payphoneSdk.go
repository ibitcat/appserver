package apix

import (
	"encoding/json"
	"fmt"

	"app-server/define"
	"app-server/models"
	"app-server/pkg/httplib"
	"app-server/pkg/utils"
)

var (
	APIX_Key_PhoneRecharge = `ea0a462880dd4de2555a8c69cfc19073` // 话费充值apix-key
	APIX_Key_DataRecharge  = `05f15ca124a94bd9670368cbc1faa41f` // 流量充值apix-key
	PhoneRechargeCallback  = define.NgrokDomain + `/v1/payphone/phone_recharge_notify`
	DataRechargeCallback   = define.NgrokDomain + `/v1/payphone/data_recharge_notify`
)

// 查询余额
func GetPayPhoneBalance() (float64, error) {
	req := httplib.Get("http://p.apix.cn/apixlife/pay/phone/user_balance")
	req.Header("accept", "application/json")
	req.Header("content-type", "application/json")
	req.Header("apix-key", APIX_Key_PhoneRecharge)

	b, err := req.Bytes()
	if err != nil {
		return 0, err
	}

	var resp struct {
		Code int
		Msg  string
		Data struct {
			UserBalance float64
			UserId      int64
		}
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return 0, err
	}
	if resp.Code != 0 {
		return 0, fmt.Errorf(resp.Msg)
	}
	return resp.Data.UserBalance, nil
}

// 根据手机号和充值额度查询商品价格(分)
func GetPhoneRechargeProductInfo(phone, price string) (*models.PhoneRechargeQueryResp, error) {
	req := httplib.Get("http://p.apix.cn/apixlife/pay/phone/recharge_query")
	req.Header("accept", "application/json")
	req.Header("content-type", "application/json")
	req.Header("apix-key", APIX_Key_PhoneRecharge)
	req.Param("phone", phone)
	req.Param("price", price)

	b, err := req.Bytes()
	if err != nil {
		return nil, err
	}

	var resp models.PhoneRechargeQueryResp
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Msg)
	}
	return &resp, nil
}

// 根据手机号和套餐查询套餐价格(分)
func GetDataRechargeProductInfo(phone string) (*models.DataRechargeQueryResp, error) {
	req := httplib.Get("http://p.apix.cn/apixlife/pay/mobile/package")
	req.Header("accept", "application/json")
	req.Header("content-type", "application/json")
	req.Header("apix-key", APIX_Key_DataRecharge)
	req.Param("phone", phone)

	b, err := req.Bytes()
	if err != nil {
		return nil, err
	}

	var resp models.DataRechargeQueryResp
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Msg)
	}
	return &resp, nil
}

// 发送充值话费请求
func ReqPhoneRecharge(phone, price, orderid string) (*models.PhoneRechargeResp, error) {
	req := httplib.Get("http://p.apix.cn/apixlife/pay/phone/phone_recharge")
	req.Header("accept", "application/json")
	req.Header("content-type", "application/json")
	req.Header("apix-key", APIX_Key_PhoneRecharge)
	req.Param("phone", phone)
	req.Param("price", price)
	req.Param("orderid", orderid)
	req.Param("callback_url", PhoneRechargeCallback)
	req.Param("sign", utils.Md5(phone+price+orderid))

	b, err := req.Bytes()
	if err != nil {
		return nil, err
	}

	resp := &models.PhoneRechargeResp{}
	err = json.Unmarshal(b, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// 发送充值流量请求
func ReqDataRecharge(phone, pkgid, orderid string) (*models.DataRechargeResp, error) {
	req := httplib.Get("http://p.apix.cn/apixlife/pay/mobile/data_recharge")
	req.Header("accept", "application/json")
	req.Header("content-type", "application/json")
	req.Header("apix-key", APIX_Key_DataRecharge)
	req.Param("phone", phone)
	req.Param("pkgid", pkgid)
	req.Param("orderid", orderid)
	req.Param("callback_url", DataRechargeCallback)
	req.Param("sign", utils.Md5(phone+pkgid+orderid))

	b, err := req.Bytes()
	if err != nil {
		return nil, err
	}

	resp := &models.DataRechargeResp{}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// 充值话费回调验签
func PhoneRechargeNotify_VerifySign(sign, orderid, ordertime string) error {
	verify := utils.Md5(APIX_Key_PhoneRecharge + orderid + ordertime)
	if verify != sign {
		return fmt.Errorf("sign failed")
	}
	return nil
}

// 充值流量回调验签
func DataRechargeNotify_VerifySign(sign, orderid, ordertime string) error {
	verify := utils.Md5(APIX_Key_DataRecharge + orderid + ordertime)
	if verify != sign {
		return fmt.Errorf("sign failed")
	}
	return nil
}
