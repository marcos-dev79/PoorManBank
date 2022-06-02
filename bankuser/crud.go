package bankuser

import (
	"context"
	"net/http"
	"poormanbank/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

func Userlist(c *gin.Context) {
	mongoclient := database.Mongo_connect()
	usersCollection := mongoclient.Database("poormanbank").Collection("users")
	results, err := usersCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": "Error in listing"})
	}
	c.JSON(http.StatusOK, results)
}

func Insertion(c *gin.Context) {
	mongoclient := database.Mongo_connect()
	var userRequest User
	if err := c.BindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
	}

	uuidn, err := uuid.New()
	userRequest.Id = uuidn

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating UUID."})
	}

	usersCollection := mongoclient.Database("poormanbank").Collection("users")
	result, err := usersCollection.InsertOne(context.TODO(), userRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when saving."})
	}

	c.JSON(http.StatusOK, result)

}
