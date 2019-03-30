package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:firstman,"omitempty"`
	Lastname  string `json:lastname,"omitempty"` //handle validate if it null data it will be to default of datatype
	//example if string it null data it will be ""
}

var users []User

func GetUsers(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(users)
}
func GetUserById(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	//if not found retrn empty object with user structure

	json.NewEncoder(w).Encode(&User{})
}
func CreateUser(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user) //mapping Struc of req.body & defind struc
	user.ID = params["id"]
	users = append(users, user)

	json.NewEncoder(w).Encode(users)
}
func DeleteUser(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

func main() {
	fmt.Println("Friendship is Magic on port 8081")

	users = append(users, User{ID: "1", Firstname: "Inwza", Lastname: "what"})
	users = append(users, User{ID: "2", Firstname: "YOYYOa", Lastname: "Where"})

	router := mux.NewRouter()
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUserById).Methods("GET")
	router.HandleFunc("/users/{id}", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", router))
}
