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
		ID:          "1",
		Title:       "Task 1",
		Status:      true,
		Description: "This is task 1",
	},
	{
		ID:          "2",
		Title:       "Task 2",
		Status:      false,
		Description: "This is task 2",
	},
}

func main() {
	r := mux.NewRouter() //? Create a new router, some mux router

	//# Our routes
	r.HandleFunc("/", HelloWorld)
	r.HandleFunc("/tasks", GetTasks).Methods("GET")
	r.HandleFunc("/tasks", CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")

	http.Handle("/", r)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	//# Get id from request
	params := mux.Vars(r)
	id := params["id"]

	//# Loop through tasks slice and delete task with id
	for index, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:index], tasks[index+1:]...)
			w.WriteHeader(http.StatusOK) //# Set status code to 200
			break
		}
	}
	w.WriteHeader(http.StatusNotFound) //# Set status code to 404
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//# Get data from request body and decode it to Task struct
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		fmt.Println(err)
	}

	//# Append new task to tasks slice
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
