package database

import (
	"context"
	"fmt"
	"image/color/palette"
	"log"
	"time"

	"github.com/MikeFors0/golang-bot/models"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection = UserData(Client, "users")

var validate = validator.New()

func gg() {
	fmt.Println("ff")
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

// Создание пользователя в бд
func AddUser(user *models.User) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

	count, err := UserCollection.CountDocuments(ctx, bson.M{"login": user.Login})
	defer cancel()
	if err != nil {
		log.Panic(err)
		fmt.Println("Error login, checking login")

	}

	if count > 0 {
		defer cancel()
		fmt.Println("user already exists")
	} else {
		validatorErr := validate.Struct(user)
		defer cancel()
		if validatorErr != nil {
			log.Panic(validatorErr.Error())
			fmt.Println(validatorErr.Error())
		}
		password := HashPassword(user.Password)
		user.Password = password
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()
		_, insertErr := UserCollection.InsertOne(ctx, user)
		if insertErr != nil {
			fmt.Sprintf("User item was not created: %v", insertErr)
		}
		defer cancel()
		fmt.Println("User created")
		return nil
	}
	return nil

}

// // Авторизация через бд
// func AuthenticateUser(login string, password string) (*models.User, error) {
// 	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	user := &models.User{}
// 	err := UserCollection.FindOne(ctx, bson.M{"login": login}).Decode(user)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			fmt.Println("user not found")
// 			return nil, fmt.Errorf("user not found")
// 		}
// 		fmt.Println("error finding user")
// 		return nil, fmt.Errorf("error finding user: %v", err)
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
// 		fmt.Println("invalid password")
// 		return nil, fmt.Errorf("invalid password")
// 	}
// 	filter := bson.M{"login": login}
// 	update := bson.M{"$set": bson.M{"logined": true}}
// 	_, err = UserCollection.UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("You are logined in system:", user.Login, user.Logined)
// 	return user, nil
// }

// // просмотр всех пользователй в бд
// func GetUsers() (*[]models.User, error) {
// 	cur, err := UserCollection.Find(context.Background(), bson.M{})
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
// 	defer cur.Close(context.Background())

// 	var users []models.User
// 	for cur.Next(context.Background()) {
// 		var user models.User
// 		err := cur.Decode(&user)
// 		if err != nil {
// 			log.Println(err)
// 			return nil, err
// 		}
// 		users = append(users, user)
// 	}
// 	if err := cur.Err(); err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
// 	return &users, nil
// }

// func SendMessage( Logined bool, message string) error {
//     if !Logined {
//         return fmt.Errorf("User not authorized")
//     }

//     _, err := UserCollection.InsertOne(nil, bson.D{{"logined", Logined}, {"message", message}})
//     if err != nil {
//         return err
//     }
// 	fmt.Println("its okey", Logined)
//     return nil
// }

func GetUser(id primitive.ObjectID) (user models.User, err error) {
	ctx := context.Background()
	filter := bson.M{"_id":id}
	err = UserCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		log.Panic(err)
	}
	return user, err
}


func AddUserTelegram(clinet *mongo.Client, tg_id int) (user models.Id_telegram, err error) {
	ctx := context.Background()
	err = UserCollection.InsertOne(ctx, tg_id)
	if err != nil {
		log.Panic(err)
	}
}

func Login(client *mongo.Client, userid primitive.ObjectID, Logined bool) error {
	ctx := context.Background()
	filter := bson.M{"_id":userid}
	_, err := UserCollection.UpdateOne(ctx, filter, bson.M{"$eq":Logined})
	return err
}


func 
