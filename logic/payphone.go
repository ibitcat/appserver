package logic

import (
	"fmt"

	"app-server/define"
	"app-server/models"
	"app-server/pkg/mongodb"

	"gopkg.in/mgo.v2/bson"
)

/////////////////////////////////////////////////////////
// 话费充值数据库查询相关
/////////////////////////////////////////////////////////
// 根据userid获取充值记录
func FindPayPhoneRecordById(tradeno string) (*models.PayPhoneRecord, error) {
	if len(tradeno) == 0 {
		return nil, fmt.Errorf("tradeno invaild")
	}

	var record models.PayPhoneRecord
	err := mongodb.SelectById(define.PayPhoneRecordCollection, bson.ObjectIdHex(tradeno), nil, &record)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

// 完成记录充值状态
func UpdatePayPhoneRecordFinish(tradeno string) error {
	if len(tradeno) == 0 {
		return fmt.Errorf("tradeno invaild")
	}

	selector := bson.M{"_id": bson.ObjectIdHex(tradeno)}
	update := bson.M{"$set": bson.M{"status": 1}}
	return mongodb.Update(define.PayPhoneRecordCollection, selector, update)
}
