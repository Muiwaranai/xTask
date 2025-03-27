package models

type MyUser struct {
	Id       int    `json:Id`
	Username string `json:Username`
}

type MyTask struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	Status      int    `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ModelCreateTask struct {
	UserId      int    `json:"UserId" validate:"required,min=1"`
	Status      int    `json:"Status" validate:"required,oneof=0 1 2"`
	Title       string `json:"Title" validate:"required,min=3,max=100"`
	Description string `json:"Description" validate:"max=500"`
}

type ModelRemoveTask struct {
	UserId int `json:"UserId" validate:"required,min=1"`
	TaskId int `json:"TaskId" validate:"required,min=1"`
}
