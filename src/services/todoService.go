package services

import (
	"golang-serverless-sample/src/domain/todoDomain"
	"log"
)

type TodoService struct {
	todoRepository todoDomain.Repository
}

func NewTodoService(todoRepository todoDomain.Repository) todoDomain.Service {
	return &TodoService{todoRepository: todoRepository}
}

func (t TodoService) FindById(id string) (*todoDomain.Model, error) {
	result, err := t.todoRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t TodoService) Create(todo *todoDomain.Model) (*todoDomain.Model, error) {
	result, err := t.todoRepository.Create(todo)
	if err != nil {
		return nil, err
	}
	log.Println("TodoService.Create - Created todo with id: ", result.Id)
	return result, nil
}

func (t TodoService) Update(id string, todo *todoDomain.Model) (*todoDomain.Model, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoService) FindAll() (*[]todoDomain.Model, error) {
	result, err := t.todoRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t TodoService) Delete(id string) error {
	_, err := t.todoRepository.GetById(id)
	if err != nil {
		return err
	}
	err = t.todoRepository.Delete(id)
	log.Println("TodoService.Delete - Deleted todo with id: ", id)
	if err != nil {
		return err
	}
	return nil
}
