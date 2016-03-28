package api

//This file is generated automatically. Do not try to edit it manually.

var resourceListingJson = `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://192.168.1.106:8080/docs",
    "apis": [
        {
            "path": "/qiniu",
            "description": "七牛云存储"
        },
        {
            "path": "/socially",
            "description": "社交"
        },
        {
            "path": "/wechatpay",
            "description": "微信支付相关"
        },
        {
            "path": "/alipay",
            "description": "支付宝相关"
        },
        {
            "path": "/enterprise_cert",
            "description": "企业认证相关"
        },
        {
            "path": "/pub",
            "description": "公共接口"
        },
        {
            "path": "/redpacket",
            "description": "抢/发红包"
        },
        {
            "path": "/scanning",
            "description": "扫红包"
        },
        {
            "path": "/unionpay",
            "description": "银联支付相关"
        },
        {
            "path": "/user",
            "description": "用户操作"
        },
        {
            "path": "/backpay",
            "description": "提现（汇款/银联/支付宝）"
        },
        {
            "path": "/oauth",
            "description": "第三方认证"
        },
        {
            "path": "/payphone",
            "description": "手机充话费/流量"
        }
    ],
    "info": {
        "title": "AppServer API",
        "description": "appserver swagger 文档测试"
    }
}`
var apiDescriptionsJson = map[string]string{"payphone": `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://192.168.1.106:8080/v1",
    "resourcePath": "/payphone",
    "apis": [
        {
            "path": "/payphone/phone_recharge_balance",
            "description": "查询余额",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "phoneRechargeBalance",
                    "type": "app-server.models.PhoneRechargeBalance",
                    "items": {},
                    "summary": "查询余额",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.PhoneRechargeBalance"
                        },
                        {
                            "code": 400,
                            "message": "错误",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/payphone/phone_recharge_query",
            "description": "根据手机号和充值额度查询商品信息",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "phoneRechargeQuery",
                    "type": "app-server.models.PhoneRechargeQueryResp",
                    "items": {},
                    "summary": "根据手机号和充值额度查询商品信息",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "phone",
                            "description": "手机号",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "price",
                            "description": "充值金额（有效值参考apix）",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "apix数据直接返回",
                            "responseType": "object",
                            "responseModel": "app-server.models.PhoneRechargeQueryResp"
                        },
                        {
                            "code": 400,
                            "message": "错误",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/payphone/phone_recharge",
            "description": "充值话费",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "phoneRecharge",
                    "type": "app-server.models.PhoneRechargeResp",
                    "items": {},
                    "summary": "充值话费",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "phone",
                            "description": "手机号",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "form",
                            "name": "price",
                            "description": "充值金额（有效值参考apix）",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "apix数据直接返回",
                            "responseType": "object",
                            "responseModel": "app-server.models.PhoneRechargeResp"
                        },
                        {
                            "code": 400,
                            "message": "错误",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/payphone/data_recharge_query",
            "description": "查询号码支持的流量套餐",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "dataRechargeQuery",
                    "type": "app-server.models.DataRechargeQueryResp",
                    "items": {},
                    "summary": "查询号码支持的流量套餐",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "phone",
                            "description": "手机号",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "apix数据直接返回",
                            "responseType": "object",
                            "responseModel": "app-server.models.DataRechargeQueryResp"
                        },
                        {
                            "code": 400,
                            "message": "错误",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/payphone/data_recharge",
            "description": "充值流量",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "dataRecharge",
                    "type": "app-server.models.DataRechargeResp",
                    "items": {},
                    "summary": "充值流量",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "phone",
                            "description": "手机号",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "form",
                            "name": "pkgid",
                            "description": "套餐ID（根据查询号码支持的流量套餐接口得到）",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "apix数据直接返回",
                            "responseType": "object",
                            "responseModel": "app-server.models.DataRechargeResp"
                        },
                        {
                            "code": 400,
                            "message": "错误",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/payphone/phone_recharge_notify",
            "description": "充值话费回调",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "phoneRechargeNotify",
                    "type": "string",
                    "items": {},
                    "summary": "充值话费回调",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "state",
                            "description": "充值状态",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "orderid",
                            "description": "商家订单号",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "ordertime",
                            "description": "订单处理时间",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "sign",
                            "description": "签名",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "err_msg",
                            "description": "失败信息",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "SUCCESS",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "错误信息",
                            "responseType": "string",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/payphone/data_recharge_notify",
            "description": "充值流量回调",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "dataRechargeNotify",
                    "type": "string",
                    "items": {},
                    "summary": "充值流量回调",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "state",
                            "description": "充值状态",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "orderid",
                            "description": "商家订单号",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "ordertime",
                            "description": "订单处理时间",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "sign",
                            "description": "签名",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "err_msg",
                            "description": "失败信息",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "SUCCESS",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "错误信息",
                            "responseType": "string",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "app-server.models.APIError": {
            "id": "app-server.models.APIError",
            "properties": {
                "code": {
                    "type": "uint32",
                    "description": "错误码",
                    "items": {},
                    "format": ""
                },
                "msg": {
                    "type": "string",
                    "description": "错误描述",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.DataRechargeData": {
            "id": "app-server.models.DataRechargeData",
            "properties": {
                "Cardname": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "Ordercash": {
                    "type": "float64",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "Phone": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "SporderId": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "State": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "UserOrderId": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.DataRechargeQueryData": {
            "id": "app-server.models.DataRechargeQueryData",
            "properties": {
                "ProviderId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "ProviderName": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "UserDataPackages": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "app-server.models.DataRechargeUserDataPackages"
                    },
                    "format": ""
                }
            }
        },
        "app-server.models.DataRechargeQueryResp": {
            "id": "app-server.models.DataRechargeQueryResp",
            "properties": {
                "Code": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "Data": {
                    "type": "app-server.models.DataRechargeQueryData",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "Msg": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.DataRechargeResp": {
            "id": "app-server.models.DataRechargeResp",
            "properties": {
                "Code": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "Data": {
                    "type": "app-server.models.DataRechargeData",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "Msg": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.DataRechargeUserDataPackages": {
            "id": "app-server.models.DataRechargeUserDataPackages",
            "properties": {
                "Cost": {
                    "type": "float64",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "DataValue": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "EffectStartTime": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "EffectTime": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "LimitTimes": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "PkgId": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "Price": {
                    "type": "float64",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "Scope": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "Support4G": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.PhoneRechargeBalance": {
            "id": "app-server.models.PhoneRechargeBalance",
            "properties": {
                "Balance": {
                    "type": "float64",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.PhoneRechargeData": {
            "id": "app-server.models.PhoneRechargeData",
            "properties": {
                "Cardid": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "Cardname": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "Cardnum": {
                    "type": "float64",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "Ordercash": {
                    "type": "float64",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "Phone": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "SporderId": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "State": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "UserOrderId": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.PhoneRechargeQueryData": {
            "id": "app-server.models.PhoneRechargeQueryData",
            "properties": {
                "Cardid": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "Cardname": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "GameArea": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "Inprice": {
                    "type": "float64",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.PhoneRechargeQueryResp": {
            "id": "app-server.models.PhoneRechargeQueryResp",
            "properties": {
                "Code": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "Data": {
                    "type": "app-server.models.PhoneRechargeQueryData",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "Msg": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.PhoneRechargeResp": {
            "id": "app-server.models.PhoneRechargeResp",
            "properties": {
                "Code": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "Data": {
                    "type": "app-server.models.PhoneRechargeData",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "Msg": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`, "wechatpay": `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://192.168.1.106:8080/v1",
    "resourcePath": "/wechatpay",
    "apis": [
        {
            "path": "/wechatpay/notify",
            "description": "微信付款异步回调api",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "wechatpayNotify",
                    "type": "app-server.models.WechatPayResult",
                    "items": {},
                    "summary": "微信付款异步回调api",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "微信支付回调成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.WechatPayResult"
                        },
                        {
                            "code": 400,
                            "message": "微信支付回调失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.WechatPayResult"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/wechatpay/params",
            "description": "获取微信支付参数",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "getWechatpayParams",
                    "type": "app-server.models.WechatPayParams",
                    "items": {},
                    "summary": "获取微信支付参数",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.WechatPayParams"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/wechatpay/backpay",
            "description": "微信提现",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "wechatBackPay",
                    "type": "string",
                    "items": {},
                    "summary": "微信提现",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "money",
                            "description": "要提现的金额",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "realname",
                            "description": "收款用户真实姓名",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "passwd",
                            "description": "密码",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "app-server.models.APIError": {
            "id": "app-server.models.APIError",
            "properties": {
                "code": {
                    "type": "uint32",
                    "description": "错误码",
                    "items": {},
                    "format": ""
                },
                "msg": {
                    "type": "string",
                    "description": "错误描述",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.WechatPayParams": {
            "id": "app-server.models.WechatPayParams",
            "properties": {
                "appid": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "mchid": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "notifyurl": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "paykey": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.WechatPayResult": {
            "id": "app-server.models.WechatPayResult",
            "properties": {
                "ReturnCode": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "ReturnMsg": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`, "alipay": `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://192.168.1.106:8080/v1",
    "resourcePath": "/alipay",
    "apis": [
        {
            "path": "/alipay/notify",
            "description": "支付宝异步回调api",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "alipay_notify",
                    "type": "string",
                    "items": {},
                    "summary": "支付宝异步回调api",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "回调成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "回调失败",
                            "responseType": "string",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/alipay/tradeno",
            "description": "生成订单号",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "outTradeNo",
                    "type": "string",
                    "items": {},
                    "summary": "生成订单号",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "生成订单成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "生成订单失败",
                            "responseType": "string",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/alipay/params",
            "description": "获取支付宝支付参数",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "alipayParams",
                    "type": "app-server.models.AlipayParams",
                    "items": {},
                    "summary": "获取支付宝支付参数",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.AlipayParams"
                        },
                        {
                            "code": 400,
                            "message": "失败",
                            "responseType": "string",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "app-server.models.AlipayParams": {
            "id": "app-server.models.AlipayParams",
            "properties": {
                "alipayacc": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "alipaypid": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "notifyurl": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`, "redpacket": `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://192.168.1.106:8080/v1",
    "resourcePath": "/redpacket",
    "apis": [
        {
            "path": "/redpacket/paytest",
            "description": "红包付款测试",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "paytest",
                    "type": "string",
                    "items": {},
                    "summary": "红包付款测试",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "id",
                            "description": "红包id",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "money",
                            "description": "支付金额",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "付款成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "付款失败",
                            "responseType": "string",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/redpacket/list",
            "description": "获取红包列表",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "redpacket",
                    "type": "app-server.models.S2C_RedpacketList",
                    "items": {},
                    "summary": "获取红包列表",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "startidx",
                            "description": "查询开始索引",
                            "dataType": "uint32",
                            "type": "uint32",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "获取红包列表成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_RedpacketList"
                        },
                        {
                            "code": 400,
                            "message": "获取红包列表失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/redpacket/record",
            "description": "抢红包的记录",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "record",
                    "type": "app-server.models.S2C_RedpktRecord",
                    "items": {},
                    "summary": "抢红包的记录",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "redpacketid",
                            "description": "红包id",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "cursor",
                            "description": "游标",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "领取红包记录",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_RedpktRecord"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/redpacket/verify",
            "description": "审核红包[app和游戏红包]",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "verify",
                    "type": "string",
                    "items": {},
                    "summary": "审核红包[app和游戏红包]",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "redpacketid",
                            "description": "红包id",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "status",
                            "description": "审核结果",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "操作成功",
                            "responseType": "string",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/redpacket/send",
            "description": "发红包",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "sendredpacket",
                    "type": "",
                    "items": {},
                    "summary": "发红包",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "data",
                            "description": "红包信息",
                            "dataType": "app-server.models.SendRedpacketBinding",
                            "type": "app-server.models.SendRedpacketBinding",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "发红包成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "发红包失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/redpacket/grab",
            "description": "抢红包",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "grabredpacket",
                    "type": "string",
                    "items": {},
                    "summary": "抢红包",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "redpacketid",
                            "description": "红包id",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "deviceid",
                            "description": "设备id",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "抢红包成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "抢红包失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/redpacket/share",
            "description": "完成红包分享任务，等待6小时后截图",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "doingtask",
                    "type": "string",
                    "items": {},
                    "summary": "完成红包分享任务，等待6小时后截图",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "redpacketid",
                            "description": "红包id",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "deviceid",
                            "description": "设备id",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "完成分享",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/redpacket/finish",
            "description": "完成红包任务",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "finishredpacket",
                    "type": "string",
                    "items": {},
                    "summary": "完成红包任务",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "redpacketid",
                            "description": "红包id",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "deviceid",
                            "description": "设备id",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "完成红包任务，并获得红包",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "完成失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/redpacket/giveup",
            "description": "放弃红包",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "giveupredpacket",
                    "type": "string",
                    "items": {},
                    "summary": "放弃红包",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "redpacketid",
                            "description": "红包id",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "放弃红包",
                            "responseType": "string",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/redpacket/pay",
            "description": "通过余额支付红包",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "paybybalance",
                    "type": "string",
                    "items": {},
                    "summary": "通过余额支付红包",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "redpacketid",
                            "description": "红包id",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "form",
                            "name": "password",
                            "description": "支付密码",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "支付成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "支付失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/redpacket/filter",
            "description": "根据用户过滤红包",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "filter",
                    "type": "app-server.models.UserExpireList",
                    "items": {},
                    "summary": "根据用户过滤红包",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "根据玩家过滤红包池中的红包",
                            "responseType": "object",
                            "responseModel": "app-server.models.UserExpireList"
                        },
                        {
                            "code": 400,
                            "message": "过滤失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/redpacket/statistics",
            "description": "获取红包的统计数据",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "Statistics",
                    "type": "app-server.models.S2C_RedpktStatistics",
                    "items": {},
                    "summary": "获取红包的统计数据",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "红包统计数据",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_RedpktStatistics"
                        },
                        {
                            "code": 400,
                            "message": "获取失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/redpacket/redpkt_recieve",
            "description": "收到的红包",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "redpacketRecieve",
                    "type": "app-server.models.S2C_RedpketRecieveInfo",
                    "items": {},
                    "summary": "收到的红包",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_RedpketRecieveInfo"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/redpacket/recievelist",
            "description": "收到的红包记录",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "RecieveList",
                    "type": "app-server.models.S2C_ReceivedList",
                    "items": {},
                    "summary": "收到的红包记录",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "date",
                            "description": "日期,（格式 20060102）",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_ReceivedList"
                        },
                        {
                            "code": 400,
                            "message": "失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/redpacket/redpkt_send",
            "description": "发出的红包",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "redpacketSend",
                    "type": "app-server.models.S2C_RedpketSendInfo",
                    "items": {},
                    "summary": "发出的红包",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "year",
                            "description": "年份",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_RedpketSendInfo"
                        },
                        {
                            "code": 400,
                            "message": "失败",
                            "responseType": "string",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/redpacket/sendlist",
            "description": "发出的红包记录",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "redpacketSend",
                    "type": "app-server.models.S2C_RedpktSendList",
                    "items": {},
                    "summary": "发出的红包记录",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "year",
                            "description": "年份",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "cursor",
                            "description": "分页游标(0开始)",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_RedpktSendList"
                        },
                        {
                            "code": 400,
                            "message": "失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/redpacket/ready",
            "description": "待发出的红包",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "tobereleased",
                    "type": "app-server.models.S2C_ToBeReleasedList",
                    "items": {},
                    "summary": "待发出的红包",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_ToBeReleasedList"
                        },
                        {
                            "code": 400,
                            "message": "失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "app-server.models.APIError": {
            "id": "app-server.models.APIError",
            "properties": {
                "code": {
                    "type": "uint32",
                    "description": "错误码",
                    "items": {},
                    "format": ""
                },
                "msg": {
                    "type": "string",
                    "description": "错误描述",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.AppInfo": {
            "id": "app-server.models.AppInfo",
            "properties": {
                "bundle_id": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "icon_url": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "keyword": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "name": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "platform": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "price": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "size": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "url_scheme": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.AreaInfo": {
            "id": "app-server.models.AreaInfo",
            "properties": {
                "city": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "country": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "district": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "province": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.GrabRecord": {
            "id": "app-server.models.GrabRecord",
            "properties": {
                "grab_money": {
                    "type": "uint32",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "grab_time": {
                    "type": "int64",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "redpacket_id": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "redpacket_name": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "redpacket_type": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.InvoiceAddress": {
            "id": "app-server.models.InvoiceAddress",
            "properties": {
                "address": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "addressee": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "tel": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "title": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.OfficialAccInfo": {
            "id": "app-server.models.OfficialAccInfo",
            "properties": {
                "name": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "title": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "url": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.RedpacketExpire": {
            "id": "app-server.models.RedpacketExpire",
            "properties": {
                "download": {
                    "type": "int64",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "is_grab": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "screenshot": {
                    "type": "int64",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "share": {
                    "type": "int64",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "type": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.RedpacketInfo": {
            "id": "app-server.models.RedpacketInfo",
            "properties": {
                "app": {
                    "type": "app-server.models.AppInfo",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "area": {
                    "type": "app-server.models.AreaInfo",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "begin_time": {
                    "type": "int64",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "end_time": {
                    "type": "int64",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "is_auth": {
                    "type": "uint8",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "number": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "official_acc": {
                    "type": "app-server.models.OfficialAccInfo",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "per_money": {
                    "type": "uint32",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "portrait": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "share": {
                    "type": "app-server.models.ShareInfo",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "type": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "user_id": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "user_name": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.RedpktRecord": {
            "id": "app-server.models.RedpktRecord",
            "properties": {
                "nickname": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "time": {
                    "type": "int64",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "userid": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.S2C_ReceivedList": {
            "id": "app-server.models.S2C_ReceivedList",
            "properties": {
                "list": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "app-server.models.GrabRecord"
                    },
                    "format": ""
                }
            }
        },
        "app-server.models.S2C_RedpacketList": {
            "id": "app-server.models.S2C_RedpacketList",
            "properties": {
                "Count": {
                    "type": "uint32",
                    "description": "该时段红包个数",
                    "items": {},
                    "format": ""
                },
                "List": {
                    "type": "array",
                    "description": "红包列表",
                    "items": {
                        "$ref": "app-server.models.RedpacketInfo"
                    },
                    "format": ""
                }
            }
        },
        "app-server.models.S2C_RedpketRecieveInfo": {
            "id": "app-server.models.S2C_RedpketRecieveInfo",
            "properties": {
                "rank": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "total": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.S2C_RedpketSendInfo": {
            "id": "app-server.models.S2C_RedpketSendInfo",
            "properties": {
                "amount": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "total": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.S2C_RedpktRecord": {
            "id": "app-server.models.S2C_RedpktRecord",
            "properties": {
                "id": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "list": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "app-server.models.RedpktRecord"
                    },
                    "format": ""
                }
            }
        },
        "app-server.models.S2C_RedpktSendList": {
            "id": "app-server.models.S2C_RedpktSendList",
            "properties": {
                "list": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "app-server.models.SendRedpacket"
                    },
                    "format": ""
                }
            }
        },
        "app-server.models.S2C_RedpktStatistics": {
            "id": "app-server.models.S2C_RedpktStatistics",
            "properties": {
                "area": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "type": "int"
                    },
                    "format": ""
                },
                "count": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "sex": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "type": "int"
                    },
                    "format": ""
                }
            }
        },
        "app-server.models.S2C_ToBeReleasedList": {
            "id": "app-server.models.S2C_ToBeReleasedList",
            "properties": {
                "list": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "app-server.models.ToBeReleasedRedpkt"
                    },
                    "format": ""
                }
            }
        },
        "app-server.models.SendRedpacket": {
            "id": "app-server.models.SendRedpacket",
            "properties": {
                "begin_time": {
                    "type": "int64",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "end_time": {
                    "type": "int64",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "per_money": {
                    "type": "uint32",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "remainder": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "title": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "total": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.SendRedpacketBinding": {
            "id": "app-server.models.SendRedpacketBinding",
            "properties": {
                "address": {
                    "type": "app-server.models.InvoiceAddress",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "app": {
                    "type": "app-server.models.AppInfo",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "area": {
                    "type": "app-server.models.AreaInfo",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "begintime": {
                    "type": "int64",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "invoice": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "official_acc": {
                    "type": "app-server.models.OfficialAccInfo",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "permoney": {
                    "type": "uint32",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "share": {
                    "type": "app-server.models.ShareInfo",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "totalnum": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "type": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.ShareInfo": {
            "id": "app-server.models.ShareInfo",
            "properties": {
                "image_uri": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "title": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.ToBeReleasedRedpkt": {
            "id": "app-server.models.ToBeReleasedRedpkt",
            "properties": {
                "begin_time": {
                    "type": "int64",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "create_time": {
                    "type": "int64",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "title": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "trade_status": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "verify": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.UserExpireList": {
            "id": "app-server.models.UserExpireList",
            "properties": {
                "list": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "app-server.models.RedpacketExpire"
                    },
                    "format": ""
                }
            }
        }
    }
}`, "scanning": `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://192.168.1.106:8080/v1",
    "resourcePath": "/scanning",
    "apis": [
        {
            "path": "/scanning/list",
            "description": "获取扫红包列表",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "list",
                    "type": "app-server.models.ScanningList",
                    "items": {},
                    "summary": "获取扫红包列表",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "tag",
                            "description": "商品标签",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "startidx",
                            "description": "查询开始索引",
                            "dataType": "uint32",
                            "type": "uint32",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "获取红包列表成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.ScanningList"
                        },
                        {
                            "code": 400,
                            "message": "获取红包列表失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/scanning/send",
            "description": "商户发送扫红包",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "send",
                    "type": "",
                    "items": {},
                    "summary": "商户发送扫红包",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "data",
                            "description": "红包信息",
                            "dataType": "app-server.models.SendScannigBinding",
                            "type": "app-server.models.SendScannigBinding",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "获取红包列表成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "获取红包列表失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/scanning/scan",
            "description": "用户扫红包",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "scan",
                    "type": "string",
                    "items": {},
                    "summary": "用户扫红包",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "redpacketid",
                            "description": "红包id",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "扫红包成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "扫红包失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "app-server.models.APIError": {
            "id": "app-server.models.APIError",
            "properties": {
                "code": {
                    "type": "uint32",
                    "description": "错误码",
                    "items": {},
                    "format": ""
                },
                "msg": {
                    "type": "string",
                    "description": "错误描述",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.ScanningItem": {
            "id": "app-server.models.ScanningItem",
            "properties": {
                "balance": {
                    "type": "uint32",
                    "description": "红包余额",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "string",
                    "description": "扫红包商品id",
                    "items": {},
                    "format": ""
                },
                "name": {
                    "type": "string",
                    "description": "商品名称",
                    "items": {},
                    "format": ""
                },
                "pic": {
                    "type": "string",
                    "description": "商品图片",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.ScanningList": {
            "id": "app-server.models.ScanningList",
            "properties": {
                "count": {
                    "type": "uint32",
                    "description": "列表长度",
                    "items": {},
                    "format": ""
                },
                "list": {
                    "type": "array",
                    "description": "扫红包列表",
                    "items": {
                        "$ref": "app-server.models.ScanningItem"
                    },
                    "format": ""
                }
            }
        },
        "app-server.models.SendScannigBinding": {
            "id": "app-server.models.SendScannigBinding",
            "properties": {
                "barcode": {
                    "type": "string",
                    "description": "条形码",
                    "items": {},
                    "format": ""
                },
                "desc": {
                    "type": "string",
                    "description": "促销信息",
                    "items": {},
                    "format": ""
                },
                "itemname": {
                    "type": "string",
                    "description": "商品名称",
                    "items": {},
                    "format": ""
                },
                "start_date": {
                    "type": "uint32",
                    "description": "红包发送日期",
                    "items": {},
                    "format": ""
                },
                "stop_date": {
                    "type": "uint32",
                    "description": "红包结束日期",
                    "items": {},
                    "format": ""
                },
                "tag": {
                    "type": "string",
                    "description": "商品分类",
                    "items": {},
                    "format": ""
                },
                "total_money": {
                    "type": "uint32",
                    "description": "红包总金额",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`, "user": `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://192.168.1.106:8080/v1",
    "resourcePath": "/user",
    "apis": [
        {
            "path": "/user/login",
            "description": "用户登录",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "login",
                    "type": "",
                    "items": {},
                    "summary": "用户登录",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "username",
                            "description": "手机号码/账号",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "form",
                            "name": "password",
                            "description": "用户密码",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "用户登录成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_TokenArray"
                        },
                        {
                            "code": 400,
                            "message": "用户登录失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/user/logout",
            "description": "用户退出",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "logout",
                    "type": "",
                    "items": {},
                    "summary": "用户退出",
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "用户退出成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "退出错误",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/user/register",
            "description": "用户手机号码注册",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "register",
                    "type": "",
                    "items": {},
                    "summary": "用户手机号码注册",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "phonenum",
                            "description": "手机号码",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "form",
                            "name": "password",
                            "description": "用户密码",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "用户登录成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_TokenArray"
                        },
                        {
                            "code": 400,
                            "message": "用户注册失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/user/userinfo",
            "description": "拉取用户的基本信息",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "userinfo",
                    "type": "app-server.models.S2C_UserData",
                    "items": {},
                    "summary": "拉取用户的基本信息",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "updatetime",
                            "description": "更新时间",
                            "dataType": "int64",
                            "type": "int64",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "拉取用户数据成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_UserData"
                        },
                        {
                            "code": 400,
                            "message": "更新失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/user/resetpwd",
            "description": "重置密码",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "resetpassword",
                    "type": "",
                    "items": {},
                    "summary": "重置密码",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "password",
                            "description": "新的密码",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "form",
                            "name": "flag",
                            "description": "操作类型 1=找回密码 0=修改密码",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "form",
                            "name": "code",
                            "description": "验证码",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "重置密码成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_TokenArray"
                        },
                        {
                            "code": 400,
                            "message": "重置密码失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/user/verifypwd",
            "description": "验证密码",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "verifypassword",
                    "type": "",
                    "items": {},
                    "summary": "验证密码",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "password",
                            "description": "原密码",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "验证密码成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "验证密码失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/user/newtoken",
            "description": "刷新token，用refresh token来刷新",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "newtoken",
                    "type": "app-server.models.S2C_TokenArray",
                    "items": {},
                    "summary": "刷新token，用refresh token来刷新",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "refresh_token",
                            "description": "刷新token",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "换取新的token成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_TokenArray"
                        },
                        {
                            "code": 400,
                            "message": "换取token失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/user/account",
            "description": "设置红包号，只能设置一次",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "account",
                    "type": "",
                    "items": {},
                    "summary": "设置红包号，只能设置一次",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "account",
                            "description": "红包账号名",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "设置红包账号成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "设置红包账号失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/user/phone",
            "description": "绑定手机号并设置密码",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "bindphone",
                    "type": "",
                    "items": {},
                    "summary": "绑定手机号并设置密码",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "phonenum",
                            "description": "手机号",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "form",
                            "name": "password",
                            "description": "密码",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "绑定成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "绑定失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/user/resetphone",
            "description": "更换手机号码",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "resetphone",
                    "type": "",
                    "items": {},
                    "summary": "更换手机号码",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "phonenum",
                            "description": "手机号",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "form",
                            "name": "password",
                            "description": "密码",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "form",
                            "name": "code",
                            "description": "手机验证码",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "更换成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "跟换失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/user/nickname",
            "description": "更改用户昵称",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "nickname",
                    "type": "",
                    "items": {},
                    "summary": "更改用户昵称",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "nickname",
                            "description": "用户昵称",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "修改昵称成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "修改昵称失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/user/sex",
            "description": "设置性别",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "area",
                    "type": "",
                    "items": {},
                    "summary": "设置性别",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "sex",
                            "description": "性别1=男0=女",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "设置性别成功",
                            "responseType": "string",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/user/area",
            "description": "设置区域",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "area",
                    "type": "",
                    "items": {},
                    "summary": "设置区域",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "area",
                            "description": "所在区域",
                            "dataType": "app-server.models.AreaInfo",
                            "type": "app-server.models.AreaInfo",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "设置区域成功",
                            "responseType": "objson",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "设置区域失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/user/signature",
            "description": "设置个性签名",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "signature",
                    "type": "",
                    "items": {},
                    "summary": "设置个性签名",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "signature",
                            "description": "个性签名",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "设置个性签名成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "设置个性签名失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/user/portrait",
            "description": "设置个人头像",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "portrait",
                    "type": "",
                    "items": {},
                    "summary": "设置个人头像",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "portrait_key",
                            "description": "七牛上的文件key",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "设置头像成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "设置头像失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/user/rank",
            "description": "排行榜",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "rankList",
                    "type": "app-server.models.S2C_RankList",
                    "items": {},
                    "summary": "排行榜",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "ranktype",
                            "description": "排行榜类型[1=红包榜 2=好友榜 3=等级榜 4=老板榜]",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "page",
                            "description": "页数[0开始]",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_RankList"
                        },
                        {
                            "code": 400,
                            "message": "失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/user/sysnotice",
            "description": "系统通知列表",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "noticeList",
                    "type": "app-server.models.S2C_SysNoticeList",
                    "items": {},
                    "summary": "系统通知列表",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "page",
                            "description": "页数[0开始]",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_SysNoticeList"
                        },
                        {
                            "code": 400,
                            "message": "失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "app-server.models.APIError": {
            "id": "app-server.models.APIError",
            "properties": {
                "code": {
                    "type": "uint32",
                    "description": "错误码",
                    "items": {},
                    "format": ""
                },
                "msg": {
                    "type": "string",
                    "description": "错误描述",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.AreaInfo": {
            "id": "app-server.models.AreaInfo",
            "properties": {
                "city": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "country": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "district": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "province": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.RankItem": {
            "id": "app-server.models.RankItem",
            "properties": {
                "id": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "name": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "rank": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "score": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.S2C_RankList": {
            "id": "app-server.models.S2C_RankList",
            "properties": {
                "list": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "app-server.models.RankItem"
                    },
                    "format": ""
                },
                "self": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "type": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.S2C_SysNoticeList": {
            "id": "app-server.models.S2C_SysNoticeList",
            "properties": {
                "list": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "app-server.models.SysNoticeItem"
                    },
                    "format": ""
                }
            }
        },
        "app-server.models.S2C_TokenArray": {
            "id": "app-server.models.S2C_TokenArray",
            "properties": {
                "access_token": {
                    "type": "app-server.models.TokenInfo",
                    "description": "access token信息",
                    "items": {},
                    "format": ""
                },
                "refresh_token": {
                    "type": "app-server.models.TokenInfo",
                    "description": "refresh token信息",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.S2C_UserData": {
            "id": "app-server.models.S2C_UserData",
            "properties": {
                "account": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "area": {
                    "type": "app-server.models.AreaInfo",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "cert": {
                    "type": "uint8",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "level": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "money": {
                    "type": "int64",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "nickname": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "oauth": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "type": "int"
                    },
                    "format": ""
                },
                "phone": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "point": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "portrait": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "sex": {
                    "type": "uint8",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "signature": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "temp_money": {
                    "type": "int64",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "update_time": {
                    "type": "int64",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.SysNoticeItem": {
            "id": "app-server.models.SysNoticeItem",
            "properties": {
                "flag": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "text": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "time": {
                    "type": "int64",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.TokenInfo": {
            "id": "app-server.models.TokenInfo",
            "properties": {
                "Expiresin": {
                    "type": "uint32",
                    "description": "token的有效期，以s为单位",
                    "items": {},
                    "format": ""
                },
                "Token": {
                    "type": "string",
                    "description": "token字符串",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`, "backpay": `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://192.168.1.106:8080/v1",
    "resourcePath": "/backpay",
    "apis": [
        {
            "path": "/backpay/common",
            "description": "提现（汇款/银联/支付宝）",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "backpayCommon",
                    "type": "string",
                    "items": {},
                    "summary": "提现（汇款/银联/支付宝）",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "money",
                            "description": "要提现的金额",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "form",
                            "name": "passwd",
                            "description": "密码",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "form",
                            "name": "type",
                            "description": "提现类型，1=汇款，2=银联，3=支付宝",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "form",
                            "name": "account",
                            "description": "汇款/银联：卡号，支付宝：账号",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "form",
                            "name": "name",
                            "description": "汇款/银联：户名，支付宝：实名",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "form",
                            "name": "bankname",
                            "description": "汇款/银联：开户银行，支付宝：空",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "错误",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "app-server.models.APIError": {
            "id": "app-server.models.APIError",
            "properties": {
                "code": {
                    "type": "uint32",
                    "description": "错误码",
                    "items": {},
                    "format": ""
                },
                "msg": {
                    "type": "string",
                    "description": "错误描述",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`, "oauth": `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://192.168.1.106:8080/v1",
    "resourcePath": "/oauth",
    "apis": [
        {
            "path": "/oauth/weixin/login",
            "description": "使用微信登陆",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "loginByweixin",
                    "type": "",
                    "items": {},
                    "summary": "使用微信登陆",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "openid",
                            "description": "微信openid",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "form",
                            "name": "token",
                            "description": "微信access_token",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "登陆成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_TokenArray"
                        },
                        {
                            "code": 400,
                            "message": "登陆失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/oauth/qq/login",
            "description": "使用QQ登陆",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "loginbyQQ",
                    "type": "",
                    "items": {},
                    "summary": "使用QQ登陆",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "openid",
                            "description": "qq openid",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "form",
                            "name": "openkey",
                            "description": "qq openkey",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "登陆成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_TokenArray"
                        },
                        {
                            "code": 400,
                            "message": "登陆失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/oauth/weibo/login",
            "description": "使用微博登陆",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "loginbyWebo",
                    "type": "",
                    "items": {},
                    "summary": "使用微博登陆",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "access_token",
                            "description": "weibo的access token",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "登陆成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_TokenArray"
                        },
                        {
                            "code": 400,
                            "message": "登陆失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/oauth/bindweixin",
            "description": "绑定微信账号",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "bindweixin",
                    "type": "",
                    "items": {},
                    "summary": "绑定微信账号",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "openid",
                            "description": "微信openid",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "form",
                            "name": "token",
                            "description": "微信access_token",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "绑定成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "绑定失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/oauth/bindqq",
            "description": "绑定QQ账号",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "bindqq",
                    "type": "",
                    "items": {},
                    "summary": "绑定QQ账号",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "openid",
                            "description": "qq openid",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "form",
                            "name": "openkey",
                            "description": "qq openkey",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "绑定成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "绑定失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/oauth/bindweibo",
            "description": "绑定微博账号",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "bindWebo",
                    "type": "",
                    "items": {},
                    "summary": "绑定微博账号",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "access_token",
                            "description": "weibo的access token",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "绑定成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "绑定失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/oauth/unbind",
            "description": "解绑第三方账号",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "unbind",
                    "type": "",
                    "items": {},
                    "summary": "解绑第三方账号",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "platform",
                            "description": "第三方平台[0=微信,1=微博,2=QQ]",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "解绑成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "绑定失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "app-server.models.APIError": {
            "id": "app-server.models.APIError",
            "properties": {
                "code": {
                    "type": "uint32",
                    "description": "错误码",
                    "items": {},
                    "format": ""
                },
                "msg": {
                    "type": "string",
                    "description": "错误描述",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.S2C_TokenArray": {
            "id": "app-server.models.S2C_TokenArray",
            "properties": {
                "access_token": {
                    "type": "app-server.models.TokenInfo",
                    "description": "access token信息",
                    "items": {},
                    "format": ""
                },
                "refresh_token": {
                    "type": "app-server.models.TokenInfo",
                    "description": "refresh token信息",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.TokenInfo": {
            "id": "app-server.models.TokenInfo",
            "properties": {
                "Expiresin": {
                    "type": "uint32",
                    "description": "token的有效期，以s为单位",
                    "items": {},
                    "format": ""
                },
                "Token": {
                    "type": "string",
                    "description": "token字符串",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`, "qiniu": `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://192.168.1.106:8080/v1",
    "resourcePath": "/qiniu",
    "apis": [
        {
            "path": "/qiniu/uploadtoken",
            "description": "获取七牛上传凭证",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "uploadtoken",
                    "type": "app-server.models.S2C_QiniuUpToken",
                    "items": {},
                    "summary": "获取七牛上传凭证",
                    "parameters": [
                        {
                            "paramType": "string",
                            "name": "key",
                            "description": "key",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "获取七牛上传凭证成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_QiniuUpToken"
                        },
                        {
                            "code": 400,
                            "message": "获取七牛上传凭证失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/qiniu/privateuploadtoken",
            "description": "获取七牛私人空间上传凭证",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "privateuploadtoken",
                    "type": "app-server.models.S2C_QiniuUpToken",
                    "items": {},
                    "summary": "获取七牛私人空间上传凭证",
                    "parameters": [
                        {
                            "paramType": "string",
                            "name": "key",
                            "description": "key",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "获取七牛上传凭证成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_QiniuUpToken"
                        },
                        {
                            "code": 400,
                            "message": "获取七牛上传凭证失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/qiniu/downloadtoken",
            "description": "获取七牛下载凭证(私有空间才需要)",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "downloadtoken",
                    "type": "app-server.models.S2C_QiniuDlUrl",
                    "items": {},
                    "summary": "获取七牛下载凭证(私有空间才需要)",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "key",
                            "description": "图片的key",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "获取七牛下载凭证成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_QiniuDlUrl"
                        },
                        {
                            "code": 400,
                            "message": "获取七牛下载凭证失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/qiniu/uptest",
            "description": "上传测试(暂时废弃)",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "uptest",
                    "type": "",
                    "items": {},
                    "summary": "上传测试(暂时废弃)",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "token",
                            "description": "七牛上传凭证",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "上传测试成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "上传测试失败",
                            "responseType": "string",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "app-server.models.APIError": {
            "id": "app-server.models.APIError",
            "properties": {
                "code": {
                    "type": "uint32",
                    "description": "错误码",
                    "items": {},
                    "format": ""
                },
                "msg": {
                    "type": "string",
                    "description": "错误描述",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.S2C_QiniuDlUrl": {
            "id": "app-server.models.S2C_QiniuDlUrl",
            "properties": {
                "dl_url": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "expires": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.S2C_QiniuUpToken": {
            "id": "app-server.models.S2C_QiniuUpToken",
            "properties": {
                "expires": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "upload_token": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`, "socially": `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://192.168.1.106:8080/v1",
    "resourcePath": "/socially",
    "apis": [
        {
            "path": "/socially/rctoken",
            "description": "获取融云token",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "rcToken",
                    "type": "app-server.models.S2C_RcToken",
                    "items": {},
                    "summary": "获取融云token",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "refresh",
                            "description": "是否需要向融云服务器获取新的token（一般填0，注意是字符串）",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "融云token",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_RcToken"
                        },
                        {
                            "code": 400,
                            "message": "获取融云token失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/socially/friend/list",
            "description": "获取好友列表",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "list",
                    "type": "app-server.models.FriendList",
                    "items": {},
                    "summary": "获取好友列表",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "返回好友列表",
                            "responseType": "object",
                            "responseModel": "app-server.models.FriendList"
                        },
                        {
                            "code": 400,
                            "message": "获取好友列表失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/socially/friend/info",
            "description": "获取好友信息",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "friendInfo",
                    "type": "app-server.models.S2C_UserData",
                    "items": {},
                    "summary": "获取好友信息",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "targetid",
                            "description": "要查询的用户id",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "query",
                            "name": "updatetime",
                            "description": "更新时间",
                            "dataType": "int64",
                            "type": "int64",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "返回好友信息",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_UserData"
                        },
                        {
                            "code": 400,
                            "message": "返回好友信息失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/socially/friend/add",
            "description": "请求加好友",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "add",
                    "type": "",
                    "items": {},
                    "summary": "请求加好友",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "targetid",
                            "description": "要添加的用户id",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "发送添加请求成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "发送添加请求失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/socially/friend/addbysearch",
            "description": "通过红包号、手机号搜素加好友",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "add",
                    "type": "",
                    "items": {},
                    "summary": "通过红包号、手机号搜素加好友",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "condition",
                            "description": "要添加的手机号或红包号",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "发送添加请求成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "发送添加请求失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/socially/friend/agree",
            "description": "同意加好友",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "agree",
                    "type": "",
                    "items": {},
                    "summary": "同意加好友",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "userid",
                            "description": "发起请求的用户id",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "form",
                            "name": "name",
                            "description": "自己的昵称",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "form",
                            "name": "portrait",
                            "description": "自己的图像",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "加好友失败",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "请求失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/socially/friend/refuse",
            "description": "拒绝添加好友请求",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "refuse",
                    "type": "",
                    "items": {},
                    "summary": "拒绝添加好友请求",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "fromid",
                            "description": "要拒绝的用户id",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "成功拒绝",
                            "responseType": "object",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/socially/friend/remove",
            "description": "解除好友关系",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "remove",
                    "type": "",
                    "items": {},
                    "summary": "解除好友关系",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "friendid",
                            "description": "好友用户id",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "解除成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "删除好友失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/socially/blacklist/add",
            "description": "解除好友关系",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "addblacklist",
                    "type": "",
                    "items": {},
                    "summary": "解除好友关系",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "targetid",
                            "description": "用户id",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 201,
                            "message": "加入成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "删除好友失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "app-server.models.APIError": {
            "id": "app-server.models.APIError",
            "properties": {
                "code": {
                    "type": "uint32",
                    "description": "错误码",
                    "items": {},
                    "format": ""
                },
                "msg": {
                    "type": "string",
                    "description": "错误描述",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.AreaInfo": {
            "id": "app-server.models.AreaInfo",
            "properties": {
                "city": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "country": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "district": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "province": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.FriendBrief": {
            "id": "app-server.models.FriendBrief",
            "properties": {
                "black": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "name": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "portrait": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "star": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "userid": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.FriendList": {
            "id": "app-server.models.FriendList",
            "properties": {
                "friends": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "$ref": "app-server.models.FriendBrief"
                    },
                    "format": ""
                }
            }
        },
        "app-server.models.S2C_RcToken": {
            "id": "app-server.models.S2C_RcToken",
            "properties": {
                "RcToken": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.S2C_UserData": {
            "id": "app-server.models.S2C_UserData",
            "properties": {
                "account": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "area": {
                    "type": "app-server.models.AreaInfo",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "cert": {
                    "type": "uint8",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "level": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "money": {
                    "type": "int64",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "nickname": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "oauth": {
                    "type": "array",
                    "description": "",
                    "items": {
                        "type": "int"
                    },
                    "format": ""
                },
                "phone": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "point": {
                    "type": "int",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "portrait": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "sex": {
                    "type": "uint8",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "signature": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "temp_money": {
                    "type": "int64",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "update_time": {
                    "type": "int64",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`, "enterprise_cert": `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://192.168.1.106:8080/v1",
    "resourcePath": "/enterprise_cert",
    "apis": [
        {
            "path": "/enterprise_cert/submit",
            "description": "认证/重新认证",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "submitEnterpriseCertMaterial",
                    "type": "string",
                    "items": {},
                    "summary": "认证/重新认证",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "data",
                            "description": "提交材料",
                            "dataType": "app-server.models.EnterpriseCertMaterialBinding",
                            "type": "app-server.models.EnterpriseCertMaterialBinding",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/enterprise_cert/edit",
            "description": "修改认证信息",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "editEnterpriseCertInfo",
                    "type": "string",
                    "items": {},
                    "summary": "修改认证信息",
                    "parameters": [
                        {
                            "paramType": "body",
                            "name": "data",
                            "description": "修改后的信息",
                            "dataType": "app-server.models.EnterpriseCertInfoBinding",
                            "type": "app-server.models.EnterpriseCertInfoBinding",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "成功",
                            "responseType": "object",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/enterprise_cert/info",
            "description": "获取认证信息",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "getEnterpriseCertInfo",
                    "type": "app-server.models.S2C_EnterpriseCertInfo",
                    "items": {},
                    "summary": "获取认证信息",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "成功",
                            "responseType": "object",
                            "responseModel": "app-server.models.S2C_EnterpriseCertInfo"
                        },
                        {
                            "code": 400,
                            "message": "失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "app-server.models.APIError": {
            "id": "app-server.models.APIError",
            "properties": {
                "code": {
                    "type": "uint32",
                    "description": "错误码",
                    "items": {},
                    "format": ""
                },
                "msg": {
                    "type": "string",
                    "description": "错误描述",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.EnterpriseCertInfoBinding": {
            "id": "app-server.models.EnterpriseCertInfoBinding",
            "properties": {
                "official_website": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "weibo": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "weixin": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.EnterpriseCertMaterialBinding": {
            "id": "app-server.models.EnterpriseCertMaterialBinding",
            "properties": {
                "business_license": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "enterprise_name": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "name": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "official_website": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "operate_id": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "operate_id_photo": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "operate_name": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "operate_place": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "trademark_cert": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "weibo": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "weixin": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "app-server.models.S2C_EnterpriseCertInfo": {
            "id": "app-server.models.S2C_EnterpriseCertInfo",
            "properties": {
                "cert_name": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "official_website": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "weibo": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "weixin": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`, "pub": `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://192.168.1.106:8080/v1",
    "resourcePath": "/pub",
    "apis": [
        {
            "path": "/pub/servertime",
            "description": "获取服务器时间",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "getServerTime",
                    "type": "string",
                    "items": {},
                    "summary": "获取服务器时间",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "获取服务器时间成功",
                            "responseType": "string",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/pub/test",
            "description": "公共接口测试",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "publiceTest",
                    "type": "string",
                    "items": {},
                    "summary": "公共接口测试",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "测试成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 401,
                            "message": "测试失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/pub/posttest",
            "description": "公共接口post测试",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "postTest",
                    "type": "string",
                    "items": {},
                    "summary": "公共接口post测试",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "测试成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 401,
                            "message": "测试失败",
                            "responseType": "string",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/pub/debug",
            "description": "调试接口",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "debug",
                    "type": "string",
                    "items": {},
                    "summary": "调试接口",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "测试成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 401,
                            "message": "测试失败",
                            "responseType": "string",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "app-server.models.APIError": {
            "id": "app-server.models.APIError",
            "properties": {
                "code": {
                    "type": "uint32",
                    "description": "错误码",
                    "items": {},
                    "format": ""
                },
                "msg": {
                    "type": "string",
                    "description": "错误描述",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`, "unionpay": `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://192.168.1.106:8080/v1",
    "resourcePath": "/unionpay",
    "apis": [
        {
            "path": "/unionpay/notify",
            "description": "银联异步回调api",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "unionpay_notify",
                    "type": "string",
                    "items": {},
                    "summary": "银联异步回调api",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "回调成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "回调失败",
                            "responseType": "string",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/unionpay/backpaynotify",
            "description": "银联提现回调api",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "unionpay_backpay",
                    "type": "string",
                    "items": {},
                    "summary": "银联提现回调api",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "回调成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "回调失败",
                            "responseType": "string",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/unionpay/tradeno",
            "description": "获取银联交易流水号",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "TradeNo",
                    "type": "string",
                    "items": {},
                    "summary": "获取银联交易流水号",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "redpktid",
                            "description": "要付款的红包id",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "获取成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "获取失败",
                            "responseType": "object",
                            "responseModel": "app-server.models.APIError"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/unionpay/tradestatus",
            "description": "查询订单交易状态",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "TradeStatus",
                    "type": "string",
                    "items": {},
                    "summary": "查询订单交易状态",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "redpktid",
                            "description": "要付款的红包id",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "失败",
                            "responseType": "string",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/unionpay/backpay",
            "description": "请求提现",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "TradeStatus",
                    "type": "string",
                    "items": {},
                    "summary": "请求提现",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "money",
                            "description": "要提现的金额",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "成功",
                            "responseType": "string",
                            "responseModel": "string"
                        },
                        {
                            "code": 400,
                            "message": "失败",
                            "responseType": "string",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "app-server.models.APIError": {
            "id": "app-server.models.APIError",
            "properties": {
                "code": {
                    "type": "uint32",
                    "description": "错误码",
                    "items": {},
                    "format": ""
                },
                "msg": {
                    "type": "string",
                    "description": "错误描述",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`}
