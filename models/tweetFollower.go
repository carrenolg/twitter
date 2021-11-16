package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TweetFollower struct {
	Id               primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId           string             `bson:"userid" json:"userId,omitempty"`
	UseridRelationed string             `bson:"useridrelationed" json:"useridRelationed,omitempty"`
	Tweet            struct {
		Id       string    `bson:"_id" json:"_id,omitempty"`
		Message  string    `bson:"message" json:"message,omitempty"`
		DateTime time.Time `bson:"datetime" json:"datetime,omitempty"`
	}
}
