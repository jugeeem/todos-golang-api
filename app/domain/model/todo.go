package model

import (
	"time"
)

type Todo struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	UserID      uint      `json:"user_id"` // 追加: ユーザーIDフィールド
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TableName はTodoモデルのテーブル名を返します
func (Todo) TableName() string {
	return "todos"
}

// NewTodo は新しいTodoを作成します
func NewTodo(title string, description string, userID uint) *Todo {
	now := time.Now()
	return &Todo{
		Title:       title,
		Description: description,
		Completed:   false,
		UserID:      userID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// ToggleCompleted はタスクの完了状態を切り替えます
func (t *Todo) ToggleCompleted() {
	t.Completed = !t.Completed
	t.UpdatedAt = time.Now()
}

// UpdateTitle はタスクのタイトルを更新します
func (t *Todo) UpdateTitle(title string, description string) {
	t.Title = title
	t.Description = description
	t.UpdatedAt = time.Now()
}
