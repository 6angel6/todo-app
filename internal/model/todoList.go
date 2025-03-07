package model

import "errors"

// TodoList @Summary		TodoList
// @Description	A todo list
type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"Title" binding:"required"`
	Description string `json:"Description"`
}

// UserList @Summary		UserList
// @Description	User to List association
type UserList struct {
	Id     int
	UserId int
	ListId int
}

// TodoItem @Summary		TodoItem
// @Description	A todo item
type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

// List @Summary		List
// @Description	A list to item association
type List struct {
	Id     int
	ListId int
	ItemId int
}
type UpadteListInput struct {
	Title       *string `json:"Title"`
	Description *string `json:"Description"`
}

type UpadteItemInput struct {
	Title       *string `json:"Title"`
	Description *string `json:"Description"`
	Done        *bool   `json:"Done"`
}

func (i UpadteItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
