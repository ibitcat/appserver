package alipay

import (
	"bufio"
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sort"
	"time"

	"app-server/define"
	"app-server/pkg/httplib"
)

const (
	AlipayPid         = "2088021966427704"                       // 支付宝pid
	AlipayEmail       = "du@shuteng8.com"                        // 商户的支付宝账邮箱
	AlipayAccName     = "广州xx信息技术有限公司"                           // 商户支付宝账号名
	NotifyUrl         = define.NgrokDomain + "/v1/alipay/notify" // 服务器回调api
	httpsVerifyUrl    = "https://mapi.alipay.com/gateway.do"     // 支付宝通知验证url
	aliPublicKeyPath  = "res/alipay/rsa_public_key.pem"          // 支付宝公钥
	aliPrivateKeyPath = "res/alipay/rsa_private_key.pem"         // 支付宝私钥
	AliMd5Key         = "3r37kqytv3spkmngqart511daog6b8ry"       // 支付宝MD5
)

var (
	aliPublicKey  *rsa.PublicKey
	aliPrivateKey *rsa.PrivateKey
)

// 初始化支付宝sdk
func InitAlipaySDK() {
	aliPublicKey = getAliRsaPublicKey()
}

// 读取支付宝rsa公钥
func getAliRsaPublicKey() *rsa.PublicKey {
	publicKeyFile, err := os.Open(aliPublicKeyPath)
	if err != nil {
		panic(err)
	}

	defer publicKeyFile.Close()

	pemfileinfo, _ := publicKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)
	if err != nil {
		panic(err)
	}

	data, _ := pem.Decode([]byte(pembytes))
	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)
	if err != nil {
		panic(err)
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)
	if !ok {
		panic(err)
	}

	return rsaPub
}

// 验证是否由支付宝发过来的通知
func getResponseVerify(notifyId string) bool {
	req := httplib.Get(httpsVerifyUrl)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Param("service", "notify_verify")
	req.Param("partner", AlipayPid)
	req.Param("notify_id", notifyId)

	result, err := req.String()
	if err != nil {
		return false
	}

	return result == "true"
}

// 排序并拼接参数
// 参照 url包的Encode() 方法
func createLinkstring(v url.Values) []byte {
	if v == nil {
		return nil
	}

	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		if k != "sign" && k != "sign_type" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys) // 排序

	for _, k := range keys {
		vs := v[k]
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(k + "=" + v)
		}
	}
	return buf.Bytes()
}

// 验证签名
func checkVerify(origData []byte, sign string) bool {
	h := sha1.New()
	h.Write(origData)
	src := h.Sum(nil)

	sig, _ := base64.StdEncoding.DecodeString(sign) //decode签名
	err := rsa.VerifyPKCS1v15(aliPublicKey, crypto.SHA1, src, sig)
	if err != nil {
		fmt.Println("getSignVerify err ", err)
		return false
	}

	return true
}

// 验证notify
// 为什么不用 gin的binding，是因为需要变量params验签
func VerifyNotify(request *http.Request) (ok bool, tradeNo, totalFee string) {
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded") //这里处理golang的bug
	request.ParseForm()

	// 坑爹 golang 的bug
	//b, _ := ioutil.ReadAll(request.Body)
	//postForm, _ := url.ParseQuery(string(b))

	ok = false
	if request.FormValue("sign_type") != "RSA" { // 固定为rsa算法
		fmt.Println("固定为rsa算法")
		return
	}

	// 验证签名
	sign := request.FormValue("sign")
	linkStr := createLinkstring(request.Form)
	if linkStr == nil || !checkVerify(linkStr, sign) {
		fmt.Println("验签失败")
		return
	}

	// 验证是否由支付宝发过来的通知
	notifyId := request.FormValue("notify_id")
	if !getResponseVerify(notifyId) {
		fmt.Println("支付宝发过来的通知")
		return
	}

	if request.FormValue("trade_status") != "TRADE_SUCCESS" {
		fmt.Println("TRADE_SUCCESS")
		return
	}

	ok = true
	tradeNo = request.FormValue("out_trade_no")
	totalFee = request.FormValue("total_fee")
	return
}

// 批量付款
func WithdrawCash() {
	params := make(map[string]string)
	params["service"] = "batch_trans_notify"
	params["partner"] = AlipayPid
	params["_input_charset"] = "utf-8"
	//params["notify_url"] = NotifyUrl
	params["sign_type"] = "MD5"
	params["email"] = AlipayEmail
	params["pay_date"] = time.Now().Format("20060102")
	params["batch_no"] = time.Now().Format("20060102150405")
	params["batch_num"] = "1"
	params["account_name"] = AlipayAccName
	params["batch_fee"] = "1.01"
	params["detail_data"] = "0315006^shui_mu98@163.com^鹏飞^1.01^hello"
	params["extend_param"] = "agent^123456"

	keys := make([]string, 0, len(params))
	for k, v := range params {
		if k == "sign" || k == "sign_type" || len(v) == 0 {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys) // 排序

	var paramStr string
	for i, k := range keys {
		if i > 0 {
			paramStr += "&"
		}
		paramStr += (k + "=" + params[k])
	}

	h := md5.New()
	h.Write([]byte(paramStr))
	sign := string(h.Sum(nil))
	params["sign"] = sign

	// http请求
	req := httplib.Get(httpsVerifyUrl)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	for k, v := range params {
		req.Param(k, v)
	}

	result, err := req.String()
	fmt.Println("---------->", result, err)
}
