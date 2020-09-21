package repositories

import (
	"Chatter-Api/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var UserRepository _UserRepository

func initUserRepository() {
	UserRepository = _UserRepository{
		collection: _Database.Collection("Users"),
	}
}

func (_UserRepository) GetUsers(userFilter bson.M, opts bson.M) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := UserRepository.collection.Find(
		ctx,
		userFilter)
	var users []models.User
	if err != nil {
		return users, err
	}
	if err := cursor.All(ctx, &users); err != nil {
		return users, err
	}
	return users, nil
}

func (_UserRepository) CreateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := UserRepository.collection.InsertOne(ctx, user)
	return err
}

func (_UserRepository) UpdateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := UserRepository.collection.ReplaceOne(
		ctx,
		bson.M{
			"_id": user.Id,
		},
		user)
	return err
}

func (_UserRepository) DeleteUser(userId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := UserRepository.collection.DeleteOne(
		ctx,
		bson.M{"_id": userId})
	return err
}
