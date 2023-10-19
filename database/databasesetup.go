package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err, "ffffff")
	}

	log.Println("> Connected to MongoDB!")

	return client
}

var Client *mongo.Client = Init()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("okoko").Collection(collectionName)
	return collection
}

func PassageData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("okoko").Collection(collectionName)
	return collection
}

func Subscription(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("okoko").Collection(collectionName)
	return collection
}


func SubscriptionData(collectionName string) *mongo.Collection {
	return Client.Database("test3").Collection(collectionName)
}
