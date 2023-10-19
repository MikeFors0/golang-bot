package database

import (
	"context"
	"fmt"

	// "go/parser"
	"log"
	"time"

	"github.com/MikeFors0/golang-bot/models"
	"github.com/MikeFors0/golang-bot/pkg/telegram"
	"github.com/MikeFors0/golang-bot/telegram"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection = UserData(Client, "users")
var PassageCollection *mongo.Collection = PassageData(Client, "passages")
var SubscriptionCollection *mongo.Collection = Subscription(Client, "subscription")
var validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

// Создание пользователя в бд
func AddUser(*models.User) error {
	user := models.User{}
	ctx := context.Background()

	count, err := UserCollection.CountDocuments(ctx, bson.M{"login": user.Login})
	if err != nil {
		log.Println("ошибка логина, проверьте введенные данные")
		return err
	}
	if count > 0 {
		log.Println("пользователь уже есть в системе")
		return err
	}
	validatorErr := validate.Struct(user)
	if validatorErr != nil {
		log.Println(validatorErr.Error())
		return validatorErr
	}
	user.ID = primitive.NewObjectID()
	user.User_ID = user.ID.Hex()
	password := HashPassword(user.Password)
	user.Password = password
	_, insertErr := UserCollection.InsertOne(ctx, user)
	if insertErr != nil {
		log.Println(insertErr)
		return insertErr
	}
	log.Println("пользователь успешно создан:", user.Login)
	return nil
}

// просмотр всех пользователй в бд
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
	for _, user := range users {
		log.Println(user.FIO_student, user.Login, user.Logined, user.Passage_student)
	}
	return &users, nil
}

// Доступ к пользователя по ID_telegram
func GetUser(tg_id int64) (*models.User, error) {
	user := &models.User{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var id_telegram models.Id_telegram
	id_telegram.Id_telegram = tg_id

	err := UserCollection.FindOne(ctx, bson.M{"tg_id.id_telegram": tg_id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("пользователь не найден", err)
			return nil, err
		}
		fmt.Println("ошибка при поиске пользователя", err)
		return nil, err
	}
	log.Println(user.FIO_student, user.Login, user.Passage_student)
	return user, nil
}

// Доступ к пользователя по fio
func GetUserByFIO(fio_student string) (*models.User, error) {
	user := &models.User{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := UserCollection.FindOne(ctx, bson.M{"fio_student": fio_student}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("пользователь не найден")
			return nil, err
		}
		fmt.Println("ошибка при поиске пользователя")
		return nil, err
	}
	// log.Println(user.FIO_student, user.Login, user.Passage_student)
	return user, nil
}

// Регистрация при /start, отправление ID_telegram в массив
func AddUserTelegram(client *mongo.Client, tg_id int64) (*models.Id_telegram, error) {
	ctx := context.Background()
	var id_telegram models.Id_telegram
	id_telegram.Id_telegram = tg_id
	_, err := UserCollection.InsertOne(ctx, id_telegram)
	if err != nil {
		log.Println("error ", err)
		return nil, err
	}
	log.Println(id_telegram)
	return &id_telegram, nil
}

// авторизация через бд - лк САМГК (с добавлением ID_telegram)
func AuthenticateUser(client *mongo.Client, tg_id int64, login string, password string) (*models.User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := &models.User{}

	err := UserCollection.FindOne(ctx, bson.M{"login": login}).Decode(&user)
	if err != nil {
		fmt.Println("ошибка при поиске пользователя")
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		fmt.Println("Неверный пароль")
		return nil, err
	}

	user.Login = login
	var id_telegram models.Id_telegram
	id_telegram.Id_telegram = tg_id
	user.Tg_id = id_telegram
	user.Logined = true

	filter := bson.M{"login": login}

	var _, errUpdateUserCollection = UserCollection.UpdateOne(ctx, filter, bson.M{"$set": user})
	if errUpdateUserCollection != nil {
		log.Println("фатальное обновление с mongoDB: ", errUpdateUserCollection)
		return nil, err
	}
	fmt.Println("Вы вошли в систему:", user.Login, user.Logined, user.Tg_id)

	return user, nil
}

// Добавление passage
func AddPassage(passage models.Passage) error {
	ctx := context.Background()

	if err := validate.Struct(passage); err != nil {
		log.Println("Ошибка проверки:", err)
		return nil
	}

	passage.Passage_ID = primitive.NewObjectID()
	passage.Passage_At = time.Now()

	if _, err := PassageCollection.InsertOne(ctx, passage); err != nil {
		log.Println("Ошибка вставки:", err)
		return err
	}
	log.Println("Passage создан:", passage.FIO_student)

	var user, err = GetUserByFIO(passage.FIO_student)
	if err != nil {
		log.Println(err)
		return nil
	}

	filter := bson.M{"_id": user.ID}
	if !user.Logined {
		log.Println("пользователь не авторизирован")
		return nil
	}
	user.Passage_student = append(user.Passage_student, passage)
	if _, err := UserCollection.UpdateOne(ctx, filter, bson.M{"$set": user}); err != nil {
		log.Println("ошибка вставки в пользователя:", err)
		return err
	}
	log.Println("Passage отправлен:", passage.FIO_student)
	return nil
}

// Просмотр всех passages
func GetAllPassages() ([]models.Passage, error) {
	ctx := context.Background()

	cur, err := PassageCollection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("Ошибка базы данных", err)
		return nil, err
	}
	defer cur.Close(ctx)

	var passages []models.Passage
	for cur.Next(ctx) {
		var passage models.Passage
		if err := cur.Decode(&passage); err != nil {
			fmt.Println("ошибка декодирования:", err)
			return nil, err
		}
		if passage.FIO_student == "" {
			continue
		}
		passages = append(passages, passage)
	}
	for _, passage := range passages {
		log.Println(passage.FIO_student)
		log.Println(passage.Passage_At)
	}
	return passages, nil
}

func CheckNewData() error {
	var lastId primitive.ObjectID // переменная для хранения последнего id в базе данных
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// поиск последнего id в базе данных
		var lastData models.Passage
		err := PassageCollection.FindOne(ctx, bson.M{}, options.FindOne().SetSort(bson.M{"_id": -1})).Decode(&lastData)
		if err != nil {
			log.Println("ошибка при поиске последних данных:", err)
			continue
		}
		lastId = lastData.Passage_ID
		log.Println(lastData)
		pass := models.Passage{
			FIO_student: "gnome",
		}
		AddPassage(pass)

		// поиск новых данных в базе данных
		cur, err := PassageCollection.Find(ctx, bson.M{"_id": bson.M{"$gt": lastId}})
		if err != nil {
			log.Println("ошибка при поиске новых данных:", err)
			continue
		}
		defer cur.Close(ctx)

		for cur.Next(ctx) {
			var passage models.Passage
			if err := cur.Decode(&passage); err != nil {
				log.Println("hh", err)
				continue
			}
			var user models.User
			if err := UserCollection.FindOne(ctx, bson.M{"fio_student": passage.FIO_student}).Decode(&user); err != nil {
				log.Println("обнаружена ошибка пользователя", err)
				continue
			}
			if !user.Logined {
				log.Print("Пользователь не может получить данные в массив, так как не авторизирован")
				continue
			}

			err = SendDataToUser(telegram.Bot,user.Tg_id.Id_telegram, passage)
			if err != nil {
				log.Println("Ошибка отправки данных пользователю")
			}
			if err := cur.Err(); err != nil {
				log.Println("yy", err)
			}
			cur.Close(ctx)
			time.Sleep(5 * time.Second)
		}
		time.Sleep(5 * time.Second)
	}
}

func AddSubscription(*models.Subscription) error {
	subscription := models.Subscription{}
	ctx := context.Background()

	count, err := SubscriptionCollection.CountDocuments(ctx, bson.M{"subscription_name": subscription.Subscription_Name})
	if err != nil {
		log.Println("ошибка имени продукта", err)
		return err
	}

	if count > 0 {
		log.Println("продукт уже существует", err)
		return err
	}
	validatorErr := validate.Struct(subscription)
	if validatorErr != nil {
		log.Println(validatorErr.Error())
		return validatorErr
	}
	subscription.Subscription_ID = primitive.NewObjectID()

	_, insertErr := SubscriptionCollection.InsertOne(ctx, subscription)
	if insertErr != nil {
		log.Println("Ошибка создания продукта", err)
		return insertErr
	}
	log.Println("Продукт создан:", subscription.Subscription_Name)
	return nil
}
