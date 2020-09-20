package repositories

import (
	"Chatter-Api/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var ConversationRepository _ConversationRepository

func initConversationRepository() {
	ConversationRepository = _ConversationRepository{
		collection: _Database.Collection("Conversations"),
	}
}

func (_ConversationRepository) GetConversations(conversationFilter bson.M, opts bson.M) ([]models.Conversation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := ConversationRepository.collection.Find(ctx, conversationFilter)
	var conversationArray []models.Conversation
	if err != nil {
		return conversationArray, err
	}
	if err := cursor.All(ctx, &conversationArray); err != nil {
		return conversationArray, err
	}
	return conversationArray, nil
}

func (_ConversationRepository) CreateConversation(conversation models.Conversation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := ConversationRepository.collection.InsertOne(ctx, conversation)
	return err
}

func (_ConversationRepository) UpdateConversation(conversation models.Conversation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := ConversationRepository.collection.ReplaceOne(
		ctx,
		bson.M{"_id": conversation.Id},
		conversation)
	return err
}

func (_ConversationRepository) DeleteConversation(conversationId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := ConversationRepository.collection.DeleteOne(
		ctx,
		bson.M{"_id": conversationId})
	return err
}
