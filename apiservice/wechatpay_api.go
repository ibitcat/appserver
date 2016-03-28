package apiservice

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"app-server/define"
	"app-server/logic"
	"app-server/models"
	"app-server/pkg/mongodb"
	"app-server/pkg/sdk/wechatpay"
	"app-server/pkg/utils"

	"gopkg.in/mgo.v2/bson"
)

func wechatXml2map(b []byte) map[string]string {
	d := xml.NewDecoder(bytes.NewReader(b))
	ms := make(map[string]string)
	key, val := "", ""
	for {
		t, err := d.Token()
		if err != nil {
			break
		}
		switch tt := t.(type) {
		case xml.StartElement:
			key = tt.Name.Local
		case xml.EndElement:
			if len(key) > 0 && len(val) > 0 {
				ms[key] = val
			}
			key, val = "", ""
		case xml.CharData:
			s := strings.TrimSpace(string(tt))
			if len(s) > 0 {
				val = s
			}
		default:
		}
	}
	return ms
}

// 微信支付异步回调处理
func WechatpayNoify(request *http.Request) error {
	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	fmt.Printf("str: %s\n", string(b))

	params := wechatXml2map(b)
	fmt.Printf("%#v\n", params)

	// 参数说明地址
	// https://pay.weixin.qq.com/wiki/doc/api/app.php?chapter=9_7&index=3
	if params["return_code"] != "SUCCESS" {
		return fmt.Errorf(params["return_msg"])
	}

	if params["result_code"] != "SUCCESS" {
		return fmt.Errorf("%s: %s", params["err_code"], params["err_code_des"])
	}

	err = wechatpay.VerifyNotify(params)
	if err != nil {
		return err
	}

	wechatTradeNo := params["transaction_id"]
	fmt.Println("微信订单号", wechatTradeNo)
	totalFee, err := strconv.Atoi(params["total_fee"])
	if err != nil {
		fmt.Println("无效金额:", params["total_fee"])
		return err
	}

	tradeNo := params["out_trade_no"]
	if len(tradeNo) >= 0 {
		fmt.Println("微信支付红包id = ", tradeNo)
		m := new(models.TradeInfo)
		m.Id_ = bson.NewObjectId()
		m.Status = 0
		m.CreatTime = time.Now().Unix()

		err = mongodb.Insert(define.TradeCollection, m)
		if err != nil {
			return err
		}

		// 上架红包
		PayRedpacketBy3rdParty(tradeNo, totalFee)
	}

	return nil
}

// 获取微信支付参数
func GetWechatPayParams() *models.WechatPayParams {
	return &models.WechatPayParams{
		AppId:     wechatpay.WechatAppId,
		MchId:     wechatpay.WechatMchId,
		NotifyUrl: wechatpay.NotifyUrl,
		PayKey:    wechatpay.WechatpayKey,
	}
}

// 微信提现
func WechatBackPay(userId string, request *http.Request) uint32 {
	request.ParseForm()
	realName := request.FormValue("realname")
	passwd := request.FormValue("passwd")

	amount, err := strconv.ParseFloat(request.FormValue("money"), 64)
	if err != nil {
		fmt.Println("无效金额", request.FormValue("money"))
		return 10451
	}
	remoteStr := strings.Split(request.RemoteAddr, ":")
	if len(remoteStr) != 2 {
		fmt.Println("无效客户端ip", request.RemoteAddr)
		return 10451
	}

	user, err := logic.FindUserDataById(userId)
	if err != nil {
		return 10451
	}

	if utils.Md5(passwd+user.Salt) != user.Password {
		return 10452
	}

	// 余额不足
	if user.Money < int64(amount*100) {
		fmt.Println("超出剩余金额", user.Money, int64(amount*100))
		return 10451
	}

	// 交易id
	tradeNo := bson.NewObjectId()

	// 请求支付
	result, err := wechatpay.ReqBackPay(user.WeixinOpenId, tradeNo.Hex(), remoteStr[0], realName, int64(amount*100))
	if err != nil {
		fmt.Println("请求支付错误", err.Error())
		return 10451
	}

	if result.ErrCode != nil && *result.ErrCode == "NOTENOUGH" {
		// TODO:
		logic.UpdateMoney(userId, -int64(amount*100))
		logic.CreateBackpayRecord(&models.BackpayRecord{
			Id_:    tradeNo,
			UserId: userId,
			Type:   4,
			Name:   realName,
			Fee:    int64(amount * 100),
			Status: 2,
			Time:   time.Now().Unix(),
		})
		return 0
	} else if result.ReturnCode == "SUCCESS" && *result.ResultCode == "SUCCESS" {
		fmt.Println("微信提现：微信订单号", *result.PaymentNo)
		logic.UpdateMoney(userId, -int64(amount*100))
		logic.CreateBackpayRecord(&models.BackpayRecord{
			Id_:    tradeNo,
			UserId: userId,
			Type:   4,
			Name:   realName,
			Fee:    int64(amount * 100),
			Status: 1,
			Time:   time.Now().Unix(),
		})
		return 0
	} else {
		fmt.Println("微信提现：ReturnCode:", result.ReturnCode, "ErrCode:", *result.ErrCode, "Desc", *result.ErrCodeDes)
		return 10451
	}
}
