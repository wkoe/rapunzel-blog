package storage

import (
	"github.com/s4kibs4mi/rapunzel-blog/models"
	"github.com/s4kibs4mi/rapunzel-blog/protos"
	"github.com/s4kibs4mi/rapunzel-blog/conn"
	"gopkg.in/mgo.v2/bson"
)

/**
 * := Coded with love by Sakib Sami on 20/1/18.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

type MongodbStorage struct {
}

func (db *MongodbStorage) Init() bool {
	return conn.NewMongodbConnection()
}

func (db *MongodbStorage) Save(user models.User) bool {
	if err := conn.GetUserCollection().Insert(user); err != nil {
		return false
	}
	return true
}

func (db *MongodbStorage) Update(user models.User) bool {
	if err := conn.GetUserCollection().Update(user, user); err != nil {
		return false
	}
	return true
}

func (db *MongodbStorage) Count() int {
	if n, err := conn.GetUserCollection().Count(); err == nil {
		return n
	}
	return 0
}

func (db *MongodbStorage) Delete(user models.User) bool {
	if err := conn.GetUserCollection().Remove(user); err != nil {
		return false
	}
	return true
}

func (db *MongodbStorage) FindByID(ID string) *models.User {
	u := models.User{}
	if err := conn.GetUserCollection().FindId(ID).One(&u); err != nil {
		return nil
	}
	return &u
}

func (db *MongodbStorage) FindByUsername(username string) *models.User {
	u := models.User{}
	if err := conn.GetUserCollection().Find(bson.M{
		"username": username,
	}).One(&u); err != nil {
		return nil
	}
	return &u
}

func (db *MongodbStorage) FindByEmail(email string) *models.User {
	u := models.User{}
	if err := conn.GetUserCollection().Find(bson.M{
		"email": email,
	}).One(&u); err != nil {
		return nil
	}
	return &u
}

func (db *MongodbStorage) FindAll() []models.User {
	var users []models.User
	if err := conn.GetUserCollection().Find(bson.M{

	}).All(&users); err != nil {
		return nil
	}
	return users
}

func (db *MongodbStorage) FindAllByQuery(query protos.Query) []models.User {
	var users []models.User
	if err := conn.GetUserCollection().Find(bson.M{
		query.Field: query.Value,
	}).All(&users); err != nil {
		return nil
	}
	return users
}