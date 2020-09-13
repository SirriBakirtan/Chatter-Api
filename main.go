package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}

type MessageObject struct {
	Sender_id primitive.ObjectID `json:"sender_id,omitempty" bson:"sender_id,omitempty"`
	Message   string             `json:"message,omitempty" bson:"message,omitempty"`
}

type Conversation struct {
	ID       primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Parties  []primitive.ObjectID `json:"parties,omitempty" bson:"parties,omitempty"`
	Messages []MessageObject      `json:"messages,omitempty" bson:"messages,omitempty"`
}

type endpoints struct {
	client *mongo.Client
}

func NewEndpoints(client *mongo.Client) *endpoints {
	return &endpoints{
		client: client,
	}
}

func (e *endpoints) GetUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var user User
	collection := e.client.Database("test").Collection("admin")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, User{ID: id}).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(user)
}

func (e *endpoints) GetUserMessageEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var conversation Conversation
	collection := e.client.Database("test").Collection("messages")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, Conversation{ID: id}).Decode(&conversation)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message1": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(conversation)
}

func (e *endpoints) CreateUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var user User
	_ = json.NewDecoder(request.Body).Decode(&user)
	collection := e.client.Database("test").Collection("admin")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	fmt.Println(user.Username)
	result, _ := collection.InsertOne(ctx, user)
	json.NewEncoder(response).Encode(result)
}

func (e *endpoints) GetAllUsersEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var people []User
	collection := e.client.Database("test").Collection("admin")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user User
		cursor.Decode(&user)
		people = append(people, user)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(people)
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://sirribakirtan26:Is99SM4ON7e9PUPw@databases.fby8y.mongodb.net/<dbname>?retryWrites=true&w=majority")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	endpoints := NewEndpoints(client)
	router := mux.NewRouter()
	router.HandleFunc("/new_user", endpoints.CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/get_users", endpoints.GetAllUsersEndpoint).Methods("GET")
	router.HandleFunc("/user/{id}", endpoints.GetUserEndpoint).Methods("GET")
	router.HandleFunc("/message/{id}", endpoints.GetUserMessageEndpoint).Methods("GET")
	http.ListenAndServe(":12345", router)
}
