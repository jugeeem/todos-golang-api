package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jugeeem/golang-todo.git/app/domain/dto"
	"github.com/jugeeem/golang-todo.git/app/usecase"
)

// UserHandler はユーザー関連のHTTPリクエストを処理する
type UserHandler struct {
	userUseCase *usecase.UserUseCase // ポインタ型に変更
}

// NewUserHandler はUserHandlerのコンストラクタ
func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

// CreateUser は新しいユーザーを作成する
func (h *UserHandler) CreateUser(c *gin.Context) {
	var input struct {
		Username        string `json:"username" binding:"required"`
		Email           string `json:"email" binding:"required,email"`
		Password        string `json:"password" binding:"required"`
		ConfirmPassword string `json:"confirmPassword" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.Password != input.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "パスワードが一致しません"})
		return
	}
	user, err := h.userUseCase.CreateUser(input.Username, input.Password, input.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToUserResponse(user))
}

// GetUserByID はIDでユーザーを取得する
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}
	user, err := h.userUseCase.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ユーザーが見つかりません"})
		return
	}

	c.JSON(http.StatusOK, dto.ToUserResponse(user))
}

// GetAllUsers は全ユーザーを取得する
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userUseCase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.ToUserResponseList(users))
}

// UpdateUser はユーザー情報を更新する
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email" binding:"email"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.userUseCase.UpdateUser(uint(id), input.Username, input.Password, input.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ユーザーが見つかりません"})
		return
	}

	c.JSON(http.StatusOK, dto.ToUserResponse(user))
}

// RemoveUser はユーザーを削除する
func (h *UserHandler) RemoveUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}
	err = h.userUseCase.RemoveUser(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ユーザーが見つかりません"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
