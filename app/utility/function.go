package utility

import (
	"os"
)

// GetEnv は環境変数の値を取得し、設定されていない場合はデフォルト値を返します
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
