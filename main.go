package main

import (
	"net/http"
	"poormanbank/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	mongoclient := database.Mongo_connect()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	router.POST("/createaccount", func(c *gin.Context) {
		mongoclient.Database("poormanbank").Collection("users")
	})

	router.Run(":8060")
}
