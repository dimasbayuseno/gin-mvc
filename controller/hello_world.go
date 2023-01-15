package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloWorld(c *gin.Context) {

	c.JSON(http.StatusOK, "Hello World")
}
