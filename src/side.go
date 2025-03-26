package main

import (
	"fmt"
	"os"
)

func genConnectionString() string {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
		fmt.Println("Error: One or more environment variables are missing.")
		os.Exit(1)
	}
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
}
