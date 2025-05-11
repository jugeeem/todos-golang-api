package dto

import "github.com/jugeeem/golang-todo.git/app/domain/model"

type TodoResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	UserID      uint   `json:"user_id"`
}

// Todoモデルから必要なフィールドだけを取り出すマッパー関数
func ToTodoResponse(todo *model.Todo) *TodoResponse {
	return &TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
		UserID:      todo.UserID,
	}
}

// スライス変換用のヘルパー関数
func ToTodoResponseList(todos []*model.Todo) []*TodoResponse {
	result := make([]*TodoResponse, len(todos))
	for i, todo := range todos {
		result[i] = ToTodoResponse(todo)
	}
	return result
}
