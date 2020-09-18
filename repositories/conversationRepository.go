package repositories

import (
	"Chatter-Api/models"
	"context"
	"errors"
	"time"
)

var ConversationRepository _ConversationRepository

func initConversationRepository() {
	ConversationRepository = _ConversationRepository{
		collection: _Database.Collection("Conversations"),
	}
}

func (_ConversationRepository) GetConversation(conversation models.Conversation) (models.Conversation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ConversationRepository.collection.FindOne(ctx, conversation).Decode(&conversation); err != nil {
		return models.Conversation{}, errors.New("failed_to_get_messages")
	}
	return conversation, nil
}

func (_ConversationRepository) CreateConversation(conversation models.Conversation) bool {
	return true
}

func (_ConversationRepository) UpdateConversation(conversation models.Conversation) bool {
	return true
}

func (_ConversationRepository) DeleteConversation(conversation string) bool {
	return true
}
