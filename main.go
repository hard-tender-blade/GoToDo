package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
}

var tasks []Task = []Task{
	{
		ID: 		"1",
		Title: 		"Task 1",
		Status: 	true,
		Description: "This is task 1",
	},
	{
		ID: 		"2",
		Title: 		"Task 2",
		Status: 	false,
		Description: "This is task 2",
	},
}

func main() {
	r := mux.NewRouter() //? Create a new router, some mux router

	r.HandleFunc("/", HelloWorld) //? adress request to the func
	r.HandleFunc("/tasks", GetTasks).Methods("GET")
	
	http.Handle("/", r)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")	
	json.NewEncoder(w).Encode(tasks)
}