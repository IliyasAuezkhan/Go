package main
import(
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres password=170719 dbname=Go sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping() //Are we connected?
	if err != nil {
		panic(err)
	}
	fmt.Println("Успешное подключение к PostgreSQL!")
	query := `CREATE TABLE IF NOT EXISTS users (
					   id SERIAL PRIMARY KEY,
					   name VARCHAR(100) NOT NULL,
					   age INT,
					   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Couldnt create the table: %v", err)
	}
	fmt.Println("Таблица 'users' успешно создана или уже существует!")
	query2 := `INSERT INTO users (name, age) values ('Iliyas', 12)`
	_, err = db.Exec(query2)
	if err != nil {
		log.Fatalf("Couldnt insert into the table: %v", err)
	}

	query = `SELECT * FROM users`
	rows, err2 := db.Query(query)
	if err2 != nil {
		log.Fatalf("Couldnt get data from the users")
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var age int
		var createdAt string
		err = rows.Scan(&id, &name, &age, &createdAt)
		if err != nil {
			log.Fatalf("error of reading: %v", err)
		}
		fmt.Printf("id: %d | Name: %s | Age: %d | Date: %s\n", id, name, age, createdAt)
	}
	if err = rows.Err(); err != nil {
		log.Fatalf("Ошибка после чтения всех строк: %v", err)
	}
}