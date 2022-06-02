package main

import (
	"net/http"
	"poormanbank/bankuser"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	router.GET("/user/list", bankuser.Userlist)
	router.POST("/user/register", bankuser.Insertion)

	router.Run(":8040")
}
