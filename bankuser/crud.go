package bankuser

import (
	"context"
	"net/http"
	"poormanbank/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Userlist(c *gin.Context) {
	mongoclient := database.Mongo_connect()
	usersCollection := mongoclient.Database("poormanbank").Collection("users")
	cursor, err := usersCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		panic(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, results)
}

func Insertion(c *gin.Context) {
	mongoclient := database.Mongo_connect()
	var userRequest User
	if err := c.BindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		panic(err)
	}

	usersCollection := mongoclient.Database("poormanbank").Collection("users")
	result, err := usersCollection.InsertOne(context.TODO(), userRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		panic(err)
	}

	c.JSON(http.StatusOK, result)

}
