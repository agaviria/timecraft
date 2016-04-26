package models

// User user to authenticate
type User struct {
	Base     `storm:"inline"`
	Username string `storm:"unique"`
	Email    string
	Password string
}
