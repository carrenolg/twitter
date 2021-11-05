package db

import (
	"context"
	"time"

	"github.com/carrenolg/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModifyProfile(user models.User, id string) (bool, error) {
	// set up database
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	database := MongoConn.Database("twitter")
	collection := database.Collection("users")

	// validate user model
	registry := make(map[string]interface{})
	registry["dateBirth"] = user.DateBirth
	if len(user.FirstName) > 0 {
		registry["fristName"] = user.FirstName
	}
	if len(user.LastName) > 0 {
		registry["lastName"] = user.LastName
	}
	if len(user.Avatar) > 0 {
		registry["avatar"] = user.Avatar
	}
	if len(user.Banner) > 0 {
		registry["banner"] = user.Banner
	}
	if len(user.Biography) > 0 {
		registry["biography"] = user.Biography
	}
	if len(user.Location) > 0 {
		registry["location"] = user.Location
	}
	if len(user.WebSite) > 0 {
		registry["website"] = user.WebSite
	}
	userModel := bson.M{
		"$set": registry,
	}

	// update database
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": bson.M{"$eq": objId}}
	_, err := collection.UpdateOne(ctx, filter, userModel)
	if err != nil {
		return false, err
	}
	return true, nil
}
