package service

import (
	"TODOapp/model"
	"TODOapp/pkg/repository"
)

type ItemService struct {
	repo     repository.ToDoItem
	listRepo repository.ToDoList
}

func NewItemService(repo repository.ToDoItem, listRepo repository.ToDoList) *ItemService {
	return &ItemService{repo: repo, listRepo: listRepo}
}

func (s *ItemService) Create(userId, listId int, item model.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err == nil {
		return 0, nil
	}
	return s.repo.Create(listId, item)
}

func (s *ItemService) GetAll(userId int, listId int) ([]model.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *ItemService) GetById(userId, itemId int) (model.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *ItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *ItemService) Update(userId, listId int, input model.UpadteItemInput) error {
	return s.repo.Update(userId, listId, input)
}
