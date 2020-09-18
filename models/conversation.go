package models

type Conversation struct {
	ID          string           `json:"_id,omitempty" bson:"_id,omitempty"`
	IsGroupChat bool             `json:"isGroupChat,omitempty" bson:"isGroupChat,omitempty"`
	Parties     []string         `json:"parties,omitempty" bson:"parties,omitempty"`
	Messages    []_MessageObject `json:"messages,omitempty" bson:"messages,omitempty"`
}

type _MessageObject struct {
	SenderId string `json:"sender_id,omitempty" bson:"sender_id,omitempty"`
	Message  string `json:"message,omitempty" bson:"message,omitempty"`
}
