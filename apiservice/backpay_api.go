package apiservice

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"app-server/logic"
	"app-server/models"
	"app-server/pkg/utils"

	"gopkg.in/mgo.v2/bson"
)

func BackPayCommon(request *http.Request, userId string) uint32 {
	user, err := logic.FindUserDataById(userId)
	if err != nil {
		return 10451
	}
	request.ParseForm()

	money, err := strconv.ParseFloat(request.FormValue("money"), 64)
	if err != nil {
		fmt.Println("无效金额", request.FormValue("money"))
		return 10451
	}

	passwd := request.FormValue("passwd")
	if utils.Md5(passwd+user.Salt) != user.Password {
		return 10452
	}

	// 余额不足
	if user.Money < int64(money*100) {
		fmt.Println("超出剩余金额", user.Money, money*100)
		return 10451
	}

	typ, err := strconv.Atoi(request.FormValue("type"))
	if err != nil {
		fmt.Println("无效类型", request.FormValue("type"))
		return 10451
	}

	tradeNo := bson.NewObjectId()
	record := &models.BackpayRecord{
		Id_:    tradeNo,
		UserId: userId,
		Fee:    int64(money * 100),
		Time:   time.Now().Unix(),
	}

	switch typ {
	default:
		fmt.Println("无效类型", typ)
		return 10451
	case 1: // 汇款
		record.Type = 1
		record.Account = request.FormValue("account")
		record.Name = request.FormValue("name")
		record.BankName = request.FormValue("bankname")
	case 2: // 银联
		record.Type = 2
		record.Account = request.FormValue("account")
		record.Name = request.FormValue("name")
		record.BankName = request.FormValue("bankname")
	case 3: // 支付宝
		record.Type = 3
		record.Account = request.FormValue("account")
		record.Name = request.FormValue("name")
		//case 4: // 微信
		//	record.Type = 4
		//	record.Name = request.FormValue("name")
	}

	logic.UpdateMoney(userId, -int64(money*100))
	err = logic.CreateBackpayRecord(record)
	if err != nil {
		fmt.Println("mongodb错误", err.Error())
		return 10451
	}
	return 0
}
