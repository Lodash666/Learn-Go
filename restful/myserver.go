package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//http://www.somkiat.cc/restful-api-with-go/
func main() {
	http.HandleFunc("/", addHandler)
	http.ListenAndServe(":8080", nil)
}
func addHandler(w http.ResponseWriter, r *http.Request) {
	type Task struct {
		ID    int64
		Title string
		Done  bool
	}
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	var resultTask Task
	resultTask.ID = task.ID
	resultTask.Title = task.Title
	resultTask.Done = task.Done

	if err := json.NewEncoder(w).Encode(resultTask); err != nil {
		log.Println(err)
		http.Error(w, "Oops", http.StatusInternalServerError)
	}
}
