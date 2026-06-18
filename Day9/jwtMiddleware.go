package main
import (
	"fmt"
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)
var jwtSecret = []byte("my_super_secret_key")
type MyClaims struct {
	UserID uint `json:"user_id"`
	Email string `json:"email"`
	jwt.RegisteredClaims //time of life for token
}

func GenerateJWT(userID uint, email string) (string, error) {
	claims := MyClaims{
		UserID: userID,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenstring, nil
}

func ValidateJWT(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*MyClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeaader:= c.GetHeader("Authorization")
		if authHeaader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeaader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		claims, err := ValidateJWT(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		token, err := GenerateJWT(52, "iliyas@gmail.com")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Token error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	protected := r.Group("/api")
	protected.Use(AuthMiddleware())

	protected.GET("/notes", func(c *gin.Context) {
		userID, _ := c.Get("userID")
		c.JSON(http.StatusOK, gin.H{
			"message": "Successful access to private data",
			"user_id": userID,
		})
	})
	r.Run(":8080")
}