package services

import (
	"Chatter-Api/models"
	"Chatter-Api/repositories"
	"go.mongodb.org/mongo-driver/bson"
)

var UserService _UserService

func init() {

}

func (_UserService) Login(user models.User) error {
	_, err := repositories.UserRepository.GetUsers(
		bson.M{
			"email":          user.Email,
			"hashedPassword": user.HashedPassword,
		},
		bson.M{})
	return err
}

func (_UserService) SignUp(user models.User) error {
	return repositories.UserRepository.CreateUser(user)
}
