package db

import (
	"context"
	"time"

	"github.com/carrenolg/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ReadTweetsFollowers(id string, page int) ([]models.TweetFollower, bool) {
	// 1. Set up database
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*15)
	defer cancel()
	db := MongoConn.Database("twitter")
	collection := db.Collection("relationship")

	// 2. Set up aggregation
	skip := (page - 1) * 20
	stages := make([]bson.M, 0)

	stages = append(stages, bson.M{
		"$match": bson.M{"userid": id},
	})

	stages = append(stages, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "useridrelationed",
			"foreignField": "userid",
			"as":           "tweet",
		},
	})

	stages = append(stages, bson.M{
		"$unwind": "$tweet",
	})

	stages = append(stages, bson.M{
		"$sort": bson.M{"tweet.datetime": -1},
	})

	stages = append(stages, bson.M{"$skip": skip})
	stages = append(stages, bson.M{"$limit": 20})

	cursor, err := collection.Aggregate(ctx, stages)
	var tweets []models.TweetFollower
	err = cursor.All(ctx, &tweets)
	if err != nil {
		return tweets, false
	}

	// 3. return state operation
	return tweets, true
}
