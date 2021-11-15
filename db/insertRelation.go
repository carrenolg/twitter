package db

import (
	"context"
	"time"

	"github.com/carrenolg/twitter/models"
)

func InsertRelation(relation models.Relation) (bool, error) {
	// 1. set up database
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoConn.Database("twitter")
	collection := db.Collection("relationship")

	// 2. Insert model
	_, err := collection.InsertOne(ctx, relation)
	if err != nil {
		return false, err
	}

	// 3. Return state operation
	return true, err
}
