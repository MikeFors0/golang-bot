package database

import (
	"context"
	"errors"
	"fmt"

	"log"
	"time"

	"github.com/MikeFors0/golang-bot/pkg/models"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection = UserData(Client, "users")
var PassageCollection *mongo.Collection = PassageData(Client, "passages")

// var SubscriptionCollection *mongo.Collection = Subscription(Client, "subscription")
var validate = validator.New()

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
		fmt.Println("User created", user.Login)
		return nil
	}
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
		log.Println(user.FIO_student, user.Login, user.Logined)
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
	log.Println(user.FIO_student, user.Login)
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
func AddUserTelegram(tg_id int64) (*models.Id_telegram, error) {
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
func AuthenticateUser(tg_id int64, login string, password string) (*models.User, error) {
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

	// var user, err = GetUserByFIO(passage.FIO_student)
	// if err != nil {
	// 	log.Println(err)
	// 	return nil
	// }

	// filter := bson.M{"_id": user.ID}
	// if !user.Logined {
	// 	log.Println("пользователь не авторизирован")
	// 	return nil
	// }
	// user.Passage_student = append(user.Passage_student, passage)
	// if _, err := UserCollection.UpdateOne(ctx, filter, bson.M{"$set": user}); err != nil {
	// 	log.Println("ошибка вставки в пользователя:", err)
	// 	return err
	// }
	// log.Println("Passage отправлен:", passage.FIO_student)
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

func SearchItemInDB() (*models.User, *models.Passage, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    options := options.FindOne().SetSort(bson.D{{"_id", -1}})
    var passage models.Passage
    err := PassageCollection.FindOne(ctx, bson.D{}, options).Decode(&passage)
    if err != nil {
        return nil, nil, err
    }
	if passage.Flag {
		log.Println("Passage уже отправлен")
		return nil, nil, err
	}

	var user models.User


	if err := UserCollection.FindOne(ctx, bson.M{"fio_student":passage.FIO_student}).Decode(&user); err != nil {
		log.Println("пользователь не найден, отрпавка отменена")
	}
	if !user.Logined {
		log.Println("пользователь не авторизирован в системе, отправка отменена")
		return nil, nil, err
	}
	passage.Flag = true
	update := bson.D{{"$set", bson.D{{"flag", true}}}}
    _, err = PassageCollection.UpdateOne(ctx, bson.M{"_id": passage.Passage_ID}, update)
	
	log.Println(&user, &passage)
    return &user, &passage, nil
}



func AddSubscription(sb *models.Subscription) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

	count, err := UserCollection.CountDocuments(ctx, bson.M{"name": sb.Name})
	defer cancel()
	if err != nil {
		log.Println("ошибка имени продукта", err)
		return err

	}
	if count > 0 {
		defer cancel()
		log.Println("продукт уже существует", err)
		return err
	} else {
		validatorErr := validate.Struct(sb)
		defer cancel()
		if validatorErr != nil {
			log.Println(validatorErr.Error())
			return validatorErr
		}
		sb.ID = primitive.NewObjectID()
		_, insertErr := UserCollection.InsertOne(ctx, sb)
		if insertErr != nil {
			log.Println("Ошибка создания продукта", err)
			return insertErr
		}
		defer cancel()
		log.Println("Продукт создан:", sb)
		return nil
	}
}

func BuySubscription(tg_id int64, subscriptionID primitive.ObjectID) error {
	user, err := GetUser(tg_id)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(user)

	var subscription models.Subscription
	err = UserCollection.FindOne(context.Background(), bson.M{"_id": subscriptionID}).Decode(&subscription)
	if err != nil {
		log.Println("подписка не найдена")
		return err
	}
	var id_telegram models.Id_telegram
	id_telegram.Id_telegram = tg_id
	subscriptionUser := models.SubscriptionUser{
		Tg_id:          id_telegram,
		SubscriptionID: subscriptionID,
		StartDate:      time.Now(),
		EndDate:        time.Now().AddDate(0, 1, 0), // Set end date to 1 month from now
		IsActive:       true,
	}

	res, err := UserCollection.InsertOne(context.Background(), subscriptionUser)
	if err != nil {
		log.Println("Ошибка добавления подписки пользователя")
		return err
	}
	log.Println(res)

	subscriptionUserID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Println("Ошибка перехода по ID подписки пользователя")
		return errors.New("Ошибка перехода по ID подписки пользователя")
	}

	order := models.Order{
		SubscriptionUserID: subscriptionUserID,
		Date:               time.Now(),
	}

	res, err = UserCollection.InsertOne(context.Background(), order)
	if err != nil {
		log.Println("Ошибка добавления заказа")
		return err
	}
	log.Println(res)

	orderID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Println("Ошибка перехода по ID заказа")
		return errors.New("Ошибка перехода по ID заказа")
	}

	payment := models.Payment{
		OrderID:    orderID,
		Amount:     subscription.Price,
		PaymentURL: "https://example.com/payment", // Replace with actual payment URL
	}

	_, err = UserCollection.InsertOne(context.Background(), payment)
	if err != nil {
		log.Println("Ошибка добавления документа оплаты")
		return err
	}
	log.Println(payment, order.SubscriptionUserID)
	return nil
}

func CheckSubscription(tg_id int64) (bool, error) {
	user, err := GetUser(tg_id)
	if err != nil {
		log.Println(err)
		return false, err
	}

	var subscriptionUser models.SubscriptionUser
	err = UserCollection.FindOne(context.Background(), bson.M{"id_tg": user.Tg_id}).Decode(&subscriptionUser)
	if err != nil {
		log.Println("Подписка не найдена")
		return false, err
	}
	log.Println(subscriptionUser)

	if !subscriptionUser.IsActive || time.Now().After(subscriptionUser.EndDate) {
		log.Println("Подписка истекла")
		return false, nil
	}
	log.Println("подиска +")
	return true, nil
}
