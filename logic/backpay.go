/**
 * 提现
 */
package logic

import (
	"app-server/define"
	"app-server/models"
	"app-server/pkg/mongodb"

	"gopkg.in/mgo.v2/bson"
)

// 创建提现记录
func CreateBackpayRecord(record *models.BackpayRecord) error {
	return mongodb.Insert(define.BackpayCollection, record)
}

// 提现处理完成
func HandleBackpayRecord(id string, success bool) error {
	var update bson.M
	if success {
		update = bson.M{"$set": bson.M{"status": 1}}
	} else {
		update = bson.M{"$set": bson.M{"status": 2}}
	}
	return mongodb.UpdateById(define.BackpayCollection, bson.ObjectIdHex(id), update)
}
