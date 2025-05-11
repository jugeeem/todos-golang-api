package infrastructure

import (
	"fmt"
	"log"
	"time"

	"github.com/jugeeem/golang-todo.git/app/utility"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBConfig はデータベース接続情報を保持します
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewDBConfigFromEnv は環境変数からDBConfigを作成します
func NewDBConfigFromEnv() *DBConfig {
	return &DBConfig{
		Host:     utility.GetEnv("DB_HOST", "localhost"),
		Port:     utility.GetEnv("DB_PORT", "5432"),
		User:     utility.GetEnv("DB_USER", "postgres"),
		Password: utility.GetEnv("DB_PASSWORD", "password"),
		DBName:   utility.GetEnv("DB_NAME", "todo_db"),
		SSLMode:  utility.GetEnv("DB_SSLMODE", "disable"),
	}
}

// PostgresConnectionString はPostgreSQLへの接続文字列を生成します
func (c *DBConfig) PostgresConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode)
}

// ConnectDB はデータベースに接続します
func ConnectDB(config *DBConfig) (*gorm.DB, error) {
	dsn := config.PostgresConnectionString()
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// ConnectDBWithRetry はリトライロジックを用いてデータベースに接続します
func ConnectDBWithRetry(config *DBConfig, attempts int, sleep time.Duration) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	for i := 0; i < attempts; i++ {
		db, err = ConnectDB(config)
		if err == nil {
			log.Println("データベース接続に成功しました")
			return db, nil
		}
		log.Printf("データベース接続に失敗しました（試行 %d/%d）: %v", i+1, attempts, err)
		if i < attempts-1 {
			log.Printf("%v 秒後に再試行します...", sleep.Seconds())
			time.Sleep(sleep)
		}
	}

	return nil, fmt.Errorf("データベース接続に失敗しました: %w", err)
}
