package db

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// exported variable
var StringConn = GetStringConn()
var clientOptions = options.Client().ApplyURI(StringConn)
var MongoConn = ConnTodb()

func ConnTodb() *mongo.Client {
	// get client connection
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}
	// check connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("mongo: successful connection")

	return client
}

func CheckConnection() int {
	err := MongoConn.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}

func GetStringConn() string {
	var envs map[string]string
	envs, err := godotenv.Read("db/.env")
	if err != nil {
		log.Printf("err: %s", err)
		log.Fatal("Error loading .env file")
	}

	stringConn := envs["STRCONNECTION"]
	return stringConn
}
