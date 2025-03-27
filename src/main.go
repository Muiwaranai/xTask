package main

import (
	"context"
	"log"
	"todo/data"
	"todo/logic"
	"todo/server"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load("../build/.env")
	if err != nil {
		log.Println("Error with setting env", err)
	}

	database := data.NewDatabse(genConnectionString())
	defer database.Conn.Close(context.Background())

	handler := logic.NewHandler(database)
	server := server.NewServer(handler, validator.New())

	e := echo.New()
	e.POST("/api/create", server.ServerCreateTask)
	e.DELETE("/api/remove", server.ServerRemoveTaskById)

	e.Logger.Fatal(e.Start(":3711"))
}
