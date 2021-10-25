package db

import (
	"context"
	"time"

	"github.com/carrenolg/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckUserExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConn.Database("twitter")
	collection := database.Collection("users")

	filter := bson.M{"email": email}
	var user models.User

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, false, user.ID.Hex()
	}
	return user, true, user.ID.Hex()
}
