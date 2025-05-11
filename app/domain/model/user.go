package model

import (
	"time"
)

type User struct {
	ID         uint      `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeleteFlag bool      `json:"delete_flag"`
}

func (User) TableName() string {
	return "users"
}

func NewUser(
	Username string,
	Password string,
	Email string,
) *User {
	now := time.Now()
	return &User{
		Username:  Username,
		Password:  Password,
		Email:     Email,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
