package main
import (
	"log"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func main() {
	http.HandleFunc("/ping", PingHandler)
	log.Println("Запуск сервера на порту 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}