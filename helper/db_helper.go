package helper

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
	"youtube/constants"
)

var Client *mongo.Client

func GetClient() {
	var err error
	// Set client options and create new client
	mongoDbUrl := os.Getenv("MONGO_DB_URL")
	if len(mongoDbUrl) == 0 {
		mongoDbUrl = constants.DefaultMongoURI
	}
	Client, err = mongo.NewClient(options.Client().ApplyURI(mongoDbUrl))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Client created successfully")
	// Create a context which will be used to set a timeout for the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
