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
	"go.mongodb.org/mongo-driver/mongo/options"
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
func GetUser(fio_student string) (user models.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = UserCollection.FindOne(ctx, bson.M{"fio_student": fio_student}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("user not found")
			return user, fmt.Errorf("user not found")
		}
		fmt.Println("error finding user")
		return user, fmt.Errorf("error finding user: %v", err)
	}

	return user, nil
}

// Регистрация при /start, отправление ID_telegram в массив
func AddUserTelegram(client *mongo.Client, tg_id int) (models.Id_telegram, error) {
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
func AuthenticateUser(client *mongo.Client, tg_id int, login string, password string) (*models.User, error) {
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

func CheckNewData() {
	var lastId primitive.ObjectID // переменная для хранения последнего id в базе данных
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// поиск последнего id в базе данных
		var lastData models.Passage
		err := UserCollection.FindOne(ctx, bson.M{}, options.FindOne().SetSort(bson.M{"_id": -1})).Decode(&lastData)
		if err != nil {
			fmt.Println("error finding last data:", err)
			continue
		}
		lastId = lastData.Passage_ID

		// поиск новых данных в базе данных
		cur, err := UserCollection.Find(ctx, bson.M{"_id": bson.M{"$gt": lastId}})
		if err != nil {
			fmt.Println("error finding new data:", err)
			continue
		}

		// отправка новых данных пользователю
		for cur.Next(ctx) {
			var data models.Passage
			err := cur.Decode(&data)
			if err != nil {
				fmt.Println("error decoding data:", err)
				continue
			}

			user, err := GetUser(data.FIO_student)
			if err != nil {
				fmt.Println("error getting user:", err)
				continue
			}

			// отправка данных пользователю
			err = SendDataToUser(user.Tg_id.Id_telegram, data)
			if err != nil {
				fmt.Println("error sending data to user:", err)
				continue
			}
		}
		cur.Close(ctx)
		fmt.Println("ererer")
		time.Sleep(1 * time.Minute) // задержка между проверками базы данных
	}
}

func SendDataToUser(chatId uint, data models.Passage) error {
	// создание нового сообщения
	msg := fmt.Sprintf("New data: %s", data.Passage_ID)

	// отправка сообщения
	fmt.Println(msg)
	return nil
}

func AddData(passage models.Passage) error {
	ctx := context.Background()
	validatorErr := validate.Struct(passage)
	if validatorErr != nil {
		log.Panic(validatorErr.Error())
		fmt.Println(validatorErr.Error())
	}
	passage.Passage_ID = primitive.NewObjectID()
	passage.Passage_At, _ = time.Parse(time.ANSIC, time.Now().Format(time.ANSIC))
	_, insertErr := UserCollection.InsertOne(ctx, passage)
	if insertErr != nil {
		fmt.Sprintf("User item was not created", insertErr)
		return nil
	}
	fmt.Println("Passage created!")
	return nil
}




func GetAllPassage() ([]models.Passage, error) {
    cur, err := UserCollection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cur.Close(context.Background())

    var passages []models.Passage
    for cur.Next(context.Background()) {
        var passage models.Passage
        if err := cur.Decode(&passage); err != nil {
            return nil, err
        }
		if passage.FIO_student == "" {
			continue
		}
        passages = append(passages, passage)
    }
    return passages, nil
}