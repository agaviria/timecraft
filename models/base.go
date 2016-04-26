package models

import (
	"gopkg.in/mgo.V2/bson"
)

// Base entity for models
type Base struct {
	ID bson.ObjectID `storm:"id"`
}
