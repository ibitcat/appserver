package mongodb

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const url = "192.168.1.112:27017" // mongo数据库

var (
	masterSession *mgo.Session
	dataBase      = "xxlb_debug"
)

func InitMasterSession() {
	if masterSession == nil {
		var err error
		masterSession, err = mgo.Dial(url)
		if err != nil {
			log.Println("dial mongodb fail, err = ", err)
			panic(err)
		}
	}
}

// 获取session,返回的是master session的一个副本
func GetSession() *mgo.Session {
	if masterSession == nil {
		InitMasterSession()
	}

	// 最大连接数 4096
	return masterSession.Clone()
}

// 执行mongodb shell
func M(collection string, f func(*mgo.Collection) error) error {
	session := GetSession()
	defer func() {
		session.Close()
		if err := recover(); err != nil {
			log.Println("M", err)
		}
	}()

	c := session.DB(dataBase).C(collection)
	return f(c)
}

/**
 * 执行查询，此方法可拆分做为公共方法
 * [Select description]
 * @param {[type]} collectionName string [description]
 * @param {[type]} query          bson.M [description]
 * @param {[type]} sort           bson.M [description]
 * @param {[type]} fields         bson.M [description]
 * @param {[type]} skip           int    [description]
 * @param {[type]} limit          int)   (results      []interface{}, err error [description]
 */
func Select(collectionName string, query bson.M, sort string, fields bson.M, skip int, limit int) (results []interface{}, err error) {
	exop := func(c *mgo.Collection) error {
		if len(sort) == 0 {
			return c.Find(query).Select(fields).Skip(skip).Limit(limit).All(&results)
		}
		return c.Find(query).Sort(sort).Select(fields).Skip(skip).Limit(limit).All(&results)
	}
	err = M(collectionName, exop)
	return
}

// 查询所有结果
// 类似Select方法,需要传入比较详细的参数列表，注意：result 必须传递一个slice指针
func SelectAllWithParam(collectionName string, query bson.M, sort string, fields bson.M, skip int, limit int, results interface{}) error {
	exop := func(c *mgo.Collection) error {
		if len(sort) == 0 {
			return c.Find(query).Select(fields).Skip(skip).Limit(limit).All(results)
		}
		return c.Find(query).Sort(sort).Select(fields).Skip(skip).Limit(limit).All(results)
	}
	return M(collectionName, exop)
}

// 查询所有结果
// 比较常用的查询，只带field筛选条件
func SelectAll(collectionName string, query, fields bson.M, results interface{}) error {
	exop := func(c *mgo.Collection) error {
		if fields == nil {
			return c.Find(query).Select(fields).All(results)
		}
		return c.Find(query).All(results)
	}
	return M(collectionName, exop)
}

// 根据id查询结果
// 类似Select方法，注意：result 必须传递一个slice指针
func SelectById(collectionName string, id interface{}, fields bson.M, results interface{}) error {
	exop := func(c *mgo.Collection) error {
		if fields == nil {
			return c.FindId(id).One(results)
		}
		return c.FindId(id).Select(fields).One(results)
	}
	return M(collectionName, exop)
}

// 查询一个结果，如果查询结果超过一个，会报错
// 类似Select方法，注意：result 必须传递一个slice指针
func SelectOne(collectionName string, query, fields bson.M, results interface{}) error {
	exop := func(c *mgo.Collection) error {
		if fields == nil {
			return c.Find(query).Select(fields).One(results)
		}
		return c.Find(query).One(results)
	}
	return M(collectionName, exop)
}

// 查询是否存在
func Exists(collectionName string, query bson.M) bool {
	var count int
	exop := func(c *mgo.Collection) error {
		var err error
		count, err = c.Find(query).Count()
		return err
	}
	err := M(collectionName, exop)
	if err != nil {
		return false
	}

	return count > 0
}

/*
	插入文档
*/
func Insert(collectionName string, docs interface{}) error {
	query := func(c *mgo.Collection) error {
		return c.Insert(docs)
	}

	return M(collectionName, query)
}

/*
	更新文档
*/
func Update(collectionName string, selector interface{}, update interface{}) error {
	query := func(c *mgo.Collection) error {
		return c.Update(selector, update)
	}

	return M(collectionName, query)
}

/*
	通过Id更新文档
*/
func UpdateById(collectionName string, id interface{}, update interface{}) error {
	query := func(c *mgo.Collection) error {
		return c.UpdateId(id, update)
	}

	return M(collectionName, query)
}

/*
	删除文档
*/
func Delete(collectionName string, selector interface{}) error {
	query := func(c *mgo.Collection) error {
		return c.RemoveId(selector)
	}

	return M(collectionName, query)
}

///////////////////////////////////////////
// test
type Account struct {
	Id_      uint32 `bson:"_id"`
	Account  string `bson:"account"`
	Password string `bson:"password"`
}

func Test() {
	m := Account{
		Id_:      1001,
		Account:  "test001",
		Password: "123456",
	}
	query := func(c *mgo.Collection) error {
		return c.Insert(&m)
	}

	err := M("account", query)
	log.Println("插入结果：", err)
}
