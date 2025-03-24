package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var connString string

func setConn() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	connString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
}

func loadEnv() error {
	err := godotenv.Load("../build/.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
		return err
	}
	return nil
}

func setEnv() error {
	err := loadEnv()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	setConn()
	return nil
}
