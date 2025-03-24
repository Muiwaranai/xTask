package server

import (
	"net/http"
	"todo/data"
	"todo/models"

	"github.com/labstack/echo/v4"
)

type MyServer struct {
	db *data.MyDatabase
}

func NewServer(_database *data.MyDatabase) *MyServer {
	return &MyServer{
		db: _database,
	}
}

func (s *MyServer) CreateTask(c echo.Context) error {
	var data models.HelperCreateTask
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	err := s.db.CreateTaskDB(data)
	if err != nil {
		// todo error with adding to database
	}

	return c.JSON(http.StatusCreated, data)
}
