package models

// type MyUser struct {
// 	Id       int    `json:Id`
// 	Username string `json:Username`
// }

// type MyTask struct {
// 	Id          int    `json:Id`
// 	UserId      int    `json:"UserId"`
// 	Status      int    `json:"Status"`
// 	Title       string `json:"Title"`
// 	Description string `json:"Description"`
// }

type HelperCreateTask struct {
	UserId      int    `json:"UserId"`
	Status      int    `json:"Status"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}
