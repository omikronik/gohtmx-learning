package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"gohtmx-learning/todo"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/index.html")
}

func GetToDoItemsHandler(items []todo.TodoItem) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	}
}

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Print("No .env found")
	}
}

func main() {
	PORT, exists := os.LookupEnv("PORT")
	if !exists {
		fmt.Println("PORT env not found")
	}

	todoItems := []todo.TodoItem{
		{id: uuid.NewString(), title: "Task 1", content: "This is the content for task 1", complete: false, createdOn: time.Now()},
		{id: uuid.NewString(), title: "Task 2", content: "This is the content for task 2", complete: false, createdOn: time.Now()},
	}

	server := http.NewServeMux()
	server.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	server.HandleFunc("/", RootHandler)
	server.HandleFunc("GET /todoitems", GetToDoItemsHandler(todoItems))

	fmt.Printf("Started server on port: %s\n", PORT)
	if err := http.ListenAndServe(PORT, server); err != nil {
		fmt.Print(err)
	}
}
