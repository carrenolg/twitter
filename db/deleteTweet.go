package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTweet(id string, userid string) error {
	// 1. set up database
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoConn.Database("twitter")
	collection := db.Collection("tweet")

	// 2. delete document (tweet)
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id":    objId,
		"userid": userid,
	}
	_, err := collection.DeleteOne(ctx, filter)

	// 3. return state (operation)
	return err

}
