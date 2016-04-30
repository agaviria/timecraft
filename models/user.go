package models

import (
	"time"
)

// User user to authenticate
type User struct {
	Base        []byte `storm:"inline"`
	Username    string `storm:"unique"`
	Email       []byte
	Password    []byte
	DateCreated time.Time
	LastLogin   time.Time
}
