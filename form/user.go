package form

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"Username" binding:"required"`
	Password string `json:"Password"`
	Fullname string `json:"Fullname"`
	Role     string `json:"Role"`
}
