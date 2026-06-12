package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name" binding:"required,min=2"`
	Email string `json:"email" binding:"required,email"`
	Age   uint   `json:"age" binding:"gte=0,lte=150"`
}

var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@gmail.com", Age: 20},
	{ID: 2, Name: "Bob", Email: "bob@gmail.com", Age: 25},
}

func main() {
	r := gin.Default()
	
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUser)
	r.POST("/users", createUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	r.Run(":8080")
}

func getUsers(c *gin.Context) {
	minAgeStr := c.DefaultQuery("min_age", "0")
	if minAgeStr == "" {
		c.JSON(http.StatusOK, users)
		return
	}
	minAge, err := strconv.Atoi(minAgeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid age",
		})
		return
	}
	var filtered []User
	for _, user := range users {
		if int(user.Age) >= minAge {
			filtered = append(filtered, user)
		}
	}
	c.JSON(http.StatusOK, filtered)
}

func getUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	for _, user := range users {
		if id == int(user.ID) {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"error": "user not found",
	})
}

func createUser(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, existingUser := range users {
		if existingUser.Email == user.Email {
			c.JSON(http.StatusConflict, gin.H{
				"error": "user with this email already exists",
			})
			return
		}
	}

	var maxID uint = 0
	for _, u := range users {
		if u.ID > maxID {
			maxID = u.ID
		}
	}
	user.ID = maxID + 1
	
	users = append(users, user)
	c.JSON(http.StatusCreated, user)
}

func updateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	
	var updatedUser User
	err := c.ShouldBindJSON(&updatedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	
	for _, existingUser := range users {
		if existingUser.Email == updatedUser.Email && int(existingUser.ID) != id {
			c.JSON(http.StatusConflict, gin.H{
				"error": "this email is already taken by another user",
			})
			return
		}
	}

	
	for i, user := range users {
		if int(user.ID) == id {
			users[i].Name = updatedUser.Name
			users[i].Email = updatedUser.Email
			users[i].Age = updatedUser.Age

			
			updatedUser.ID = user.ID
			c.JSON(http.StatusOK, updatedUser)
			return
		}
	}
	
	c.JSON(http.StatusNotFound, gin.H{
		"error": "user not found",
	})
}

func deleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	for i, user := range users {
		if int(user.ID) == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "user deleted",
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"error": "user not found",
	})
}