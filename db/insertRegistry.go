package db

import (
	"context"
	"time"

	"github.com/carrenolg/twitter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertRegistry(user models.User) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConn.Database("twitter")
	collection := database.Collection("users")

	user.Password, _ = EncryptPassword(user.Password)

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}

	objectId, _ := result.InsertedID.(primitive.ObjectID)
	return objectId.String(), true, nil
}
