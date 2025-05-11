package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// GetUserID はコンテキストからユーザーIDを取得します
func GetUserID(c *gin.Context) (uint, error) {
	userID, exists := c.Get("userID")
	if !exists {
		return 0, errors.New("ユーザーIDが見つかりません")
	}
	id, ok := userID.(uint)
	if !ok {
		return 0, errors.New("ユーザーIDの型が無効です")
	}

	return id, nil
}
