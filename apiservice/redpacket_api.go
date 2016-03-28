package apiservice

import (
	"fmt"
	"net/http"
	"strconv"

	"app-server/define"
	"app-server/logic"
	"app-server/models"
	"app-server/pkg/mongodb"
	"app-server/pkg/token"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// 获取红包列表
func GetRedpacketList(index string, req *http.Request) *models.S2C_RedpacketList {
	startIdx, err := strconv.ParseUint(index, 10, 32)
	if err != nil {
		return nil
	}

	var userId string
	jwtBackend := token.GetJwtBackend()
	tokenObj, err := jwtBackend.ParseToken(req)
	if err == nil {
		userId = jwtBackend.GetUserIdFromToken(tokenObj)
	}

	idList := logic.GetRedpktIdList(uint32(startIdx))
	if len(idList) == 0 {
		return nil
	}

	requester, _ := logic.GetUserData(userId)
	retList := &models.S2C_RedpacketList{}
	retList.List = make([]*models.RedpacketInfo, 0, 10)
	for _, id := range idList {
		con := logic.GetRedpacketCon(id)
		if con != nil {
			redpkt := con.GetRedpacketData()
			sender, getErr := logic.GetUserData(redpkt.SenderId)
			if getErr != nil {
				fmt.Println("[ERROR] 找不到发送者------>", redpkt.SenderId)
				continue
			}

			// 筛选区域
			if requester != nil {
				// if len(redpkt.Area) > 0 && redpkt.Area != "全国" && requester.Area != redpkt.Area {
				// 	continue
				// }
			}

			info := new(models.RedpacketInfo)
			info.Id = id
			info.BeginTime = redpkt.BeginTime
			info.EndTime = redpkt.EndTime
			info.UserId = redpkt.SenderId // 设置红包发送者信息
			info.UserName = sender.NickName
			info.IsAuth = sender.Cert
			info.PerMoney = redpkt.PerMoney - uint32(float32(redpkt.PerMoney)*(float32(redpkt.Rebate)/100.0))
			info.Number = redpkt.Remainder
			info.Type = redpkt.Type
			info.Area = redpkt.Area

			switch info.Type {
			case 1:
				info.Share = &redpkt.Share
			case 2, 3:
				info.App = &redpkt.App
			case 4:
				info.OfficialAcc = &redpkt.OfficialAcc
			}

			retList.List = append(retList.List, info)
		}
	}

	retList.Count = uint32(len(retList.List))
	return retList
}

// 创建红包
func CreateRedpacket(c *gin.Context) (string, error) {
	var b models.SendRedpacketBinding
	err := c.BindJSON(&b)
	fmt.Println("err = ", err)
	if err != nil {
		return "", err
	}

	userId := c.MustGet("userId").(string)
	return logic.CreateRedpacket(userId, &b)
}

// 抢红包
func GrabRedpacket(userId, redpacketId, deviceid string) uint32 {
	return logic.GrabRedpacket(userId, redpacketId, deviceid)
}

// 红包任务进行中
func FinishShare(userId, redpacketId, deviceid string) uint32 {
	return logic.DoingRedpacket(userId, redpacketId, deviceid)
}

// 完成红包任务
func FinishRedpacket(userId, redpacketId, deviceid string) uint32 {
	ecode := logic.FinishRedpacket(userId, redpacketId, deviceid)
	if ecode == 0 {
		perMoeny := logic.GetRedpacketPerMoney(redpacketId)
		fmt.Println("单个红包的金额 = ", perMoeny)
		logic.UpdateMoney(userId, int64(perMoeny))
	}

	return ecode
}

// 放弃红包
func GiveupRedpacket(userId, redpacketId string) {
	logic.GiveupRedpacket(userId, redpacketId)
}

// 红包余额支付
func PayRedpacketByBalance(userId, redpacketId string) bool {
	return logic.PayRedpacketByBalance(userId, redpacketId)
}

// 第三方支付
func PayRedpacketBy3rdParty(redpktId string, fee int) bool {
	return logic.PayRedpacketBy3rdParty(redpktId, fee)
}

// 过滤红包
func FilterRedpacket(userId string) *models.UserExpireList {
	return logic.Filter(userId)
}

// 红包领取记录
func GetRedpktRecordList(redpktId, cursor string) *models.S2C_RedpktRecord {
	cursorIdx, err := strconv.Atoi(cursor)
	if err != nil {
		return nil
	}

	recordList := &models.S2C_RedpktRecord{}
	recordList.Id = redpktId
	userList := logic.ScanRecordList(redpktId, cursorIdx)
	if len(userList) > 0 {
		recordList.List = make([]*models.RedpktRecord, 0, 10)
		for userId, tm := range userList {
			user, getErr := logic.GetUserData(userId)
			if getErr != nil {
				continue
			}

			record := new(models.RedpktRecord)
			record.UserId = userId
			record.UserName = user.NickName
			record.Time = tm
			recordList.List = append(recordList.List, record)
		}
	}

	return recordList
}

// 审核游戏或app红包
func VerifyAppRedpacket(redpktId, status string) {
	result, _ := strconv.Atoi(status)
	if result == 1 || result == 2 {
		logic.UpdateVerify(redpktId, result)
	}
}

// 获取收到的红包信息，发出的红包个数和排名等
func GetRedpketRecieveInfo(userId string) *models.S2C_RedpketRecieveInfo {
	userData, err := logic.GetUserData(userId)
	if err != nil {
		return nil
	}

	info := new(models.S2C_RedpketRecieveInfo)
	info.Total = userData.Income
	info.Rank = logic.GetIncomeRank(userId)

	return info
}

// 收到的红包列表
func GetRedpketRecieveList(userId, date string) *models.S2C_ReceivedList {
	query := bson.M{"userid": userId, "grab_date": date}
	sort := "-grab_time"
	fields := bson.M{"_id": 0, "userid": 0, "user_name": 0, "grab_date": 0}

	var list []*models.GrabRecord
	err := mongodb.SelectAllWithParam(define.GrabRecordCollection, query, sort, fields, 0, 10, &list)
	if err != nil {
		return nil
	}

	record := &models.S2C_ReceivedList{}
	record.List = list
	return record
}

// 获取发出的红包信息，个数和总金额
func GetRedpketSendInfo(userId, year string) *models.S2C_RedpketSendInfo {
	userData, err := logic.GetUserData(userId)
	if err != nil {
		return nil
	}

	info := models.S2C_RedpketSendInfo{}
	if outcome, ok := userData.Outcome[year]; ok {
		if len(outcome) == 2 {
			info.Total = outcome[0]
			info.Amount = outcome[1]
			return &info
		}
	}

	return &info
}

// 获取发出的红包列表
func GetRedpketSendList(userId, year string) *models.S2C_RedpktSendList {
	query := bson.M{"sender_id": userId, "year": year}
	sort := "-create_time"

	var list []*models.RedPacket
	err := mongodb.SelectAllWithParam(define.RedpacketCollection, query, sort, nil, 0, 10, &list)
	if err != nil {
		return nil
	}

	record := &models.S2C_RedpktSendList{}
	for _, v := range list {
		info := new(models.SendRedpacket)
		info.Id = v.Id_.Hex()
		info.BeginTime = v.BeginTime
		info.EndTime = v.EndTime
		info.Title = ""
		info.PerMoney = v.PerMoney
		info.Total = v.Total
		info.Remainder = v.Remainder
		switch v.Type {
		case define.ERedpkt_Share:
			info.Title = v.Share.Title
		case define.ERedpkt_App, define.ERedpkt_Game:
			info.Title = v.App.Name
		case define.ERedpkt_OA:
			info.Title = v.OfficialAcc.Title
		}
		record.List = append(record.List, info)
	}
	return record
}

// 获取待发布红包列表
func GetTobeReleasedRedpktList(userId string) *models.S2C_ToBeReleasedList {
	condition := []bson.M{bson.M{"verify": bson.M{"$ne": 1}}, bson.M{"trade_status": 0}}
	query := bson.M{"sender_id": userId, "$or": condition}
	sort := "-create_time"

	var list []*models.RedPacket
	err := mongodb.SelectAllWithParam(define.RedpacketCollection, query, sort, nil, 0, 0, &list)
	if err != nil {
		return nil
	}

	record := &models.S2C_ToBeReleasedList{}
	record.List = make([]*models.ToBeReleasedRedpkt, 0, 10)
	for _, v := range list {
		info := new(models.ToBeReleasedRedpkt)
		info.Id = v.Id_.Hex()
		info.CreateTime = v.CreateTime
		info.BeginTime = v.BeginTime
		info.Verify = v.Verify
		info.TradeStatus = v.TradeStatus
		info.Title = ""
		switch v.Type {
		case define.ERedpkt_Share:
			info.Title = v.Share.Title
		case define.ERedpkt_App, define.ERedpkt_Game:
			info.Title = v.App.Name
		case define.ERedpkt_OA:
			info.Title = v.OfficialAcc.Title
		}
		record.List = append(record.List, info)
	}
	return record
}

// 获取红包的统计数据
func GetRedpacketStatistics(userId, redpktId string) *models.S2C_RedpktStatistics {
	var data models.S2C_RedpktStatistics
	data.Id = redpktId

	statistics := logic.GetStatistics(redpktId)
	if statistics != nil {
		data.RedpktStatistics = *statistics
	}

	return &data
}
