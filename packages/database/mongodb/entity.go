package mongodb

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Entity struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	InsertedAt time.Time     `json:"inserted_at" bson:"inserted_at"`
	LastUpdate time.Time     `json:"last_update" bson:"last_update"`
}
