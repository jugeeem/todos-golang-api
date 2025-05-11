package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jugeeem/golang-todo.git/app/infrastructure/middleware"
	"github.com/jugeeem/golang-todo.git/app/interface/handler"
)

// SetupRouter はアプリケーションのルーターを設定します
func SetupRouter(
	userHandler *handler.UserHandler,
	authHandler *handler.AuthHandler,
	todoHandler *handler.TodoHandler,
) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://web:3000"}, // フロントエンドのオリジン
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Set-Cookie"},
		AllowCredentials: true,         // Cookieの送受信を許可
		MaxAge:           12 * 60 * 60, // プリフライトリクエストのキャッシュ時間（12時間）
	}))
	public := r.Group("/api/v1")
	{
		public.POST("/token", authHandler.Signin)
		public.POST("/register", userHandler.CreateUser)
	}
	authorized := r.Group("/api/v1")
	authorized.Use(middleware.JWTAuthMiddleware())
	{
		users := authorized.Group("/users")
		{
			users.GET("/", userHandler.GetAllUsers)
			users.GET("/:id", userHandler.GetUserByID)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.RemoveUser)
		}
		todos := authorized.Group("/todos")
		{
			todos.GET("/", todoHandler.GetAllTodos)
			todos.POST("/", todoHandler.CreateTodo)
			todos.GET("/:id", todoHandler.GetTodoByID)
			todos.PUT("/:id", todoHandler.UpdateTodo)
			todos.DELETE("/:id", todoHandler.DeleteTodo)
			todos.GET("/my", todoHandler.GetTodosByUser)
		}
	}

	return r
}
