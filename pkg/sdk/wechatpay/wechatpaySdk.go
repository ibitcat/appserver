package wechatpay

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"app-server/define"
	"app-server/models"
	"app-server/pkg/utils"
)

const (
	WechatAppId = `wx4f2bab4031677dad` // appid
	WechatMchId = `1281006301`         // 商户号
	NotifyUrl   = define.NgrokDomain + "/v1/wechatpay/notify"

	WechatpayKey = `4m5GYJrXeC3XbQd9WdqsfHy7o0BuisR5` // 支付key

	wechatCertPath = "res/wechatpay/apiclient_cert.pem" // 证书
	wechatKeyPath  = "res/wechatpay/apiclient_key.pem"  // 证书密钥
	wechatCAPath   = "res/wechatpay/rootca.pem"         // CA证书

	//wechatRefundURL = "https://api.mch.weixin.qq.com/secapi/pay/refund" // 退款url
	wechatBackPayURL = "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers" // 提现url
)

var (
	_tlsConfig *tls.Config // 证书tls.Config
)

// 签名
func doSign(vals map[string]string) string {
	if vals == nil {
		return ""
	}

	keys := make([]string, 0, len(vals))
	for k, v := range vals {
		if k != "sign" && len(v) > 0 {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys) // 排序

	var buf bytes.Buffer
	for _, k := range keys {
		vs := vals[k]
		buf.WriteString(k + "=" + vs + "&")
	}
	buf.WriteString("key=" + WechatpayKey)
	return strings.ToUpper(utils.Md5(string(buf.Bytes())))
}

// 排序并拼接参数
// 参照 url包的Encode() 方法
func parseSignParams(val url.Values) map[string]string {
	if val == nil {
		return nil
	}

	ret := make(map[string]string)
	for k, v := range val {
		if k != "sign" && len(v) > 0 && len(v[0]) > 0 {
			ret[k] = v[0]
		}
	}
	return ret
}

// 验证签名
func checkVerify(origData []byte, sign string) bool {
	if utils.Md5(string(origData)) != sign {
		return false
	}
	return true
}

// 验证notify
func VerifyNotify(vals map[string]string) error {
	// 验证签名
	sign, ok := vals["sign"]
	if !ok {
		return fmt.Errorf("没有sign参数")
	}
	verify := doSign(vals)
	if verify != sign {
		fmt.Println("签名失败")
		return fmt.Errorf("签名失败")
	}

	return nil
}

func getTLSConfig() (*tls.Config, error) {
	if _tlsConfig != nil {
		return _tlsConfig, nil
	}

	// load cert
	cert, err := tls.LoadX509KeyPair(wechatCertPath, wechatKeyPath)
	if err != nil {
		fmt.Println("load wechat keys fail", err)
		return nil, err
	}

	// load root ca
	caData, err := ioutil.ReadFile(wechatCAPath)
	if err != nil {
		fmt.Println("read wechat ca fail", err)
		return nil, err
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caData)

	_tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      pool,
	}
	return _tlsConfig, nil
}

// https post 请求
func securePost(url string, xmlContent []byte) (*http.Response, error) {
	tlsConfig, err := getTLSConfig()
	if err != nil {
		return nil, err
	}

	tr := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: tr}
	return client.Post(url, "text/xml", bytes.NewBuffer(xmlContent))
}

// 发送提现请求
func ReqBackPay(openId, tradeNo, clientIp, realName string, amount int64) (*models.WechatBackPayResult, error) {
	// https://pay.weixin.qq.com/wiki/doc/api/mch_pay.php?chapter=14_2
	type ParamsXml struct {
		XMLName        xml.Name `xml:"xml"`
		MchAppid       string   `xml:"mch_appid"`
		Mchid          string   `xml:"mchid"`
		NonceStr       string   `xml:"nonce_str"`
		PartnerTradeNo string   `xml:"partner_trade_no"`
		Openid         string   `xml:"openid"`
		CheckName      string   `xml:"check_name"`
		ReUserName     string   `xml:"re_user_name"`
		Amount         string   `xml:"amount"`
		Desc           string   `xml:"desc"`
		SpbillCreateIp string   `xml:"spbill_create_ip"`
		Sign           string   `xml:"sign"`
	}
	var paramsXml = ParamsXml{
		MchAppid:       WechatAppId,
		Mchid:          WechatMchId,
		NonceStr:       utils.RandString(32),
		PartnerTradeNo: tradeNo,
		Openid:         openId,
		CheckName:      "FORCE_CHECK", // FORCE_CHECK/NO_CHECK/OPTION_CHECK
		ReUserName:     realName,
		Amount:         strconv.FormatInt(amount, 10),
		Desc:           "提现",
		SpbillCreateIp: clientIp,
	}

	err := func(p *ParamsXml) error {
		var buf bytes.Buffer
		buf.WriteString("amount=" + p.Amount)
		buf.WriteString("&check_name=" + p.CheckName)
		buf.WriteString("&desc=" + p.Desc)
		buf.WriteString("&mch_appid=" + p.MchAppid)
		buf.WriteString("&mchid=" + p.Mchid)
		buf.WriteString("&nonce_str=" + p.NonceStr)
		buf.WriteString("&openid=" + p.Openid)
		buf.WriteString("&partner_trade_no=" + p.PartnerTradeNo)
		buf.WriteString("&re_user_name=" + p.ReUserName)
		buf.WriteString("&spbill_create_ip=" + p.SpbillCreateIp)
		buf.WriteString("&key=" + WechatpayKey)

		p.Sign = strings.ToUpper(utils.Md5(string(buf.Bytes())))
		if p.Sign == "" {
			return fmt.Errorf("签名错误")
		}
		return nil
	}(&paramsXml)
	if err != nil {
		return nil, err
	}

	xmlContent, err := xml.Marshal(&paramsXml)
	if err != nil {
		return nil, err
	}

	resp, err := securePost(wechatBackPayURL, xmlContent)
	if err != nil {
		return nil, fmt.Errorf("安全请求失败 %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("微信提现：安全请求失败", err.Error())
		return nil, fmt.Errorf("安全请求失败2 %s", err.Error())
	}
	fmt.Println("微信提现：返回xml", string(b))

	var result models.WechatBackPayResult
	err = xml.Unmarshal(b, &result)
	if err != nil {
		fmt.Println("微信提现：xml错误", err.Error())
		return nil, err
	}

	return &result, nil
}
