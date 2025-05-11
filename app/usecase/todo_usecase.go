package usecase

import (
	"errors"

	"github.com/jugeeem/golang-todo.git/app/domain/model"
	"github.com/jugeeem/golang-todo.git/app/domain/repository"
)

// TodoUseCase はTodoアプリケーションユースケースを提供します
type TodoUseCase struct {
	todoRepo repository.TodoRepository
}

// NewTodoUseCase は新しいTodoUseCaseのインスタンスを作成します
func NewTodoUseCase(todoRepo repository.TodoRepository) *TodoUseCase {
	return &TodoUseCase{
		todoRepo: todoRepo,
	}
}

// GetAllTodos は全てのTodoタスクを取得します
func (uc *TodoUseCase) GetAllTodos() ([]*model.Todo, error) {
	return uc.todoRepo.FindAll()
}

// GetTodoByID は指定されたIDのTodoタスクを取得します
func (uc *TodoUseCase) GetTodoByID(id uint) (*model.Todo, error) {
	return uc.todoRepo.FindByID(id)
}

// GetTodosByUserID は指定されたユーザーIDのTodoタスクを取得します
func (uc *TodoUseCase) GetTodosByUserID(userID uint) ([]*model.Todo, error) {
	return uc.todoRepo.FindByUserID(userID)
}

// CreateTodo は新しいTodoタスクを作成します
func (uc *TodoUseCase) CreateTodo(
	title string,
	description string,
	userID uint,
) (*model.Todo, error) {
	if title == "" {
		return nil, errors.New("タイトルは必須です")
	}
	if userID == 0 {
		return nil, errors.New("ユーザーIDは必須です")
	}
	todo := model.NewTodo(title, description, userID)
	err := uc.todoRepo.Create(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

// UpdateTodo は既存のTodoタスクを更新します
func (uc *TodoUseCase) UpdateTodo(
	id uint,
	title string,
	description string,
	completed *bool,
	currentUserID uint,
) (*model.Todo, error) {
	todo, err := uc.todoRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if todo.UserID != currentUserID {
		return nil, errors.New("このTodoを編集する権限がありません")
	}
	if title != "" {
		todo.UpdateTitle(title, description)
	}
	if completed != nil {
		if *completed != todo.Completed {
			todo.ToggleCompleted()
		}
	}
	err = uc.todoRepo.Update(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

// DeleteTodo は指定されたIDのTodoタスクを削除します
func (uc *TodoUseCase) DeleteTodo(id uint, currentUserID uint) error {
	todo, err := uc.todoRepo.FindByID(id)
	if err != nil {
		return err
	}
	if todo.UserID != currentUserID {
		return errors.New("このTodoを削除する権限がありません")
	}

	return uc.todoRepo.Delete(id)
}
