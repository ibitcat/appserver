// 第三方登陆相关的model

package models

type WeixinUserInfo struct {
	OpenId   string `json:"openid"`
	NickName string `json:"nickname"`
	Sex      int    `json:"sex"`
	Portrait string `json:"headimgurl"`
}

type QQUserInfo struct {
	RetCode  int    `json:"ret"`
	NickName string `json:"nickname"`
	Sex      string `json:"gender"`
	Portrait string `json:"figureurl"`
}

type WeiboUserInfo struct {
	Uid      int64  `json:"id"`
	NickName string `json:"screen_name"`
	Sex      string `json:"gender"`
	Portrait string `json:"profile_image_url"`
}

// 支付宝异步通知
// type AlipayNotify struct {
// 	NotifyTime       string `json:"notify_time"`                   // 通知时间
// 	NotifyType       string `json:"notify_type"`                   // 通知类型
// 	NotifyId         string `json:"notify_id"`                     // 通知校验ID（用来验证通知是否由支付宝发来的）
// 	SignType         string `json:"sign_type"`                     // 签名方式（固定用rsa签名）
// 	Sign             string `json:"sign"`                          // 签名
// 	OutTradeNo       string `json:"out_trade_no,omitempty"`        // app服务器生成的唯一订单号，可以为空
// 	Subject          string `json:"subject,omitempty"`             // 商品名称，可以为空
// 	PaymentType      string `json:"payment_type,omitempty"`        // 支付类型，可以为空
// 	TradeNo          string `json:"trade_no"`                      // 支付宝交易号
// 	TradeStatus      string `json:"trade_status"`                  // 交易状态
// 	SellerId         string `json:"seller_id"`                     // 卖家支付宝用户号
// 	SellerEmail      string `json:"seller_email"`                  // 卖家支付宝账号
// 	BuyerId          string `json:"buyer_id"`                      // 买家支付宝用户号
// 	BuyerEmail       string `json:"buyer_email"`                   // 买家支付宝账号，即app用户的支付宝账号
// 	TotalFee         int    `json:"total_fee"`                     // 交易金额
// 	Quantity         int    `json:"quantity,omitempty"`            // 购买数量，可以为空
// 	Price            int    `json:"price,omitempty"`               // 商品单价，可以为空
// 	Body             string `json:"body,omitempty"`                // 商品描述，可以为空
// 	GmtCreate        string `json:"gmt_create,omitempty"`          // 交易创建时间，可以为空
// 	GmtPayment       string `json:"gmt_payment,omitempty"`         // 交易付款时间，可以为空
// 	IsTotalFeeAdjust string `json:"is_total_fee_adjust,omitempty"` // 是否调整总价，可以为空
// 	UseCoupon        string `json:"use_coupon,omitempty"`          // 是否使用了红包，可以为空
// 	Discount         string `json:"discount,omitempty"`            // 折扣，可以为空
// 	RefundStatus     string `json:"refund_status,omitempty"`       // 退款状态，可以为空
// 	GmtRefund        string `json:"gmt_refund,omitempty"`          // 退款时间，可以为空
// }
