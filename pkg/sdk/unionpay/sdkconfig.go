package unionpay

import (
	"app-server/define"
)

const (
	sdkSignCertPath     = "res/unionpay/700000000000001_acp.pfx"                    // 商户私钥证书(签名)
	sdkVerifyCertPath   = "res/unionpay/verify_sign_acp.cer"                        // 银联公钥证书(验签)
	sdkEncryptCertPath  = "res/unionpay/acp_test_enc.cer"                           // 加密证书
	sdkSignCertPwd      = "000000"                                                  // 签名证书密码
	sdkAppRequstUrl     = "https://101.231.204.80:5000/gateway/api/appTransReq.do"  // App交易地址
	sdkSingleQueryUrl   = "https://101.231.204.80:5000/gateway/api/queryTrans.do"   // 交易订单状态查询地址
	sdkBackTransUrl     = "https://101.231.204.80:5000/gateway/api/backTransReq.do" // 银联提现api地址
	sdkBackPayNotifyUrl = define.NgrokDomain + "/v1/unionpay/backpaynotify"         // 后台取现回调地址，由银联服务器回调
	sdkBackNotifyUrl    = define.NgrokDomain + "/v1/unionpay/notify"                // 后台通知地址，由银联服务器回调
	sdkMerId            = "802440048160542"                                         // 商户号
)
