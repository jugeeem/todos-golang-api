package persistence

import (
	"github.com/jugeeem/golang-todo.git/app/domain/model"
	"github.com/jugeeem/golang-todo.git/app/domain/repository"
	"gorm.io/gorm"
)

// TodoRepository はTodoRepositoryインターフェースの実装
type TodoRepository struct {
	DB *gorm.DB
}

// NewTodoRepository は新しいTodoRepositoryのインスタンスを作成します
func NewTodoRepository(db *gorm.DB) repository.TodoRepository {
	return &TodoRepository{
		DB: db,
	}
}

// FindByID は指定されたIDのTodoを検索します
func (r *TodoRepository) FindByID(id uint) (*model.Todo, error) {
	var todo model.Todo
	result := r.DB.First(&todo, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return &todo, nil
}

// FindAll はすべてのTodoを取得します
func (r *TodoRepository) FindAll() ([]*model.Todo, error) {
	var todos []*model.Todo
	result := r.DB.Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}

	return todos, nil
}

// FindByUserID は指定されたユーザーIDに関連するTodoを検索します
func (r *TodoRepository) FindByUserID(userID uint) ([]*model.Todo, error) {
	var todos []*model.Todo
	result := r.DB.Where("user_id = ?", userID).Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}

	return todos, nil
}

// Create は新しいTodoを作成します
func (r *TodoRepository) Create(todo *model.Todo) error {
	result := r.DB.Create(todo)

	return result.Error
}

// Update は既存のTodoを更新します
func (r *TodoRepository) Update(todo *model.Todo) error {
	result := r.DB.Save(todo)

	return result.Error
}

// Delete は指定されたIDのTodoを削除します
func (r *TodoRepository) Delete(id uint) error {
	result := r.DB.Delete(&model.Todo{}, id)

	return result.Error
}
