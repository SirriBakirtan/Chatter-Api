package models

type User struct {
	Id             string `json:"_id,omitempty" bson:"_id,omitempty"`
	Username       string `json:"username,omitempty" bson:"username,omitempty"`
	HashedPassword string `json:"hashedPassword,omitempty" bson:"hashedPassword,omitempty"`
}
