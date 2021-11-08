package db

import (
	"context"
	"log"
	"time"

	"github.com/carrenolg/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadTweets(id string, page int64) ([]*models.TweetDocument, bool) {
	// 1. set up database
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoConn.Database("twitter")
	collection := db.Collection("tweet")

	// 2. get all tweets from database
	var tweets []*models.TweetDocument
	options := options.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "datetime", Value: -1}})
	options.SetSkip((page - 1) * 20)

	filter := bson.M{
		"userid": id,
	}
	cursor, err := collection.Find(ctx, filter, options)
	if err != nil {
		log.Fatal(err.Error())
		return tweets, false
	}

	for cursor.Next(context.TODO()) {
		var tweet models.TweetDocument
		err := cursor.Decode(&tweet)
		if err != nil {
			return tweets, false
		}
		tweets = append(tweets, &tweet)
	}

	// 3. return results
	return tweets, true
}
