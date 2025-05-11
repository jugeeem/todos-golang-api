package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jugeeem/golang-todo.git/app/domain/dto"
	"github.com/jugeeem/golang-todo.git/app/infrastructure/middleware"
	"github.com/jugeeem/golang-todo.git/app/usecase"
)

// TodoHandler はTodo関連のHTTPリクエストを処理します
type TodoHandler struct {
	todoUseCase *usecase.TodoUseCase
}

// NewTodoHandler は新しいTodoHandlerのインスタンスを作成します
func NewTodoHandler(todoUseCase *usecase.TodoUseCase) *TodoHandler {
	return &TodoHandler{
		todoUseCase: todoUseCase,
	}
}

// GetAllTodos は全てのTodoタスクを取得するエンドポイント
func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	todos, err := h.todoUseCase.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.ToTodoResponseList(todos))
}

// GetTodoByID は特定のTodoタスクを取得するエンドポイント
func (h *TodoHandler) GetTodoByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}
	todo, err := h.todoUseCase.GetTodoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todoが見つかりません"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// CreateTodo は新しいTodoタスクを作成するエンドポイント
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "認証が必要です"})
		return
	}
	var input struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo, err := h.todoUseCase.CreateTodo(input.Title, input.Description, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToTodoResponse(todo))
}

// UpdateTodo はTodoタスクを更新するエンドポイント
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "認証が必要です"})
		return
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Completed   *bool  `json:"completed"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo, err := h.todoUseCase.UpdateTodo(
		uint(id),
		input.Title,
		input.Description,
		input.Completed,
		userID,
	)
	if err != nil {
		if err.Error() == "このTodoを編集する権限がありません" {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, dto.ToTodoResponse(todo))
}

// GetTodosByUser は現在ログイン中のユーザーのTodoタスクを取得するエンドポイント
func (h *TodoHandler) GetTodosByUser(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "認証が必要です"})
		return
	}
	todos, err := h.todoUseCase.GetTodosByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToTodoResponseList(todos))
}

// DeleteTodo はTodoタスクを削除するエンドポイント
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "認証が必要です"})
		return
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}
	if err := h.todoUseCase.DeleteTodo(uint(id), userID); err != nil {
		if err.Error() == "このTodoを削除する権限がありません" {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todoを削除しました"})
}
