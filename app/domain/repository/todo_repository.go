package repository

import "github.com/jugeeem/golang-todo.git/app/domain/model"

// TodoRepository はTodoの永続化を担当するインターフェース
type TodoRepository interface {
	FindByID(id uint) (*model.Todo, error)
	FindAll() ([]*model.Todo, error)
	FindByUserID(userID uint) ([]*model.Todo, error)
	Create(todo *model.Todo) error
	Update(todo *model.Todo) error
	Delete(id uint) error
}
