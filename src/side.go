package main

import (
	"fmt"
	"log"
	"os"
)

func genConnectionString() string {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	if dbUser == "" || dbHost == "" || dbPort == "" || dbName == "" { // There is no password required in my app cuz its local
		log.Fatal("Error: One or more environment variables are missing.")
	}
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
}

// func genApp() *server.MyServer {
// 	database := data.NewDatabse(genConnectionString())
// 	defer database.Conn.Close(context.Background())

// 	handler := logic.NewHandler(database)
// 	server := server.NewServer(handler, validator.New())
// 	return server
// }
