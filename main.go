package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type TodoItem struct {
	Id       int
	title    string
	content  string
	createOn time.Time
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/index.html")
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

	server := http.NewServeMux()

	server.HandleFunc("/", RootHandler)

	fmt.Printf("Started server on port: %s\n", PORT)
	if err := http.ListenAndServe(PORT, server); err != nil {
		fmt.Print(err)
	}
}
