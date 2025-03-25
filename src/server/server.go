package server

import (
	"net/http"
	"todo/logic"
	"todo/models"

	"github.com/labstack/echo/v4"
)

type MyServer struct {
	handler *logic.MyHandler
}

func NewServer(_handler *logic.MyHandler) *MyServer {
	return &MyServer{
		handler: _handler,
	}
}

func (s *MyServer) CreateTask(c echo.Context) error {
	var data models.HelperCreateTask
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, "error invalid request")
	}

	return c.JSON(http.StatusCreated, data)
}

func (s *MyServer) RemoveTaskById(c echo.Context) error {
	var id int

	return c.JSON(http.StatusCreated, id)
}
