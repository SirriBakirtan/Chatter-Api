package services

import (
	"Chatter-Api/models"
	"Chatter-Api/repositories"
	"errors"
)

var ConversationService _ConversationService

func init() {

}

func (_ConversationService) GetConversation(userId string, conversation models.Conversation) (models.Conversation, error) {
	conversation, err := repositories.ConversationRepository.GetConversation(conversation)
	if err != nil {
		return models.Conversation{}, err
	}
	for _, party := range conversation.Parties {
		if party == userId {
			return conversation, nil
		}
	}
	return models.Conversation{}, errors.New("user_not_included")
}
