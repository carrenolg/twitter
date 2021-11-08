package db

import (
	"context"
	"time"

	"github.com/carrenolg/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertTweet(tweet models.Tweet) (string, bool, error) {
	// set up database
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoConn.Database("twitter")
	collection := db.Collection("tweet")

	// convert model (tweet) to bson
	btweet := bson.M{
		"userid":   tweet.UserId,
		"message":  tweet.Message,
		"datetime": tweet.DateTime,
	}

	// insert tweet into database
	result, err := collection.InsertOne(ctx, btweet)
	if err != nil {
		return string(""), false, err
	}

	// get Id of last tweet inserted
	id, _ := result.InsertedID.(primitive.ObjectID)

	// return
	return id.Hex(), true, nil
}
