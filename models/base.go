package models

import (
	"labix.org/v2/mgo/bson"
)

// Base entity for models
type Base struct {
	ID bson.ObjectId `bson: "_id"`
}
