package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TweetDocument struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId   string             `bson:"userid" json:"userid,omitempty"`
	Message  string             `bson:"message" json:"message,omitempty"`
	DateTime time.Time          `bson:"datetime" json:"datetime,omitempty"`
}
