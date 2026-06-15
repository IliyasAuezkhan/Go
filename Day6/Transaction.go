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
	ID         int
	Name       string
	Age        int
	Balance    int       
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

	
	id1 := insertUser(pool, "Bro", 12, 1000)  
	id2 := insertUser(pool, "Aksha", 27, 500)  

	fmt.Println("\n--- Баланс до перевода ---")
	getUserByID(pool, id1)
	getUserByID(pool, id2)

	
	fmt.Println("\n--- Запуск транзакции ---")
	transferMoney(pool, id1, id2, 300)

	fmt.Println("\n--- Баланс после перевода ---")
	getUserByID(pool, id1)
	getUserByID(pool, id2)
}

func create_table(pool *pgxpool.Pool) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	
	_, _ = pool.Exec(ctx, "DROP TABLE IF EXISTS users")

	query := `CREATE TABLE users (
				id SERIAL PRIMARY KEY,
				name VARCHAR(100) NOT NULL,
				age INT,
				balance INT NOT NULL DEFAULT 0,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := pool.Exec(ctx, query)
	if err != nil {
		log.Fatalf("Couldnt create the table: %v", err)
	}
	fmt.Println("Successefully created the table!!!")
}

func insertUser(pool *pgxpool.Pool, name string, age int, balance int) int {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var insertedID int
	query := `INSERT INTO users (name, age, balance) VALUES($1, $2, $3) RETURNING id`
	err := pool.QueryRow(ctx, query, name, age, balance).Scan(&insertedID)
	if err != nil {
		log.Fatalf("Error of inserting of user: %v", err)
	}
	fmt.Printf("User %s inserted with ID: %d and Balance: %d\n", name, insertedID, balance)
	return insertedID
}

func getUserByID(pool *pgxpool.Pool, id int) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var u User
	query := `SELECT id, name, age, balance, created_at FROM users WHERE id = $1`
	err := pool.QueryRow(ctx, query, id).Scan(&u.ID, &u.Name, &u.Age, &u.Balance, &u.Created_at)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			fmt.Printf("User with this id: %d not found\n", id)
			return
		}
		log.Fatalf("ERROR of searching user: %v", err)
	}
	fmt.Printf("User: %s (ID: %d) | Age: %d | Balance: %d | Date: %s\n", u.Name, u.ID, u.Age, u.Balance, u.Created_at)
}

func transferMoney(pool *pgxpool.Pool, fromID, toID, amount int) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := pool.Begin(ctx)
	if err != nil {
		log.Fatalf("Couldnt start transition: %v", err)
	}
	defer tx.Rollback(ctx)

	res1, err := tx.Exec(ctx, `UPDATE users SET balance = balance - $1 WHERE id = $2`, amount, fromID)
	if err != nil {
		fmt.Printf("Error minus balance: %v", err)
		return
	}
	if res1.RowsAffected() == 0 {
		fmt.Println("Sender not founded")
		return
	}

	res2, err := tx.Exec(ctx, `UPDATE users SET balance = balance + $1 WHERE id = $2`, amount, toID)
	if err != nil {
		fmt.Printf("Error plus balance: %v", err)
	}
	if res2.RowsAffected() == 0 {
		fmt.Println("Receiver not founded")
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Fatalf("Couldnt commit transaction: %v", err)
		return
	}

	fmt.Printf("Done! Transferred %d from id %d to id %d", amount, fromID, toID)
}