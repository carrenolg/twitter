package db

import (
	"context"
	"time"

	"github.com/carrenolg/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckRelation(relation models.Relation) (bool, error) {
	// 1. Set up database
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*15)
	defer cancel()
	db := MongoConn.Database("twitter")
	collection := db.Collection("relationship")

	// 2. Get relation from db
	filter := bson.M{
		"userid":           relation.Userid,
		"useridrelationed": relation.UseridRelationed,
	}
	var mrelation models.Relation
	err := collection.FindOne(ctx, filter).Decode(&mrelation)
	if err != nil {
		return false, err
	}

	// 3. Return state operation
	return true, err
}
