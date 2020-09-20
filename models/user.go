package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email           string             `json:"email,omitempty" bson:"email,omitempty"`
	HashedPassword  string             `json:"hashedPassword,omitempty" bson:"hashedPassword,omitempty"`
	FirstName       string             `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName        string             `json:"lastName,omitempty" bson:"lastName,omitempty"`
	ThumbnailUrl    string             `json:"thumbnailUrl,omitempty" bson:"thumbnailUrl,omitempty"`
	FriendsIdList   []string           `json:"friends,omitempty" bson:"friends,omitempty"`
	PhotosBucketUrl string             `json:"photosBucketUrl,omitempty" bson:"photosBucketUrl,omitempty"`
	VideosBucketUrl string             `json:"videoBucketUrl,omitempty" bson:"videosBucketUrl,omitempty"`
	Status          bool               `json:"status,omitempty" bson:"status,omitempty"`
	Active          bool               `json:"active,omitempty" bson:"active,omitempty"`
}
