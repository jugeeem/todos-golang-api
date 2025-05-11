package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jugeeem/golang-todo.git/app/usecase"
)

// AuthHandler は認証関連のHTTPリクエストを処理します
type AuthHandler struct {
	authUseCase *usecase.AuthUseCase // ポインタ型に変更
}

// NewAuthHandler は新しいAuthHandlerのインスタンスを作成します
func NewAuthHandler(authUseCase *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

// Signin はユーザーのログイン処理を行います
func (h *AuthHandler) Signin(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := h.authUseCase.Signin(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie(
		"token", // key
		token,   // value
		86400,   // expire
		"/",     // path
		"",      // domain
		false,   // secure
		true,    // HTTPOnly
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "ログイン成功",
		"token":   token,
	})
}
