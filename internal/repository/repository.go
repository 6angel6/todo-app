package repository

import (
	"TODOapp/internal/model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type ToDoList interface {
	Create(userId int, list model.TodoList) (int, error)
	GetAll(userId int) ([]model.TodoList, error)
	GetById(userId int, listId int) (model.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input model.UpadteListInput) error
}

type ToDoItem interface {
	Create(listId int, item model.TodoItem) (int, error)
	GetAll(userId int, listId int) ([]model.TodoItem, error)
	GetById(userId, itemId int) (model.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input model.UpadteItemInput) error
}

type Repository struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		ToDoList:      NewTodoListPostgres(db),
		ToDoItem:      NewItemPostgres(db),
	}
}
