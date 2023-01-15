package controller

import (
	"gin-mvc/db"
	"gin-mvc/form"
	"gin-mvc/model"
	"gin-mvc/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input form.User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Username: input.Username,
		Password: input.Password,
		Fullname: input.Fullname,
	}
	db.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func GetAllUsers(c *gin.Context) {
	var users []model.User
	pagination := utils.GeneratePaginationFromRequest(c)
	offset := (pagination.Page - 1) * pagination.Limit
	db.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Find(&users)

	c.JSON(http.StatusOK, users)
}

func SearchUsers(c *gin.Context) {
	var users []model.User
	searchQuery := "%" + c.Query("search") + "%"
	db.DB.Where("fullname LIKE ?", searchQuery).Find(&users)
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	var user model.User
	if err := db.DB.Where("id = ?", c.Query("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func PostUser(c *gin.Context) {
	var input form.User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product := model.User{
		Username: input.Username,
		Password: input.Password,
		Fullname: input.Fullname,
	}
	db.DB.Create(&product)

	c.JSON(http.StatusCreated, product)
}

func UpdateUserByID(c *gin.Context) {
	var product model.User
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input form.User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&product).Updates(input)

	c.JSON(http.StatusOK, product)
}

func DeleteUserByID(c *gin.Context) {
	var product model.User
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
