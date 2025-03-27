package server

import (
	"fmt"
	"net/http"
	"strconv"
	"todo/logic"
	"todo/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type MyServer struct {
	handler   *logic.MyHandler
	validator *validator.Validate
}

func NewServer(_handler *logic.MyHandler, _validator *validator.Validate) *MyServer {
	return &MyServer{
		handler:   _handler,
		validator: _validator,
	}
}

func (s *MyServer) ServerCreateTask(c echo.Context) error {
	var data models.ModelCreateTask

	if err := c.Bind(&data); err != nil {
		fmt.Println("Bind Error:", err)
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	if err := s.validator.Struct(data); err != nil {
		fmt.Println("Validation Error:", err)
		return c.JSON(http.StatusBadRequest, "Input data doesn't satisfy requirements")
	}

	s.handler.HandlerCreateTask(data)

	return c.JSON(http.StatusCreated, data)
}

func (s *MyServer) ServerRemoveTaskById(c echo.Context) error {
	var data models.ModelRemoveTask
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	if err := s.validator.Struct(data); err != nil {
		fmt.Println("Validation Error:", err)
		return c.JSON(http.StatusBadRequest, "Input data doesn't satisfy requirements")
	}

	err := s.handler.HandlerRemoveTask(data)
	if err != nil {
		fmt.Println("rror with deleting task:", err)
		return c.JSON(http.StatusBadRequest, "Error with deleting task")
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("Task %s removed", strconv.Itoa(data.TaskId)))
}
