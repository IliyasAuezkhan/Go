package main
import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	appPort := os.Getenv("APP_PORT")

	if appPort == "" {
		appPort = "8080"
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Error of pool: %v", err)
	}
	defer pool.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := pool.Ping(r.Context())
		fmt.Fprintf(w, "Status of db: %v", err == nil)
	})

	log.Printf("Запуск на порту %s", appPort)
	log.Fatal(http.ListenAndServe(":"+appPort, nil))
}