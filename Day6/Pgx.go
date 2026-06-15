package main

import (
	"context" 
	"errors"  
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"          
	"github.com/jackc/pgx/v5/pgxpool"
)


type User struct {
	ID        int
	Name      string
	Age       int
	Created_at time.Time
}

func main() {
	ctx := context.Background()
	connStr := "postgres://postgres:170719@localhost:5432/Go?sslmode=disable"
	
	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatalf("Couldnt create pool: %v", err)
	}
	defer pool.Close()

	err = pool.Ping(ctx) 
	if err != nil {
		log.Fatalf("Couldnt connect: %v", err)
	}
	fmt.Println("We have connected!")

	create_table(pool)

	id1 := insertUser(pool, "Bro", 12)
	id2 := insertUser(pool, "Aksha", 27)

	getUserByID(pool, id2)

	updateUserAge(pool, id1, 99)

	getAllUsers(pool)

	deleteUser(pool, id2)
	fmt.Println("\nAfter delete:")
	getAllUsers(pool)
}


func create_table(pool *pgxpool.Pool) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	query := `CREATE TABLE IF NOT EXISTS users (
				id SERIAL PRIMARY KEY,
				name VARCHAR(100) NOT NULL,
				age INT,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := pool.Exec(ctx, query)
	if err != nil {
		log.Fatalf("Couldnt create the table: %v", err)
	}
	fmt.Println("Successefully created the table!!!")
}


func insertUser(pool *pgxpool.Pool, name string, age int) int {
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var insertedID int
	query := `INSERT INTO users (name, age) VALUES($1, $2) RETURNING id`
	err := pool.QueryRow(ctx, query, name, age).Scan(&insertedID)
	if err != nil {
		log.Fatalf("Error of inserting of user: %v", err)
	}
	
	fmt.Printf("User %s inserted with ID: %d\n", name, insertedID)
	return insertedID
}

func getUserByID(pool *pgxpool.Pool, id int) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var u User
	query := `SELECT id, name, age, created_at FROM users WHERE id = $1`
	err := pool.QueryRow(ctx, query, id).Scan(&u.ID, &u.Name, &u.Age, &u.Created_at)
	if err != nil {
		
		if errors.Is(err, pgx.ErrNoRows) {
			fmt.Printf("User with this id: %d not found\n", id)
			return
		}
		log.Fatalf("ERROR of searching user: %v", err)
	}
	fmt.Printf("Found user: ID: %d | Name: %s | Age: %d | Date: %s\n", u.ID, u.Name, u.Age, u.Created_at)
}

func getAllUsers(pool *pgxpool.Pool) {
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	query := `SELECT id, name, age, created_at FROM users`
	rows, err := pool.Query(ctx, query)
	if err != nil {
		log.Fatalf("Error of getting list: %v", err)
	}
	defer rows.Close()
	fmt.Println("List of every User in the data base: ")
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Age, &u.Created_at)
		if err != nil {
			log.Fatalf("Error of reading lines: %v", err)
		}
		fmt.Printf("ID: %d | Name: %s | Age: %d | Created_at: %s\n", u.ID, u.Name, u.Age, u.Created_at)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalf("Error after reading lines: %v", err)
	}
}

func updateUserAge(pool *pgxpool.Pool, id int, newage int) {
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	query := `UPDATE users SET age = $1 WHERE id = $2`
	commandtag, err := pool.Exec(ctx, query, newage, id)
	if err != nil {
		log.Fatalf("Error of Updating: %v", err)
	}

	if commandtag.RowsAffected() == 0 {
		fmt.Printf("No one updated, user with ID: %d doesnt exist\n", id)
	} else {
		fmt.Printf("Age of user with id %d updated to new age: %d\n", id, newage)
	}
}

func deleteUser(pool *pgxpool.Pool, id int) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	query := `DELETE FROM users WHERE id = $1`
	
	_, err := pool.Exec(ctx, query, id)
	if err != nil {
		log.Fatalf("Error of deleting: %v", err)
	}
	
	fmt.Printf("User with id %d deleted\n", id)
}

