package utility

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecretKey []byte
var costFactorValue int

func init() {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	costFactorStr := os.Getenv("BCRYPT_COST_FACTOR")
	if secretKey == "" {
		secretKey = "fallback_development_key_do_not_use_in_production"
		fmt.Println("警告: JWT_SECRET_KEYが設定されていません。開発用のキーを使用します。")
	}
	jwtSecretKey = []byte(secretKey)
	if costFactorStr == "" {
		costFactorValue = 12
		fmt.Println("警告: BCRYPT_COST_FACTORが設定されていません。デフォルトのコスト値を使用します。")
	} else {
		cfValue, err := strconv.Atoi(costFactorStr)
		if err != nil {
			fmt.Println("警告: BCRYPT_COST_FACTORの値が無効です。デフォルトの値を使用します。")
			costFactorValue = 12
		} else {
			costFactorValue = cfValue
		}
	}
}

// JWTClaims はJWTのペイロード部分です
type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// hashPassword はパスワードをbcryptでハッシュ化します
func HashPassword(password string) (string, error) {
	costFactor := costFactorValue
	combinedPassword := password + string(jwtSecretKey)
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(combinedPassword), costFactor)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

// comparePassword はパスワードとハッシュを比較します
func VerifyPassword(hashedPassword, password string) error {
	combinedPassword := password + string(jwtSecretKey)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(combinedPassword))
	if err != nil {
		return err
	}

	return nil
}

// CheckPasswordHash はパスワードとハッシュが一致するか検証します
func CheckPasswordHash(password, hash string) bool {
	combinedPassword := password + string(jwtSecretKey)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(combinedPassword))
	return err == nil
}

// GenerateToken はユーザー情報からJWTトークンを生成します
func GenerateToken(userID uint, username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JWTClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "golang-todo-app",
			Subject:   fmt.Sprintf("%d", userID),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken はトークンを検証し、有効であればクレームを返します
func ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecretKey, nil
		},
	)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
