package database

import (
	"context"
	"log"
	"time"

	"github.com/MikeFors0/golang-bot/models"
	"go.mongodb.org/mongo-driver/bson"
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


func UserData(collectionName string) *mongo.Collection {
	return Client.Database("mydb").Collection(collectionName)
}

func SubscriptionData(collectionName string) *mongo.Collection {
	return Client.Database("mydb").Collection(collectionName)
}

// AddUser adds a new user to the users collection.
func AddUser(user *models.User) error {
	collection := UserData("users")
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// GetUsers returns all users from the users collection.
func GetUsers() (*[]models.User, error) {
	collection := UserData("users")
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cur.Close(context.Background())

	var users []models.User
	for cur.Next(context.Background()) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		users = append(users, user)
	}
	if err := cur.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	return &users, nil
}
