package main

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Глобальные переменные
var db *gorm.DB
var jwtSecret = []byte("my_super_secret_key")

// Модель для хранения Refresh-токенов в PostgreSQL
type RefreshToken struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"index;not null"`      // К какому юзеру относится
	Token     string    `gorm:"uniqueIndex;not null"` // Сам токен (случайная строка)
	ExpiresAt time.Time `gorm:"not null"`            // До какого времени живет
	CreatedAt time.Time
}

// Структура для JWT (Access)
type MyClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// 1. Функция генерации короткого Access-токена (JWT на 15 минут)
func GenerateAccessToken(userID uint) (string, error) {
	claims := MyClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)), // Короткий срок!
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// 2. Функция генерации длинного Refresh-токена (Случайная строка на 30 дней)
func GenerateAndSaveRefreshToken(userID uint) (string, error) {
	// Генерируем 32 случайных байта и переводим в текст
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	tokenString := hex.EncodeToString(b)

	// Настраиваем время жизни на 30 дней
	refreshToken := RefreshToken{
		UserID:    userID,
		Token:     tokenString,
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour), // 30 дней
	}

	// Сохраняем в PostgreSQL через GORM
	if err := db.Create(&refreshToken).Error; err != nil {
		return "", err
	}

	return tokenString, nil
}

func main() {
	// Подключение к твоей базе данных
	dsn := "host=localhost user=postgres password=170719 dbname=bro port=5432 sslmode=disable TimeZone=Asia/Almaty"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB: " + err.Error())
	}

	// Мигрируем новую таблицу refresh_tokens
	db.AutoMigrate(&RefreshToken{})
	fmt.Println("Migration successful!")

	r := gin.Default()

	// ЭНДПОИНТ ЛОГИНА: Выдает сразу ДВЕ штуки
	r.POST("/login", func(c *gin.Context) {
		userID := uint(52) // Представим, что зашел юзер с ID 52

		// Генерируем Access (JWT)
		accessToken, err := GenerateAccessToken(userID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create access token"})
			return
		}

		// Генерируем и пишем в базу Refresh
		refreshToken, err := GenerateAndSaveRefreshToken(userID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create refresh token"})
			return
		}

		// Отдаем оба токена клиенту
		c.JSON(http.StatusOK, gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
	})

	// ЭНДПОИНТ ОБНОВЛЕНИЯ: Принимает старый Refresh, выдает новые токены
	r.POST("/refresh", func(c *gin.Context) {
		// Описываем структуру того, что ждем от фронтенда
		var input struct {
			RefreshToken string `json:"refresh_token" binding:"required"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token is required"})
			return
		}

		var storedToken RefreshToken
		// Ищем токен в базе данных PostgreSQL и проверяем, что он не просрочен
		result := db.Where("token = ? AND expires_at > ?", input.RefreshToken, time.Now()).First(&storedToken)
		
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
			return
		}

		// Если токен валидный — удаляем старый из базы (чтобы им нельзя было воспользоваться дважды)
		db.Delete(&storedToken)

		// Генерируем новую пару для этого же пользователя
		newAccess, _ := GenerateAccessToken(storedToken.UserID)
		newRefresh, _ := GenerateAndSaveRefreshToken(storedToken.UserID)

		c.JSON(http.StatusOK, gin.H{
			"access_token":  newAccess,
			"refresh_token": newRefresh,
		})
	})

	r.Run(":8080")
}
