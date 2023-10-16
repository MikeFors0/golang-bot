package database

import (
	"context"
	"fmt"

	// "go/parser"
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


// // просмотр всех пользователй в бд
func GetUsers() (*[]models.User, error) {
	cur, err := UserCollection.Find(context.Background(), bson.M{})
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


// Доступ к пользователя по ID_telegram
func GetUser(tg_id string) (user models.User, err error) {
	ctx := context.Background()
	filter := bson.M{"tg_id": tg_id}
	err = UserCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		log.Panic(err)
	}
	return user, err
}


// Регистрация при /start, отправление ID_telegram в массив
func AddUserTelegram(client *mongo.Client, tg_id string) (models.Id_telegram, error) {
	ctx := context.Background()
	var id_telegram models.Id_telegram
	id_telegram.Id_telegram = tg_id
	_, err := UserCollection.InsertOne(ctx, id_telegram)
	if err != nil {
		log.Panic("error ", err)
	}
	fmt.Println(id_telegram)
	return id_telegram, err
}


// авторизация через бд - лк САМГК (с добавлением ID_telegram)
func AuthenticateUser(client *mongo.Client, tg_id string, login string, password string) (*models.User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := &models.User{}

	err := UserCollection.FindOne(ctx, bson.M{"login": login}).Decode(&user)
	if err != nil {
		fmt.Println("ere error finding user")
		return nil, fmt.Errorf("error finding user: %v", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		fmt.Println("invalid password")
		return nil, fmt.Errorf("invalid password")
	}

	user.Login = login
	var id_telegram models.Id_telegram
	id_telegram.Id_telegram = tg_id
	user.Tg_id = id_telegram
	user.Logined = true

	filter := bson.M{"login": login}

	var _, errUpdateUserCollection = UserCollection.UpdateOne(ctx, filter, bson.M{"$set": user})
	if errUpdateUserCollection != nil {
		log.Fatalf("fatal update with mongoDB: ", errUpdateUserCollection)
	}
	fmt.Println("You are logined in system:", user.Login, user.Logined, user.Tg_id)

	return user, nil
}

