package dto

import "github.com/jugeeem/golang-todo.git/app/domain/model"

// UserResponse はユーザー情報を表す構造体です
type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Userモデルから必要なフィールドだけを取り出すマッパー関数
func ToUserResponse(user *model.User) *UserResponse {
	return &UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}

// スライス変換用のヘルパー関数
func ToUserResponseList(users []*model.User) []*UserResponse {
	result := make([]*UserResponse, len(users))
	for i, user := range users {
		result[i] = ToUserResponse(user)
	}
	return result
}
