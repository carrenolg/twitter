package models

type Relation struct {
	Userid           string `bson:"userid" json:"userid"`
	UseridRelationed string `bson:"useridrelationed" json:"useridRelationed"`
}
