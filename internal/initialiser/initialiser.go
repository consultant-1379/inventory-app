package initialiser

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/config"
	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/connection"
	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/helpers"
	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMenuStruct(app *config.AppConfig) models.Menu {
	var menuItems models.Menu
	// fmt.Println("from Init")
	// fmt.Println(app.DBName)
	instances := helpers.GetAllFromCollectiion("Instances", app)
	for i := range instances {
		menuItems.Instance = append(menuItems.Instance, models.MenuInstance{Name: instances[i]["name"].(string)})
		clusters := helpers.GetByForeginKey("instance_id", instances[i]["_id"].(primitive.ObjectID), "Clusters", app)
		for c := range clusters {
			menuItems.Instance[i].Clusters = append(menuItems.Instance[i].Clusters, models.MenuCluster{Name: clusters[c]["name"].(string)})
			deployments := helpers.GetByForeginKey("cluster_id", clusters[c]["_id"].(primitive.ObjectID), "Deployments", app)
			for d := range deployments {
				menuItems.Instance[i].Clusters[c].Deployments = append(menuItems.Instance[i].Clusters[c].Deployments, models.MenuDeployment{Name: deployments[d]["name"].(string)})
			}
		}
		servers := helpers.GetByForeginKey("instance_id", instances[i]["_id"].(primitive.ObjectID), "Servers", app)
		for s := range servers {
			menuItems.Instance[i].Servers = append(menuItems.Instance[i].Servers, models.MenuServer{Name: servers[s]["name"].(string)})
			deployments := helpers.GetByForeginKey("server_id", servers[s]["_id"].(primitive.ObjectID), "Deployments", app)
			for d := range deployments {
				menuItems.Instance[i].Servers[s].Deployments = append(menuItems.Instance[i].Servers[s].Deployments, models.MenuDeployment{Name: deployments[d]["name"].(string)})
			}
		}
		vpods := helpers.GetByForeginKey("instance_id", instances[i]["_id"].(primitive.ObjectID), "Vpods", app)
		for v := range vpods {
			menuItems.Instance[i].Vpods = append(menuItems.Instance[i].Vpods, models.MenuVpod{Name: vpods[v]["name"].(string)})
			deployments := helpers.GetByForeginKey("vpod_id", vpods[v]["_id"].(primitive.ObjectID), "Deployments", app)
			for d := range deployments {
				menuItems.Instance[i].Vpods[v].Deployments = append(menuItems.Instance[i].Vpods[v].Deployments, models.MenuDeployment{Name: deployments[d]["name"].(string)})
			}
			clusters := helpers.GetByForeginKey("vpod_id", vpods[v]["_id"].(primitive.ObjectID), "Clusters", app)

			for c := range clusters {
				menuItems.Instance[i].Vpods[v].Clusters = append(menuItems.Instance[i].Vpods[v].Clusters, models.MenuCluster{Name: clusters[c]["name"].(string)})
				deployments := helpers.GetByForeginKey("cluster_id", clusters[c]["_id"].(primitive.ObjectID), "Deployments", app)
				for d := range deployments {
					menuItems.Instance[i].Vpods[v].Clusters[c].Deployments = append(menuItems.Instance[i].Vpods[v].Clusters[c].Deployments, models.MenuDeployment{Name: deployments[d]["name"].(string)})
				}
			}
		}
	}

	return menuItems
}

// FROM JSON TO STRUCT

func MenuItemsToStruct(app *config.AppConfig) models.JsonInstances {
	jsonFile, err := os.Open(app.MenuJson)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return models.JsonInstances{}
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return models.JsonInstances{}
	}
	// we initialize our Users array
	var instances models.JsonInstances

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'instances' which we defined above

	if err := json.Unmarshal(byteValue, &instances); err != nil {
		fmt.Println(err)
		return models.JsonInstances{}
	}

	return instances
}

// Package to initialize DB from JSON

func InitDB(app *config.AppConfig) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db_exist, err := connection.MongoDBClient.Database(app.DBName).ListCollectionNames(ctx, bson.M{})
	fmt.Println("Deleting DB " + app.DBName + " with collections:")
	fmt.Println(db_exist)
	helpers.CheckError(err)
	if db_exist != nil {
		connection.MongoDBClient.Database(app.DBName).Drop(ctx)
	}

	//fmt.Println("FROM INIT DB")
	instances := MenuItemsToStruct(app)
	for i := range instances.Instance {
		//fmt.Println("Instance: " + instances.Instance[i].Name)
		inst_result, err := connection.MongoDBClient.Database(app.DBName).Collection("Instances").InsertOne(ctx, bson.D{
			{Key: "name", Value: instances.Instance[i].Name},
		})
		helpers.CheckError(err)

		for c := range instances.Instance[i].Clusters {
			//fmt.Println("Cluster: " + instances.Instance[i].Clusters[c].Name)
			cluster_result, err := connection.MongoDBClient.Database(app.DBName).Collection("Clusters").InsertOne(ctx, bson.D{
				{Key: "name", Value: instances.Instance[i].Clusters[c].Name},
				{Key: "instance_id", Value: inst_result.InsertedID},
			})
			helpers.CheckError(err)

			for d := range instances.Instance[i].Clusters[c].Deployments {
				_, err := connection.MongoDBClient.Database(app.DBName).Collection("Deployments").InsertOne(ctx, bson.D{
					{Key: "name", Value: instances.Instance[i].Clusters[c].Deployments[d].Name},
					{Key: "cluster_id", Value: cluster_result.InsertedID},
				})
				helpers.CheckError(err)
				//fmt.Println("Dep: " + instances.Instance[i].Clusters[c].Deployments[d].Name)
			}
		}
		for s := range instances.Instance[i].Servers {
			//fmt.Println("Server: " + instances.Instance[i].Servers[s].Name)
			server_result, err := connection.MongoDBClient.Database(app.DBName).Collection("Servers").InsertOne(ctx, bson.D{
				{Key: "name", Value: instances.Instance[i].Servers[s].Name},
				{Key: "instance_id", Value: inst_result.InsertedID},
			})
			helpers.CheckError(err)

			for d := range instances.Instance[i].Servers[s].Deployments {
				_, err := connection.MongoDBClient.Database(app.DBName).Collection("Deployments").InsertOne(ctx, bson.D{
					{Key: "name", Value: instances.Instance[i].Servers[s].Deployments[d].Name},
					{Key: "server_id", Value: server_result.InsertedID},
				})
				helpers.CheckError(err)
				//fmt.Println("Dep: " + instances.Instance[i].Servers[s].Deployments[d].Name)
			}
		}
		for v := range instances.Instance[i].Vpods {
			//fmt.Println("Vpod: " + instances.Instance[i].Vpods[v].Name)
			vpods_result, err := connection.MongoDBClient.Database(app.DBName).Collection("Vpods").InsertOne(ctx, bson.D{
				{Key: "name", Value: instances.Instance[i].Vpods[v].Name},
				{Key: "instance_id", Value: inst_result.InsertedID},
			})
			helpers.CheckError(err)
			for d := range instances.Instance[i].Vpods[v].Deployments {
				_, err := connection.MongoDBClient.Database(app.DBName).Collection("Deployments").InsertOne(ctx, bson.D{
					{Key: "name", Value: instances.Instance[i].Vpods[v].Deployments[d].Name},
					{Key: "vpod_id", Value: vpods_result.InsertedID},
				})
				helpers.CheckError(err)
				//fmt.Println("Dep: " + instances.Instance[i].Vpods[v].Deployments[d].Name)
			}
			for c := range instances.Instance[i].Vpods[v].Clusters {
				//fmt.Println("Cluster: " + instances.Instance[i].Vpods[v].Clusters[c].Name)
				cluster_result, err := connection.MongoDBClient.Database(app.DBName).Collection("Clusters").InsertOne(ctx, bson.D{
					{Key: "name", Value: instances.Instance[i].Vpods[v].Clusters[c].Name},
					{Key: "vpod_id", Value: vpods_result.InsertedID},
				})
				helpers.CheckError(err)
				for d := range instances.Instance[i].Vpods[v].Clusters[c].Deployments {
					_, err := connection.MongoDBClient.Database(app.DBName).Collection("Deployments").InsertOne(ctx, bson.D{
						{Key: "name", Value: instances.Instance[i].Vpods[v].Clusters[c].Deployments[d].Name},
						{Key: "cluster_id", Value: cluster_result.InsertedID},
					})
					helpers.CheckError(err)
					//fmt.Println("Dep: " + instances.Instance[i].Vpods[v].Clusters[c].Deployments[d].Name)
				}
			}
		}
	}

	db_exist, err = connection.MongoDBClient.Database(app.DBName).ListCollectionNames(ctx, bson.M{})
	helpers.CheckError(err)
	fmt.Println("DB" + app.DBName + " created with collections:")
	fmt.Println(db_exist)
}
