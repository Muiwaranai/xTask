package main

import (
	"context"
	"log"
	"todo/data"
	"todo/server"

	"github.com/labstack/echo/v4"
)

func main() {
	err := setEnv() // connString
	if err != nil {
		log.Fatal("Error with starting prepearing environment")
	}

	db := data.NewDatabse(connString)
	defer db.Conn.Close(context.Background())

	s := server.NewServer(db)

	e := echo.New()
	e.POST("/api/create", s.CreateTask)

	e.Logger.Fatal(e.Start(":3711"))
}
