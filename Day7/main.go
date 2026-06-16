package main
import (
	"context"
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)
type Note struct {
	Id int `json:"id"`
	Title string `json:"title" binding:"required,max=100"`
	Content string `json:"content" binding:"required,min=1"`
	Created_at time.Time `json:"created_at"`
}
var db *pgxpool.Pool

func main() {
	connStr := "postgres://postgres:170719@localhost:5432/Practice?sslmode=disable"
	var err error
	db, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Couldnt connect to the database: %v", err)
	}
	defer db.Close()
	r := gin.Default()

	r.POST("/notes", createNote)
	r.GET("/notes", getAllNotes)
	r.GET("/notes/:id", getNoteByID)
	r.PUT("/notes/:id", updateNote)
	r.DELETE("/notes/:id", deleteNote)

	r.Run(":8080")
}

func createNote(c *gin.Context) {
	var input Note
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}
	query := `INSERT INTO notes (title, content) VALUES ($1, $2) RETURNING id, created_at`
	err = db.QueryRow(context.Background(), query, input.Title, input.Content).Scan(&input.Id, &input.Created_at)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during the creation of the new note"})
		return
	}
	c.JSON(http.StatusCreated, input)
}

func getAllNotes(c *gin.Context) {
	query := `SELECT id, title, content, created_at FROM notes ORDER BY id DESC`
	rows, err := db.Query(context.Background(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during the process of getting data"})
		return
	}
	defer rows.Close()
	notes := []Note{}
	for rows.Next() {
		var n Note
		err = rows.Scan(&n.Id, &n.Title, &n.Content, &n.Created_at)
		if err != nil {
			continue
		}
		notes = append(notes, n)
	}
	c.JSON(http.StatusOK, notes)
}

func getNoteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id is not right"})
		return
	}
	var n Note
	query := `SELECT id, title, content, created_at FROM notes WHERE id = $1`
	err = db.QueryRow(context.Background(), query, id).Scan(&n.Id, &n.Title, &n.Content, &n.Created_at)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "The note didnt found with this id"})
		return
	}
	c.JSON(http.StatusOK, n)
}

func updateNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not correct id"})
		return
	}
	var input Note
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation error": err.Error()})
		return
	}
	query := `UPDATE notes SET title = $1, content = $2 WHERE id = $3 RETURNING created_at`
	err = db.QueryRow(context.Background(), query, input.Title, input.Content, id).Scan(&input.Created_at)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error of updating"})
		return
	}
	input.Id = id
	c.JSON(http.StatusOK, input)
}

func deleteNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not correct id"})
		return
	}
	query := `DELETE FROM notes WHERE id = $1`
	cmdtag, err := db.Exec(context.Background(), query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error of deleting"})
		return
	}
	if cmdtag.RowsAffected() == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note with this id doesnt exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
}