package repository

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database

//COLLECTION constant
const COLLECTION = "logs"

//DAO data access object
type DAO struct {
	Server   string
	Database string
}

//Connect method
func (m *DAO) Connect() error {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
	err = session.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Mongo db connected successful!")
	return nil
}

//FindAll method
func (m *DAO) FindAll(limit, offset int) ([]Record, error) {
	var records []Record
	err := db.C(COLLECTION).Find(bson.M{}).Limit(limit).Skip(offset).Sort("log_time").All(&records)
	return records, err
}

//Count method
func (m *DAO) Count() (int, error) {
	return db.C(COLLECTION).Find(bson.M{}).Count()
}

//FindByID method
func (m *DAO) FindByID(id string) (Record, error) {
	var record Record
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&record)
	return record, err
}

//Insert method
func (m *DAO) Insert(raw Record) error {
	err := db.C(COLLECTION).Insert(&raw)
	return err
}

//Truncate collection method
func (m *DAO) Truncate() error {
	var i interface{}
	_, err := db.C(COLLECTION).RemoveAll(&i)
	return err
}

//Delete method
func (m *DAO) Delete(raw Record) error {
	err := db.C(COLLECTION).Remove(&raw)
	return err
}

//Update method
func (m *DAO) Update(raw Record) error {
	err := db.C(COLLECTION).UpdateId(raw.ID, &raw)
	return err
}
