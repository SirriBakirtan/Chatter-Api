package repositories

import (
	"Chatter-Api/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var _Database *mongo.Database

type _UserRepository struct {
	collection *mongo.Collection
}

type _ConversationRepository struct {
	collection *mongo.Collection
}

func init() {
	clientOptions := options.Client().ApplyURI(config.GetDatabaseConnectionString())
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_Client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = _Client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	_Database = _Client.Database(config.DbName)
	go initConversationRepository()
	go initUserRepository()
}
