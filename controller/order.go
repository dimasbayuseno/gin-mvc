package controller

import (
	"gin-mvc/db"
	"gin-mvc/model"
	"gin-mvc/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllOrders(c *gin.Context) {
	var orders []model.Order
	pagination := utils.GeneratePaginationFromRequest(c)
	offset := (pagination.Page - 1) * pagination.Limit
	db.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Find(&orders)

	c.JSON(http.StatusOK, orders)
}

func SearchOrders(c *gin.Context) {
	var orders []model.Order
	searchQuery := "%" + c.Query("search") + "%"
	db.DB.Find(&orders).Where("item_name LIKE ?", searchQuery)

	c.JSON(http.StatusOK, orders)
}

func GetOrderByID(c *gin.Context) {
	var input model.Order
	if err := db.DB.Where("id = ?", c.Query("id")).First(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, input)
}

func PostOrder(c *gin.Context) {
	var input model.Order
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product := model.Order{
		UserId:   input.UserId,
		ItemName: input.ItemName,
		Count:    input.Count,
		Status:   model.WAITING,
	}
	db.DB.Create(&product)

	c.JSON(http.StatusCreated, product)
}

func UpdateOrderByID(c *gin.Context) {
	var product model.Order
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input model.Order
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&product).Updates(input)

	c.JSON(http.StatusOK, product)
}

func DeleteOrderByID(c *gin.Context) {
	var product model.Order
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
