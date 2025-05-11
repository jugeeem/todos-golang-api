package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jugeeem/golang-todo.git/app/infrastructure/migration"
)

func main() {
	// データベースの接続設定
	dbHost := getEnv("DB_HOST", "")
	dbUser := getEnv("DB_USER", "")
	dbPass := getEnv("DB_PASSWORD", "")
	dbName := getEnv("DB_NAME", "")
	dbPort := getEnv("DB_PORT", "")

	// コマンドライン引数の定義
	rollback := flag.Bool("rollback", false, "ロールバックを実行する場合はtrue")
	flag.Parse()

	// データベースURLを環境変数から取得
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			dbUser, dbPass, dbHost, dbPort, dbName)
	}

	// マイグレーションの実行
	var err error
	if *rollback {
		err = migration.RollbackMigration(dbURL)
	} else {
		err = migration.RunMigration(dbURL)
	}

	if err != nil {
		log.Fatalf("Migration error: %v", err)
	}

	fmt.Println("Migration command executed successfully")
}

// getEnv は環境変数の値を取得し、設定されていない場合はデフォルト値を返します
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
