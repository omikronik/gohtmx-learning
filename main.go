package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"gohtmx-learning/templates"
	"gohtmx-learning/todo"

	"github.com/joho/godotenv"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Root hit")
	http.ServeFile(w, r, "web/index.html")
}

func CreateToDoItemsHandler(items *[]todo.TodoItem) http.HandlerFunc {
	fmt.Println("Create to do items hit")
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var t todo.TodoItemRequest
		err := decoder.Decode(&t)
		if err != nil {
			fmt.Println("Something happen")
			panic(err)
		}
		items := append(*items, todo.NewToDoItem(t.Title, t.Content, false))
		w.Header().Set("Content-Type", "application/json")
		templates.TodoList(items).Render(r.Context(), w)
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
		todo.NewToDoItem("Task 1", "This is the content for task 1", false),
		todo.NewToDoItem("Task 2", "This is the content for task 2", false),
	}

	server := http.NewServeMux()
	server.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	server.HandleFunc("POST /createtodoitem", CreateToDoItemsHandler(&todoItems))
	server.HandleFunc("/", RootHandler)
	fmt.Printf("Started server on port: %s\n", PORT)
	if err := http.ListenAndServe(PORT, server); err != nil {
		fmt.Print(err)
	}
}
