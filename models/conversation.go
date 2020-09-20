package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Conversation struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Parties  []string           `json:"parties,omitempty" bson:"parties,omitempty"`
	Messages []MessageObject
}

type MessageObject struct {
	SenderId string `json:"senderId,omitempty" bson:"senderId,omitempty"`
	Message  string `json:"message,omitempty" bson:"message,omitempty"`
}
