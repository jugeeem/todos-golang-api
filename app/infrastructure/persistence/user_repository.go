package persistence

import (
	"github.com/jugeeem/golang-todo.git/app/domain/model"
	"github.com/jugeeem/golang-todo.git/app/domain/repository"
	"gorm.io/gorm"
)

// UserRepository はUserRepositoryインターフェースの実装
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository は新しいUserRepositoryのインスタンスを作成します
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// FindByID はIDでユーザーを検索します
func (r *UserRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	result := r.DB.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return &user, nil
}

// FindByUsername はユーザー名でユーザーを検索します
func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	result := r.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return &user, nil
}

// FindByEmail はメールアドレスでユーザーを検索します
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	result := r.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return &user, nil
}

// FindByUsernameAndPassword はユーザー名とパスワードでユーザーを検索します
func (r *UserRepository) FindByUsernameAndPassword(username, password string) (*model.User, error) {
	var user model.User
	result := r.DB.Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return &user, nil
}

// FindByEmailAndPassword はメールアドレスとパスワードでユーザーを検索します
func (r *UserRepository) FindByEmailAndPassword(email, password string) (*model.User, error) {
	var user model.User
	result := r.DB.Where("email = ? AND password = ?", email, password).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return &user, nil
}

// FindByUsernameOrEmail はユーザー名またはメールアドレスでユーザーを検索します
func (r *UserRepository) FindByUsernameOrEmail(username, email string) (*model.User, error) {
	var user model.User
	result := r.DB.Where("username = ? OR email = ?", username, email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return &user, nil
}

// FindByUsernameAndEmail はユーザー名とメールアドレスでユーザーを検索します
func (r *UserRepository) FindByUsernameAndEmail(username, email string) (*model.User, error) {
	var user model.User
	result := r.DB.Where("username = ? AND email = ?", username, email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return &user, nil
}

// FindAll は全てのユーザーを取得します
func (r *UserRepository) FindAll() ([]*model.User, error) {
	var users []*model.User
	result := r.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

// Create は新しいユーザーを作成します
func (r *UserRepository) Create(user *model.User) (*model.User, error) {
	result := r.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// Update は既存のユーザーを更新します
func (r *UserRepository) Update(user *model.User) (*model.User, error) {
	result := r.DB.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// Remove はユーザーを削除します
func (r *UserRepository) Remove(id uint) error {
	result := r.DB.Delete(&model.User{}, id)

	return result.Error
}
