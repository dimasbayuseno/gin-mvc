package main

import (
	"gin-mvc/controller"
	"gin-mvc/db"
	"gin-mvc/middleware"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	r := gin.Default()

	db.SetupDatabaseConnection()

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "gin-mvc",
		Key:             []byte("secret key"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     middleware.IdentityKey,
		PayloadFunc:     middleware.PayloadFunc,
		IdentityHandler: middleware.IdentityHandler,
		Authenticator:   middleware.Authenticator,
		Authorizator:    middleware.Authorizator,
		Unauthorized:    middleware.Unauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		SendCookie:      true,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	r.Use(CORSMiddleware())

	r.GET("/hello-world", controller.HelloWorld)
	r.POST("/register", controller.Register)
	r.POST("/login", authMiddleware.LoginHandler)
	r.GET("/logout", authMiddleware.LogoutHandler)

	auth := r.Group("")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/users", controller.GetAllUsers)
		auth.GET("/users/search", controller.SearchUsers)
		auth.GET("/user", controller.GetUserByID)
		auth.PATCH("/user/patch", controller.UpdateUserByID)
		auth.DELETE("/user/delete", controller.DeleteUserByID)
		auth.POST("/user", controller.PostUser)

		auth.GET("/orders", controller.GetAllOrders)
		auth.GET("/orders/search", controller.SearchOrders)
		auth.GET("/order", controller.GetOrderByID)
		auth.PATCH("/order/patch", controller.UpdateOrderByID)
		auth.DELETE("/order/delete", controller.DeleteOrderByID)
		auth.POST("/order", controller.PostOrder)
	}

	log.Fatal(r.Run(":" + port))
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Whitelist FE localhost
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
