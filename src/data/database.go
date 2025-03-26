package data

import (
	"context"
	"fmt"
	"log"
	"todo/models"

	"github.com/jackc/pgx/v5"
)

type MyDatabase struct {
	Conn *pgx.Conn
}

func NewDatabse(connString string) *MyDatabase {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Println("Error with connection", err)
	}

	return &MyDatabase{Conn: conn}
}

func (db *MyDatabase) CreateTaskDB(task models.HelperCreateTask) error {
	query := `
		INSERT INTO tasks (UserId, Status, Title, Description)
		VALUES ($1, $2, $3, $4)
	`
	_, err := db.Conn.Exec(context.Background(), query, task.UserId, task.Status, task.Title, task.Description)
	if err != nil {
		log.Println("Error inserting task:", err)
		return err
	}

	return nil
}

func (db *MyDatabase) RemoveTaskDB(id int) error {
	query := `
		Remove task from db
	`
	fmt.Println(query)
	return nil
}
