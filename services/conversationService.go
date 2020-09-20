package services

import (
	"Chatter-Api/models"
	"Chatter-Api/repositories"
	"go.mongodb.org/mongo-driver/bson"
)

var ConversationService _ConversationService

func init() {

}

func (_ConversationService) GetConversationsOfUser(userId string) ([]models.Conversation, error) {
	return repositories.ConversationRepository.GetConversations(
		bson.M{"parties": userId},
		bson.M{})
}
