package main
import("encoding/json"; "fmt"; "sync"; "net/http")
type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

var (
	users = make(map[string]User)
	usersMu sync.Mutex
)

func userHandler(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")

	switch r.Method {
	case "GET":
		usersMu.Lock()
		user, exists := users[id]
		usersMu.Unlock()
		if !exists || id == "" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "User not found"}`))
			return
		}
		json.NewEncoder(w).Encode(user)

	case "POST":
		var newUser User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil || newUser.ID == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Invalid request body"}`))
			return 
		}	
		usersMu.Lock()
		users[newUser.ID] = newUser
		usersMu.Unlock()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)

	case "PUT":
		var updatedUser User
		err := json.NewDecoder(r.Body).Decode(&updatedUser)
		if err != nil || id == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Invalid data or missing ID"}`))
			return 
		}
		usersMu.Lock()
		_, exists := users[id]
		if !exists {
			usersMu.Unlock()
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "User not found to update"}`))
			return 
		}
		updatedUser.ID = id
		users[id] = updatedUser
		usersMu.Unlock()
		json.NewEncoder(w).Encode(updatedUser)

	case "DELETE":
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Missing id"}`))
			return 
		}

		usersMu.Lock()
		if _, exists := users[id]; !exists {
			usersMu.Unlock()
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "User not found"}`))
			return
		}
		delete(users, id)
		usersMu.Unlock()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "User deleted successfully"}`))
	
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "Method not allowed"}`))
	}
}

func main() {
	http.HandleFunc("/user", userHandler)
	fmt.Println("Сервер запущен на порту :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера...", err)
	}
}