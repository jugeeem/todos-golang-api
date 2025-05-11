package repository

import (
	"github.com/jugeeem/golang-todo.git/app/domain/model"
)

// UserRepository はユーザー情報の永続化を担当するインターフェース
type UserRepository interface {
	FindByID(id uint) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByUsernameAndPassword(username, password string) (*model.User, error)
	FindByEmailAndPassword(email, password string) (*model.User, error)
	FindByUsernameOrEmail(username, email string) (*model.User, error)
	FindByUsernameAndEmail(username, email string) (*model.User, error)
	FindAll() ([]*model.User, error)
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Remove(id uint) error
}
