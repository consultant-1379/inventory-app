package helpers

import (
	"context"
	"fmt"
	"log"

	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/config"
	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/connection"
	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Mongo DB helpers

func InsertOneDeployment(deployment models.Deployment, app *config.AppConfig) {
	inserted, err := connection.MongoDBClient.Database(app.DBName).Collection("Deployments").InsertOne(context.Background(), deployment)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Inserted 1 deployment with id:", inserted.InsertedID)

}

func UpdateOneDeployment(deploymentID string, app *config.AppConfig) {

	id, _ := primitive.ObjectIDFromHex(deploymentID)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"inUse": true}}

	updated, err := connection.MongoDBClient.Database(app.DBName).Collection("Deployments").UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count:", updated.ModifiedCount)
}

func DeletOneDeployment(deploymentID string, app *config.AppConfig) {
	id, _ := primitive.ObjectIDFromHex(deploymentID)

	filter := bson.M{"_id": id}

	deleted, err := connection.MongoDBClient.Database(app.DBName).Collection("Deployments").DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count: ", deleted.DeletedCount)
}

func GetOneDeployment(deploymentID string, app *config.AppConfig) (models.Deployment, error) {
	id, _ := primitive.ObjectIDFromHex(deploymentID)
	var result models.Deployment
	err := connection.MongoDBClient.Database(app.DBName).Collection("Deployments").FindOne(context.Background(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return models.Deployment{}, err
		}
		panic(err)
	}

	fmt.Println("found", result)
	return result, nil
}

func DeleteAllDeployments(app *config.AppConfig) {
	deleted, err := connection.MongoDBClient.Database(app.DBName).Collection("Deployments").DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count: ", deleted.DeletedCount)
}

func GetAllFromCollectiion(colName string, app *config.AppConfig) []primitive.M {
	cursor, err := connection.MongoDBClient.Database(app.DBName).Collection(colName).Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var results []primitive.M

	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	defer cursor.Close(context.Background())
	return results
}

func GetByForeginKey(foreginKey string, foreginId primitive.ObjectID, colName string, app *config.AppConfig) []primitive.M {
	cursor, err := connection.MongoDBClient.Database(app.DBName).Collection(colName).Find(context.Background(), bson.M{foreginKey: foreginId})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var results []primitive.M

	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	defer cursor.Close(context.Background())
	return results
}

func GetOneByName(colName string, name string, app *config.AppConfig) []primitive.M {
	cursor, err := connection.MongoDBClient.Database(app.DBName).Collection(colName).Find(context.Background(), bson.M{"name": name})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var results []primitive.M

	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	defer cursor.Close(context.Background())
	return results
}
