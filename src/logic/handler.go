package logic

import "todo/data"

type MyHandler struct {
	db *data.MyDatabase
}

func NewHandler(_db *data.MyDatabase) *MyHandler {
	return &MyHandler{
		db: _db,
	}
}
