package unionpay

import (
	"bufio"
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"app-server/pkg/httplib"
	"golang.org/x/crypto/pkcs12"
)

var (
	PrivateCertId *big.Int        // 商户私钥证书序列号
	PublicCertId  *big.Int        // 银联公钥证书序列号
	EncryptCertId *big.Int        // 加密证书序列号
	RsaPrivateKey *rsa.PrivateKey // 私钥pem
	RsaPublicKey  *rsa.PublicKey  // 公钥
	RsaEncryptKey *rsa.PublicKey  // 加密证书内的公钥
	TestKey       *rsa.PublicKey  // 加密证书内的公钥
)

// 初始化银联sdk
func InitUnionPaySdk() {
	getCertIdAndPrivateKey()
	getCertIdAndPublicKey()
	getEncryptCert()

	fmt.Println("商户私钥cerid = ", PrivateCertId)
	fmt.Println("银联公钥cerid = ", PublicCertId)
	fmt.Println("加密证书cerid = ", EncryptCertId)
}

// 获取证书序列号serialNumber和私钥 (.pfx文件中获取)
func getCertIdAndPrivateKey() {
	pfxFile, err := os.Open(sdkSignCertPath)
	if err != nil {
		panic(err)
	}
	defer pfxFile.Close()

	fileInfo, statErr := pfxFile.Stat()
	if statErr != nil {
		panic(statErr)
	}

	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	buffer := bufio.NewReader(pfxFile)
	_, readErr := buffer.Read(bytes)
	if readErr != nil {
		panic(readErr)
	}

	prvKey, cert, decodeErr := pkcs12.Decode(bytes, sdkSignCertPwd)
	if decodeErr != nil {
		panic(decodeErr)
	}

	priv := prvKey.(*rsa.PrivateKey)
	if e := priv.Validate(); e != nil {
		panic(e)
	}
	RsaPrivateKey = priv
	PrivateCertId = cert.SerialNumber
	TestKey = cert.PublicKey.(*rsa.PublicKey)
}

// 从银联公钥证书中获取serialNumber (.cer文件中获取)
func getCertIdAndPublicKey() {
	certFile, err := os.Open(sdkVerifyCertPath)
	if err != nil {
		panic(err)
	}
	defer certFile.Close()

	info, statErr := certFile.Stat()
	if statErr != nil {
		panic(statErr)
	}

	var size int64 = info.Size()
	bytes := make([]byte, size)
	buffer := bufio.NewReader(certFile)
	_, readErr := buffer.Read(bytes)
	if readErr != nil {
		panic(readErr)
	}

	block, _ := pem.Decode(bytes)
	if block == nil {
		panic("failed to parse certificate PEM")
	}

	certificate, parseErr := x509.ParseCertificate(block.Bytes)
	if parseErr != nil {
		panic(parseErr)
	}

	RsaPublicKey = certificate.PublicKey.(*rsa.PublicKey)
	PublicCertId = certificate.SerialNumber
}

// 读取加密证书
func getEncryptCert() {
	certFile, err := os.Open(sdkEncryptCertPath)
	if err != nil {
		panic(err)
	}
	defer certFile.Close()

	info, statErr := certFile.Stat()
	if statErr != nil {
		panic(statErr)
	}

	var size int64 = info.Size()
	bytes := make([]byte, size)
	buffer := bufio.NewReader(certFile)
	_, readErr := buffer.Read(bytes)
	if readErr != nil {
		panic(readErr)
	}

	block, _ := pem.Decode(bytes)
	if block == nil {
		panic("failed to parse certificate PEM")
	}

	certificate, parseErr := x509.ParseCertificate(block.Bytes)
	if parseErr != nil {
		fmt.Println("========>", parseErr)
		//panic(parseErr)
	}

	EncryptCertId = certificate.SerialNumber
	RsaEncryptKey = certificate.PublicKey.(*rsa.PublicKey)

	// RsaPublicKey = certificate.PublicKey.(*rsa.PublicKey)

	// if RsaPublicKey.N.Cmp(RsaPrivateKey.N) != 0 {
	// 	fmt.Println("私钥和公钥不匹配")
	// }

	// PublicCertId = certificate.SerialNumber
}

// 解析字符串 （"key=vale&key1=value2"）
func coverStringToMap(str string) map[string]string {
	params := make(map[string]string)
	for _, kv := range strings.Split(str, "&") {
		s := strings.SplitN(kv, "=", 2)
		if len(s) == 2 {
			params[s[0]] = s[1]
		}
	}

	return params
}

// 排序参数并接连起来
func coverParamsToBytes(params map[string]string) []byte {
	var buf bytes.Buffer

	keys := make([]string, 0, len(params))
	for k, v := range params {
		if len(v) > 0 && k != "signature" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys) // 排序

	for i, k := range keys {
		if i > 0 {
			buf.WriteString("&")
		}
		buf.WriteString(k + "=" + params[k])
	}

	return buf.Bytes()
}

// 获取参数的签名
func signRequestParams(params map[string]string) string {
	paramsBytes := coverParamsToBytes(params)

	// sha1
	h := sha1.New()
	h.Write(paramsBytes)
	src := h.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src) //哈希后的摘要转为16进制

	// 使用sh1+ras私钥对签名加证书
	hashed := sha1.Sum(dst)
	sig, err := rsa.SignPKCS1v15(rand.Reader, RsaPrivateKey, crypto.SHA1, hashed[:])
	if err != nil {
		fmt.Println("签名失败 = ", err)
		return ""
	}

	return base64.StdEncoding.EncodeToString(sig)
}

// 发送请求，获取交易流水号
func sendHttpRequest(url string, params map[string]string) (map[string]string, error) {
	params["signature"] = signRequestParams(params) //签名

	req := httplib.Post(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}) //跳过证书验证
	for k, v := range params {
		if len(v) > 0 {
			req.Param(k, v)
		}
	}

	resp, err := req.String()
	if err != nil {
		return nil, err
	}

	fmt.Println("resp = ", resp)
	respParams := coverStringToMap(resp)
	fmt.Println("response certid = ", respParams["certId"])
	return respParams, nil
}

// 获取银联订单号
func GetTradeNo(orderId string, money int) (string, error) {
	params := make(map[string]string)
	params["version"] = "5.0.0"                             //版本号
	params["encoding"] = "utf-8"                            //编码方式
	params["certId"] = PrivateCertId.String()               //私钥证书ID
	params["txnType"] = "01"                                //交易类型
	params["txnSubType"] = "01"                             //交易子类
	params["bizType"] = "000201"                            //业务类型
	params["backUrl"] = sdkBackNotifyUrl                    //后台通知地址
	params["signMethod"] = "01"                             //签名方法
	params["channelType"] = "08"                            //渠道类型，07-PC，08-手机
	params["accessType"] = "0"                              //接入类型
	params["merId"] = sdkMerId                              //商户代码
	params["orderId"] = orderId                             //商户订单号，8-40位数字字母
	params["txnTime"] = time.Now().Format("20060102150405") //订单发送时间
	params["txnAmt"] = strconv.Itoa(money)                  //交易金额，单位分
	params["currencyCode"] = "156"                          //交易币种,默认人民币
	params["orderDesc"] = "desc"                            //订单描述，可不上送，上送时控件中会显示该信息
	params["reqReserved"] = "中文"                            //请求方保留域，透传字段，查询、通知、对账文件中均会原样出现

	respParams, err := sendHttpRequest(sdkAppRequstUrl, params)
	if err != nil {
		return "", err
	}

	verifyErr := Verify(respParams)
	if verifyErr != nil {
		return "", verifyErr
	}
	return respParams["tn"], nil
}

// 验签
func Verify(params map[string]string) error {
	// 验证结果
	if params["respCode"] != "00" {
		fmt.Println("银联处理错误，respmsg = ", params["respMsg"])
		return errors.New(params["respMsg"])
	}

	if params["certId"] != PublicCertId.String() { //公钥证书id不一样
		return errors.New("certid error")
	}

	// 排序重组报文域
	paramsBytes := coverParamsToBytes(params)

	// sha1
	h := sha1.New()
	h.Write(paramsBytes)
	src := h.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src) //哈希后的摘要转为16进制

	// 使用sh1+ras私钥对签名加证书
	hashed := sha1.Sum(dst)

	sig, _ := base64.StdEncoding.DecodeString(params["signature"]) //decode签名
	verifyErr := rsa.VerifyPKCS1v15(RsaPublicKey, crypto.SHA1, hashed[:], sig)
	fmt.Println("验签结果 err = ", verifyErr)
	if verifyErr != nil {
		return verifyErr
	}

	return nil
}

// 查询订单状态
func QueryTradeStatus(orderId string) error {
	params := make(map[string]string)
	params["version"] = "5.0.0"                             //版本号
	params["encoding"] = "utf-8"                            //编码方式
	params["certId"] = PrivateCertId.String()               //私钥证书ID
	params["signMethod"] = "01"                             //签名方法
	params["txnType"] = "00"                                //交易类型
	params["txnSubType"] = "00"                             //交易子类
	params["bizType"] = "000000"                            //业务类型
	params["accessType"] = "0"                              //接入类型
	params["channelType"] = "07"                            //渠道类型，07-PC，08-手机
	params["orderId"] = orderId                             //商户订单号，8-40位数字字母
	params["merId"] = sdkMerId                              //商户代码
	params["txnTime"] = time.Now().Format("20060102150405") //订单发送时间

	respParams, err := sendHttpRequest(sdkSingleQueryUrl, params)
	fmt.Printf("QueryTradeStatus response = %+v \n", respParams)
	if err != nil {
		return err
	}

	verifyErr := Verify(respParams)
	if verifyErr != nil {
		return verifyErr
	}
	return nil
}

// 获取持卡人身份信息
func getCustomerInfoStrNew() string {
	customerInfo := make(map[string]string)
	customerInfo["certifTp"] = "01"
	customerInfo["certifId"] = "341126197709218366"
	customerInfo["customerNm"] = "全渠道"

	var str string
	var count int = 0
	for k, v := range customerInfo {
		if count > 0 {
			str += "&"
		}
		str += (k + "=" + v)
		count++
	}

	outStr := "{" + "certifId=341126197709218366&certifTp=01&customerNm=全渠道" + "}"
	return base64.StdEncoding.EncodeToString([]byte(outStr))
}

// 请求提现
func WithdrawCash(orderId string, money int) error {
	fmt.Println("orderid = ", orderId)
	out, rsaErr := rsa.EncryptPKCS1v15(rand.Reader, TestKey, []byte("6216261000000000018"))
	if rsaErr != nil {
		return rsaErr
	}
	accNo := base64.StdEncoding.EncodeToString(out)
	fmt.Println("加密后的accNo=", accNo, "\n")

	params := make(map[string]string)
	params["version"] = "5.0.0"                             //版本号
	params["encoding"] = "utf-8"                            //编码方式
	params["certId"] = PrivateCertId.String()               //私钥证书ID
	params["encryptCertId"] = "123131"                      //EncryptCertId.String()        //公钥证书ID
	params["signMethod"] = "01"                             //签名方法
	params["txnType"] = "12"                                //交易类型
	params["txnSubType"] = "00"                             //交易子类
	params["bizType"] = "000401"                            //业务类型
	params["accessType"] = "0"                              //接入类型
	params["channelType"] = "08"                            //渠道类型，07-PC，08-手机
	params["orderId"] = orderId                             //商户订单号，8-40位数字字母
	params["merId"] = sdkMerId                              //商户代码
	params["txnTime"] = time.Now().Format("20060102150405") //订单发送时间
	params["txnAmt"] = strconv.Itoa(money)                  //交易金额，单位分
	params["currencyCode"] = "156"                          //交易币种,默认人民币
	params["reqReserved"] = "reqReserved"                   //请求方保留域，透传字段，查询、通知、对账文件中均会原样出现
	params["backUrl"] = sdkBackPayNotifyUrl                 //后台通知地址
	params["accNo"] = accNo                                 //"6216261000000000018"                 //收款卡号
	params["customerInfo"] = getCustomerInfoStrNew()        //银行卡验证信息及身份信息，证件信息与姓名至少出现一个

	respParams, err := sendHttpRequest(sdkBackTransUrl, params)
	//fmt.Printf("%+v\n", respParams)
	sig, _ := base64.StdEncoding.DecodeString(respParams["accNo"]) //decode签名
	acc, _ := rsa.DecryptPKCS1v15(rand.Reader, RsaPrivateKey, sig)

	fmt.Println("返回的accNo=", respParams["accNo"], "\n")
	fmt.Println("解密后的accNo = ", string(acc))

	if err != nil {
		return err
	}

	verifyErr := Verify(respParams)
	if verifyErr != nil {
		return verifyErr
	}

	return nil
}
