package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//how you create a type in go 
type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}


