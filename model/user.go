package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"-"`
	Password string `json:"-"`
	Fullname string
}

type UserModel struct {
	gorm.Model
	Username string
	Fullname string
}
