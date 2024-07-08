package model

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type List struct {
	Id     int
	ListId int
	ItemId int
}
