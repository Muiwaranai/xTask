package logic

import (
	"errors"
	"fmt"
	"log"
	"todo/data"
	"todo/models"
)

type MyHandler struct {
	db *data.MyDatabase
}

func NewHandler(_db *data.MyDatabase) *MyHandler {
	return &MyHandler{
		db: _db,
	}
}

func (h *MyHandler) HandlerCreateTask(model models.ModelCreateTask) error {

	err := h.db.DatabaseCreateTask(model.UserId, model.Status, model.Title, model.Description)
	if err != nil {
		log.Println("Error with adding task to the database:", err)
		return err
	}

	return nil
}

func (h *MyHandler) HandlerRemoveTask(model models.ModelRemoveTask) error {
	data, err := h.db.DatabaseGetTaskById(model.TaskId)
	if err != nil {
		log.Println("Error with getting task to from database:", err)
		return err
	}

	if data.UserId != model.UserId {
		log.Printf("This user is not authorized to delete chad %d", model.TaskId)
		return errors.New("unauthorized action")
	}

	err = h.db.DatabaseRemoveTaskById(model.TaskId)
	if err != nil {
		log.Println("Error removing task:", err)
		return err
	}

	fmt.Println("Done")

	return nil
}
