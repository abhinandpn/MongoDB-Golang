package repo

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/abhinandpn/MongoDB-Golang/model"
	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoClint() *mongo.Client {

	mongoTestClint, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb+srv://abhinandpn:6AJfU9xz2xQY1bLb@testecom.jioh5.mongodb.net/?retryWrites=true&w=majority&appName=TestEcom"))
	if err != nil {
		log.Fatal("Error connecting MongoDB client:", err)
	}

	log.Println("mongo client connected successfully")

	err = mongoTestClint.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed", err)
	}

	log.Println("ping success")

	return mongoTestClint
}

func TestMongoOpration(t *testing.T) {

	mongotestClint := NewMongoClint()
	defer mongotestClint.Disconnect(context.Background())

	Usr1 := uuid.New().String()
	// Usr2 := uuid.New()
	RandomUUID := uuid.New().String()
	// connect to collection
	coll := mongotestClint.Database("TestEcomDB").Collection("user_test")

	UsrRepo := UserRepo{MongoCollection: coll}

	// Insert user
	t.Run("insert user 1", func(t *testing.T) {
		usr := model.User{
			UserID:    RandomUUID, // Generate MongoDB ObjectID
			Name:      "abhinand",
			Email:     "abhinand@gmail.com",
			Number:    "1234567890",
			CreatedAt: time.Now(),
		}

		result, err := UsrRepo.InsertUser(&usr)
		if err != nil {
			t.Fatal("insert 1 operation failed", err)
		}
		t.Log("insert 1 success", result)
	})

	// get user - 1 data
	t.Run("get user 1", func(t *testing.T) {

		result, err := UsrRepo.FindUserByID(RandomUUID)
		if err != nil {
			t.Fatal("get opration failed", err)
		}

		t.Log("get user 1 success", result.Name)

		time.Sleep(500 * time.Millisecond) // Wait for MongoDB to commit
	})

	// get all users
	t.Run("get all users", func(t *testing.T) {

		result, err := UsrRepo.GetAllUsers()
		if err != nil {
			t.Fatal("get all users failed", err)
		}

		t.Log("get all users success", result)

		time.Sleep(500 * time.Millisecond) // Wait for MongoDB to commit
	})

	// update  user - 1 data
	t.Run("update user 1", func(t *testing.T) {

		usr := model.User{
			Name:   "new name",
			Email:  "nemail@gmail.com",
			Number: "9090909090",
			UserID: Usr1,
		}

		result, err := UsrRepo.UpdateUserById(Usr1, &usr)
		if err != nil {
			t.Fatal("update user 1 failed", err)
		}

		t.Log("update count", result)

		time.Sleep(500 * time.Millisecond) // Wait for MongoDB to commit
	})

	// delete user 1 data
	t.Run("delete user 1", func(t *testing.T) {

		result, err := UsrRepo.DeleteUserById(Usr1)
		if err != nil {
			t.Fatal("delete user 1 failed", err)
		}
		t.Log("delete user 1 success", result)

		time.Sleep(500 * time.Millisecond) // Wait for MongoDB to commit
	})

	// get all user after delete

	t.Run("get all users after delete", func(t *testing.T) {

		result, err := UsrRepo.GetAllUsers()
		if err != nil {
			t.Fatal("get all users failed", err)
		}

		t.Log("get all users success", result)

		time.Sleep(500 * time.Millisecond) // Wait for MongoDB to commit
	})
}
