package db

import (
	"context"
	"log"
	"time"

	"github.com/carrenolg/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchProfile(id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoConn.Database("twitter")
	collection := db.Collection("users")

	var profile models.User
	objId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": objId}

	err := collection.FindOne(ctx, filter).Decode(&profile)
	// set password to emtpy
	profile.Password = string("")

	if err != nil {
		log.Printf("registry not find - %s", err.Error())
		return profile, err
	}

	return profile, nil
}
