package service

import (
	"TODOapp/internal/model"
	"TODOapp/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ToDoList interface {
	Create(userId int, list model.TodoList) (int, error)
	GetAll(userId int) ([]model.TodoList, error)
	GetById(userId int, liistId int) (model.TodoList, error)
	Delete(userId, liistId int) error
	Update(userId, listId int, input model.UpadteListInput) error
}

type ToDoItem interface {
	Create(userId, listId int, item model.TodoItem) (int, error)
	GetAll(userId int, listId int) ([]model.TodoItem, error)
	GetById(userId, itemId int) (model.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input model.UpadteItemInput) error
}

type Service struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		ToDoList:      NewTodoListService(repo.ToDoList),
		ToDoItem:      NewItemService(repo.ToDoItem, repo.ToDoList),
	}
}
