package repo

import (
	"context"
	"fmt"

	"github.com/abhinandpn/MongoDB-Golang/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	MongoCollection *mongo.Collection
}

func (r *UserRepo) InsertUser(usr *model.User) (interface{}, error) {

	result, err := r.MongoCollection.InsertOne(context.Background(), usr)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (r *UserRepo) FindUserByID(userID string) (*model.User, error) {
    var user model.User
    err := r.MongoCollection.FindOne(context.TODO(), bson.M{"user_id": userID}).Decode(&user)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepo) GetAllUsers() ([]model.User, error) {

	result, err := r.MongoCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	var users []model.User
	err = result.All(context.Background(), &users)
	if err != nil {
		return nil, fmt.Errorf("error while decoding users: %s", err.Error())
	}

	return users, nil
}

func (r *UserRepo) UpdateUserById(usrId string, UpdateUsr *model.User) (int64, error) {

	result, err := r.MongoCollection.UpdateOne(context.Background(),
		bson.D{{Key: "user_id", Value: usrId}},
		bson.D{{Key: "$set", Value: UpdateUsr}})

	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

func (r *UserRepo) DeleteUserById(usrID string) (int64, error) {

	result, err := r.MongoCollection.DeleteOne(context.Background(),
		bson.D{{Key: "user_id", Value: usrID}})

	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

func (r *UserRepo) DeleteAllUsers() (int64, error) {
	result, err := r.MongoCollection.DeleteMany(context.Background(),
		bson.D{})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
