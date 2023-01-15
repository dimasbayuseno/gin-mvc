package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId   int
	ItemName string
	Count    int
	Status   string
}

var PAID string = "paid"
var WAITING string = "waiting"
