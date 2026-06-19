package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	var err error

	err = godotenv.Load()
	if err != nil {
		panic("env file not found")
	}

	connectDB()
	defer db.Close(context.Background())

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("POST /create-user", createUserHandler)
	mux.HandleFunc("GET /users", getUserHandler)
	mux.HandleFunc("GET /users/{id}", getSingleUserById)
	mux.HandleFunc("PUT /users/{id}", updateUser)
	mux.HandleFunc("DELETE /users/{id}", deleteUser)

	fmt.Println("Server is running on http://localhost:5000")

	err = http.ListenAndServe(":5000", mux)
	if err != nil {
		fmt.Println("Server Error", err)
	}
}
