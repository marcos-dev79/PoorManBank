package bankuser

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Since    int64     `json:"since"`
	Birthday int64     `json:"birthday"`
	Address  string    `json:"address"`
	City     string    `json:"city"`
	Country  string    `json:"country"`
}
