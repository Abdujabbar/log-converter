package repository

import "gopkg.in/mgo.v2/bson"

//Record Log record
type Record struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Time   int64         `bson:"log_time" json:"log_time"`
	Msg    string        `bson:"msg" json:"msg"`
	Format string        `bson:"format" json:"format"`
}
