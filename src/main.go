package main

import (
	"context"
	"log"
	"todo/data"
	"todo/logic"
	"todo/server"

	"github.com/labstack/echo/v4"
)

func main() {
	err := setEnv() // connString
	if err != nil {
		log.Fatal("Error with starting prepearing environment")
	}

	database := data.NewDatabse(connString)
	defer database.Conn.Close(context.Background())

	handler := logic.NewHandler(database)
	server := server.NewServer(handler)

	e := echo.New()
	e.POST("/api/create", server.CreateTask)
	e.POST("/api/remove", server.RemoveTaskById)

	e.Logger.Fatal(e.Start(":3711"))
}
