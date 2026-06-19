package main
import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"
	"golang.org/x/time/rate"
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Success bool `json:"success"`
	Data interface{} `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level : slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	r := gin.New()
	//Middleware
	r.Use(CORS())
	r.Use(SlogErrorHandlerAndLogger())
	r.Use(RateLimiter())

	r.GET("/api/success", func(c *gin.Context) {
		c.JSON(http.StatusOK, APIResponse{
			Success: true,
			Data: gin.H{"message": " День 10 успешно завершен"},
		})
	})
	r.GET("/api/error", func(c *gin.Context) {
		errMock := errors.New("ошибка подключения к PostgreSQL заметок")
		c.Error(errMock)
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Error: "Error of internal server",
		})
	})

	slog.Info("Сревер успешно запущен на порту 8080")
	r.Run(":8080")
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func SlogErrorHandlerAndLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		//pause
		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		if len(c.Errors) > 0 {
			for _, ginErr := range c.Errors {
				slog.Error("Запрос завершился ошибкой",
					slog.String("method", method),
					slog.String("path", path),
					slog.Int("status", status),
					slog.String("error", ginErr.Error()),
					slog.Duration("latency", latency),
			)
			}
			return
		}
		
		slog.Info("Запрос успешно обработан", 
			slog.String("method", method),
			slog.String("path", path),
			slog.Int("status", status),
			slog.Duration("latency", latency),
		)
	}
}

func RateLimiter() gin.HandlerFunc {
	limiter := rate.NewLimiter(5, 5)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, APIResponse{
				Success: false,
				Error: "Слишком много запросов. Пожалуйста, подождите.",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}