package repository

import (
	"fmt"
	"log"
	"os"

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
func (m *DAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
	err = session.Ping()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("Mongo db connected successful!")
}

//FindAll method
func (m *DAO) FindAll() ([]Record, error) {
	var records []Record
	err := db.C(COLLECTION).Find(bson.M{}).All(&records)
	return records, err
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
