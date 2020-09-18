package repositories

import (
	"Chatter-Api/models"
	"context"
	"errors"
	"time"
)

var UserRepository _UserRepository

func initUserRepository() {
	UserRepository = _UserRepository{
		collection: _Database.Collection("Users"),
	}
}

func (_UserRepository) GetUser(user models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := UserRepository.collection.FindOne(ctx, user).Decode(&user); err != nil {
		return models.User{}, errors.New("failed")
	}
	return user, nil
}

func (_UserRepository) CreateUser(user models.User) bool {
	return true
}

func (_UserRepository) UpdateUser(user models.User) bool {
	return true
}

func (_UserRepository) DeleteUser(userId string) bool {
	return true
}
