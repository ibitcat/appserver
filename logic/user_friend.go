// 用户好友相关操作

package logic

import (
	"time"

	"app-server/define"
	"app-server/models"
	"app-server/pkg/mongodb"

	"gopkg.in/mgo.v2/bson"
)

// 将用户的最新信息更新到自己的好友列表
func UpdateNameToFriends(userId, nickName string) {
	var friends models.FriendList
	id := bson.ObjectIdHex(userId)
	feilds := bson.M{"friends": 1, "_id": 0}

	mongodb.SelectById(define.AccountCollection, id, feilds, &friends)

	for _, friend := range friends.Friends {
		selector := bson.M{
			"_id":            bson.ObjectIdHex(friend.UserId),
			"friends.userid": userId,
		}
		update := bson.M{
			"$set": bson.M{"friends.$.name": nickName},
		}

		go mongodb.Update(define.AccountCollection, selector, update)
	}
}

// 添加好友入库
func AddFriendToDB(userId, friendId, friendName, friendPor string) {
	var friend models.FriendBrief
	friend.UserId = friendId
	friend.Name = friendName
	friend.Portrait = friendPor
	friend.Time = time.Now().Unix()
	update := bson.M{"$addToSet": bson.M{"friends": friend}}
	UpdateUserById(userId, update)
}

// 是否已经是好友了
func Isfriends(fromId, targetId string) bool {
	query := bson.M{"_id": bson.ObjectIdHex(fromId), "friends.userid": targetId}
	return mongodb.Exists(define.AccountCollection, query)
}
