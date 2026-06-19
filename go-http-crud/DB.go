package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn

func connectDB() {
	var err error

	connectStr := os.Getenv("DATABASE_URL")

	db, err = pgx.Connect(context.Background(), connectStr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database Connected Successfully")
}
