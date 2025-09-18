package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go-rest-api/db"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

//how you create a type in go 
type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}


func main() {

	client := db.ConnectDB()
	userCollection = client.Database("isfdb").Collection("users")

	r := mux.NewRouter()
	r.HandleFunc("api/users", GetUsers).Methods("GET")
}


func GetClient(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	defer cursor.Close(ctx)


	var users []User 
	if err = cursor.All(ctx, &users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	json.NewEncoder(w).Encode(users)

}


func CreateClient(w http.ResponseWriter, r *http.Request) {
	var newUser User
	_ = json.NewDecoder(r.Body).Decode(&newUser)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newUser)
}


