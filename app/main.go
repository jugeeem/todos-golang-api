package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/jugeeem/golang-todo.git/app/infrastructure"
	"github.com/jugeeem/golang-todo.git/app/infrastructure/persistence"
	"github.com/jugeeem/golang-todo.git/app/interface/handler"
	"github.com/jugeeem/golang-todo.git/app/interface/router"
	"github.com/jugeeem/golang-todo.git/app/usecase"
	"github.com/jugeeem/golang-todo.git/app/utility"
)

func main() {
	dbConfig := infrastructure.NewDBConfigFromEnv()
	gormDB, err := infrastructure.ConnectDBWithRetry(dbConfig, 5, 3*time.Second)
	if err != nil {
		log.Fatalf("データベース接続エラー: %v", err)
	}
	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatalf("sql.DB取得エラー: %v", err)
	}
	defer sqlDB.Close()
	userRepo := persistence.NewUserRepository(gormDB)
	todoRepo := persistence.NewTodoRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepo)
	authUseCase := usecase.NewAuthUseCase(userRepo)
	todoUseCase := usecase.NewTodoUseCase(todoRepo)
	userHandler := handler.NewUserHandler(userUseCase)
	authHandler := handler.NewAuthHandler(authUseCase)
	todoHandler := handler.NewTodoHandler(todoUseCase)
	router := router.SetupRouter(userHandler, authHandler, todoHandler)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	port := utility.GetEnv("PORT", "8080")
	log.Printf("サーバーを起動しています: :%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("サーバー起動エラー: %v", err)
	}
}
