package usecase

import (
	"errors"

	"github.com/jugeeem/golang-todo.git/app/domain/model"
	"github.com/jugeeem/golang-todo.git/app/domain/repository"
	"github.com/jugeeem/golang-todo.git/app/utility"
)

// AuthUseCase は認証関連のビジネスロジックを提供します
type AuthUseCase struct {
	userRepo repository.UserRepository
}

// NewAuthUseCase は新しいAuthUseCaseのインスタンスを作成します
func NewAuthUseCase(userRepo repository.UserRepository) *AuthUseCase {
	return &AuthUseCase{
		userRepo: userRepo,
	}
}

// Signin はユーザー認証を行い、JWTトークンを返します
func (uc *AuthUseCase) Signin(usernameOrEmail, password string) (string, error) {
	var user *model.User
	var err error
	user, err = uc.userRepo.FindByUsername(usernameOrEmail)
	if err != nil {
		return "", err
	}
	if user == nil {
		user, err = uc.userRepo.FindByEmail(usernameOrEmail)
		if err != nil {
			return "", err
		}
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

// Register は新しいユーザーを登録します
func (uc *AuthUseCase) Register(username, password, email string) (*model.User, error) {
	existingUser, err := uc.userRepo.FindByUsernameOrEmail(username, email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("ユーザー名またはメールアドレスは既に使用されています")
	}
	hashedPassword, err := utility.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user := model.NewUser(username, hashedPassword, email)
	createdUser, err := uc.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

// VerifyToken はJWTトークンを検証し、ユーザーIDとユーザー名を返します
func (uc *AuthUseCase) VerifyToken(tokenString string) (uint, string, error) {
	claims, err := utility.ValidateToken(tokenString)
	if err != nil {
		return 0, "", err
	}

	return claims.UserID, claims.Username, nil
}
