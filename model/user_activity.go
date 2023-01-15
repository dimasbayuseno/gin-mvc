package model

import "gorm.io/gorm"

type UserActivity struct {
	gorm.Model
	Username string
	Activity string
}

var LOGIN string = "login"
var LOGOUT string = "logout"
var INTRUSION string = "intrusion"
