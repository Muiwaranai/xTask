package data

import (
	"context"
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

	return &MyDatabase{
		Conn: conn,
	}
}

func (db *MyDatabase) DatabaseCreateTask(_userId int, _status int, _title string, _description string) error {
	query := `
		INSERT INTO tasks (UserId, Status, Title, Description)
		VALUES ($1, $2, $3, $4)
	`
	_, err := db.Conn.Exec(context.Background(), query, _userId, _status, _title, _description)
	if err != nil {
		log.Println("Error inserting task:", err)
		return err
	}

	return nil
}

func (db *MyDatabase) DatabaseGetTaskById(_id int) (*models.MyTask, error) {
	query := `
		SELECT id, userid, status, title, description FROM tasks
		WHERE id = $1
	`

	var data models.MyTask

	err := db.Conn.QueryRow(context.Background(), query, _id).Scan(
		&data.Id, &data.UserId, &data.Status, &data.Title, &data.Description,
	)
	if err != nil {
		log.Println("Error getting task:", err)
		return nil, err
	}

	return &data, nil
}

func (db *MyDatabase) DatabaseRemoveTaskById(_id int) error {
	query := `
		DELETE FROM tasks 
		WHERE id = $1
	`

	_, err := db.Conn.Exec(context.Background(), query, _id)
	if err != nil {
		log.Println("Error removing task:", err)
		return err
	}

	return nil
}
