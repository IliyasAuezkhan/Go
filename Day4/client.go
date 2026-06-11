package main
import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)
const url = "http://localhost:8080/user"

func sendAndPrint(req *http.Request) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Eror of sending request:", err)
	}
	defer resp.Body.Close()
	fmt.Println("Статус ответа:", resp.Status)
	fmt.Print("Тело ответа: ")
	io.Copy(os.Stdout, resp.Body)
	fmt.Println("\n------------------------------")
}

func main() {
	fmt.Println("=== 1. Отправка POST запроса (Создание) ===")
	bodyPost := bytes.NewBuffer([]byte(`{"id": "10", "name": "Tom"}`))
	reqPost, _ := http.NewRequest("POST", url, bodyPost)
	reqPost.Header.Set("Content-Type", "application/json")
	sendAndPrint(reqPost)

	fmt.Println("=== 2. Отправка GET запроса (Получение) ===")
	reqGet, _ := http.NewRequest("GET", url+"?id=10", nil)
	sendAndPrint(reqGet)

	fmt.Println("=== 3. Отправка PUT запроса (Обновление) ===")
	bodyPut := bytes.NewBuffer([]byte(`{"name": "Tomas"}`))
	reqPut, _ := http.NewRequest("PUT", url+"?id=10", bodyPut)
	reqPut.Header.Set("Content-Type", "application/json")
	sendAndPrint(reqPut)

	fmt.Println("=== 4. Отправка DELETE запроса (Удаление) ===")
	reqDelete, _ := http.NewRequest("DELETE", url+"?id=10", nil)
	sendAndPrint(reqDelete)

	fmt.Println("=== 5. Повторный GET запрос после удаления ===")
	reqCheck, _ := http.NewRequest("GET", url+"?id=10", nil)
	sendAndPrint(reqCheck)
}