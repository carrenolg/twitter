package db

import (
	"context"
	"log"
	"time"

	"github.com/carrenolg/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadAllUsers(id string, page int64, search string, typeUser string) ([]*models.User, bool) {
	// 1. Set up database
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*15)
	defer cancel()
	db := MongoConn.Database("twitter")
	collection := db.Collection("users")

	// 2. Search users (keep in mind order about setskip, setlimit )
	var users []*models.User
	options := options.Find()
	options.SetSkip((page - 1) * 20)
	options.SetLimit(20)
	filter := bson.M{
		"fristName": bson.M{
			"$regex": `(?i)` + search,
		},
	}
	cursor, err := collection.Find(ctx, filter, options)
	if err != nil {
		log.Println(err)
		return users, false
	}

	// 3. Check relation by each user found
	for cursor.Next(ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Println(err)
			return users, false
		}

		relation := models.Relation{
			Userid:           id,
			UseridRelationed: user.ID.Hex(),
		}

		var isRelationed bool = false
		var add bool = false
		// Note: error is not check in this step
		// only is needed knowing if there's relation
		isRelationed, _ = CheckRelation(relation)
		/*if err != nil {
			log.Println(err)
			return users, false
		}*/

		if isRelationed == false && typeUser == "new" {
			add = true
		}
		if isRelationed == true && typeUser == "follow" {
			add = true
		}
		if id == relation.UseridRelationed {
			add = false
		}

		if add == true {
			user.Password = string("")
			user.Biography = string("")
			user.WebSite = string("")
			user.Location = string("")
			user.Banner = string("")
			user.Email = string("")

			users = append(users, &user)
		}
	}
	err = cursor.Err()
	if err != nil {
		log.Println(err)
		return users, false
	}
	cursor.Close(ctx)
	return users, true
}
