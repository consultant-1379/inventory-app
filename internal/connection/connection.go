package connection

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBClient is exported Mongo Database client
var MongoDBClient *mongo.Client

// SOME more connection settings

func ConnectDatabase(dsn string) {
	log.Println("Database connecting...")
	// Set client options
	clientOptions := options.Client().ApplyURI(dsn)
	clientOptions.SetConnectTimeout(1 * time.Second)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	MongoDBClient = client
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = MongoDBClient.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database Connected.")
}
