package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/jugeeem/golang-todo.git/app/domain/model"
	"github.com/jugeeem/golang-todo.git/app/domain/repository"
	"github.com/jugeeem/golang-todo.git/app/utility"
)

// UserUseCase はユーザーアプリケーションユースケースを提供します
type UserUseCase struct {
	userRepo repository.UserRepository
}

// NewUserUseCase はUserUseCaseの新しいインスタンスを作成します
func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

// GetAllUsers は全てのユーザーを取得します
func (uc *UserUseCase) GetAllUsers() ([]*model.User, error) {
	return uc.userRepo.FindAll()
}

// GetUserByID は指定されたIDのユーザーを取得します
func (uc *UserUseCase) GetUserByID(id uint) (*model.User, error) {
	return uc.userRepo.FindByID(id)
}

// GetUserByUsername はユーザー名でユーザーを検索します
func (uc *UserUseCase) GetUserByUsername(username string) (*model.User, error) {
	return uc.userRepo.FindByUsername(username)
}

// GetUserByEmail はメールアドレスでユーザーを検索します
func (uc *UserUseCase) GetUserByEmail(email string) (*model.User, error) {
	return uc.userRepo.FindByEmail(email)
}

// GetUserByUsernameAndPassword はユーザー名とパスワードでユーザーを検索します
func (uc *UserUseCase) GetUserByUsernameAndPassword(username, password string) (*model.User, error) {
	user, err := uc.userRepo.FindByUsernameAndPassword(username, password)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("ユーザーが見つかりません")
	}

	return user, nil
}

// GetUserByEmailAndPassword はメールアドレスとパスワードでユーザーを検索します
func (uc *UserUseCase) GetUserByEmailAndPassword(email, password string) (*model.User, error) {
	user, err := uc.userRepo.FindByEmailAndPassword(email, password)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("ユーザーが見つかりません")
	}

	return user, nil
}

// GetUserByUsernameOrEmail はユーザー名またはメールアドレスでユーザーを検索します
func (uc *UserUseCase) GetUserByUsernameOrEmail(username, email string) (*model.User, error) {
	user, err := uc.userRepo.FindByUsernameOrEmail(username, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("ユーザーが見つかりません")
	}

	return user, nil
}

// GetUserByUsernameAndEmail はユーザー名とメールアドレスでユーザーを検索します
func (uc *UserUseCase) GetUserByUsernameAndEmail(username, email string) (*model.User, error) {
	user, err := uc.userRepo.FindByUsernameAndEmail(username, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("ユーザーが見つかりません")
	}

	return user, nil
}

// CreateUser は新しいユーザーを作成します
func (uc *UserUseCase) CreateUser(username, password, email string) (*model.User, error) {
	start := time.Now()
	defer func() {
		fmt.Printf("CreateUser took %v\n", time.Since(start))
	}()
	if username == "" || password == "" || email == "" {
		return nil, errors.New("ユーザー名、パスワード、メールアドレスは必須です")
	}
	hashedPassword, err := utility.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user := model.NewUser(username, hashedPassword, email)

	return uc.userRepo.Create(user)
}

// UpdateUser は既存のユーザーを更新します
func (uc *UserUseCase) UpdateUser(id uint, username, password, email string) (*model.User, error) {
	user, err := uc.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("ユーザーが見つかりません")
	}

	if username != "" {
		user.Username = username
	}
	if password != "" {
		user.Password = password
	}
	if email != "" {
		user.Email = email
	}

	return uc.userRepo.Update(user)
}

// RemoveUser は指定されたIDのユーザーを削除します
func (uc *UserUseCase) RemoveUser(id uint) error {
	user, err := uc.userRepo.FindByID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("ユーザーが見つかりません")
	}
	return uc.userRepo.Remove(id)
}

// Signin はユーザー名とパスワードを検証し、成功時にJWTトークンを返します
func (uc *UserUseCase) Signin(username, password string) (string, error) {
	user, err := uc.GetUserByUsername(username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("ユーザーが見つかりません")
	}
	if !utility.CheckPasswordHash(password, user.Password) {
		return "", errors.New("パスワードが正しくありません")
	}
	token, err := utility.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
