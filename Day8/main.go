package main
import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"fmt"
	"errors"
	"time"
)

type User struct {
	ID uint `gorm:"primaryKey;<-:create"`
	Status int `gorm:"comment:Подсказка"`
	TrackingNum string `gorm:"uniqueIndex"`
	Name string `gorm:"not null"`
	Password string `gorm:"column:password_hash"`
	Age int `gorm:"default:18;check:age_range, age >= 0 AND age <= 122"`
	Email string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}

func main() {
	dsn := "host=localhost user=postgres password=170719 dbname=gromP port=5432 sslmode=disable TimeZone=Asia/Almaty"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Couldnt connect to the database: " + err.Error())
	}
	fmt.Println("1) Successfully connected to the database!", db)

	err = db.AutoMigrate(&User{}, &Category{})
	if err != nil {
		panic("Error during the migration: " + err.Error())
	}
	fmt.Println("2) Successfully did migration")

	bro := User {
		Name: "Iliyas",
		TrackingNum: "XYZ",
		Password: "Utyr",
		Age: 18,
		Email: "iliyas@gmail.com",
	}
	result := db.Create(&bro) //insert
	if result.Error != nil {
		fmt.Println("Error:", result.Error)
	} else {
		fmt.Println("User created!", result.RowsAffected)
		fmt.Println("Automatic id:", bro.ID) // automatically can get new id
	}
	
	fmt.Println("\n--- [FIND / FIRST] ---")
	var foundUser User
	resultFind := db.Where("password_hash = ?", "Utyr").First(&foundUser)
	if resultFind.Error != nil {
		if errors.Is(resultFind.Error, gorm.ErrRecordNotFound) {
			fmt.Println("Error: User not found!")
		} else {
			fmt.Println("Database error:",resultFind.Error)
		}
	}
	fmt.Println("Found user:", foundUser)
	var allUsers []User
	db.Find(&allUsers)
	fmt.Println("All users:",allUsers)

	fmt.Println("\n--- SAVE ---")
	foundUser.Name = "Iliyas Updated"
	foundUser.Status = 1
	resultSave := db.Save(&foundUser) //Saving changings
	if resultSave.Error != nil {
		fmt.Println("Save failed:", resultSave.Error)
	} else {
		fmt.Println("User successfully updated and saved! ", foundUser)
	}

	guy := User {
		Name: "Iliyas",
		TrackingNum: "123",
		Password: "Utyr",
		Age: 18,
		Email: "as@gmail.com",
	}
	result = db.Create(&guy) //insert
	if result.Error != nil {
		fmt.Println("Error:", result.Error)
	} else {
		fmt.Println("User created!", result.RowsAffected)
		fmt.Println("Automatic id:", guy.ID) // automatically can get new id
	}

	fmt.Println("\n--- DELETE ---")
	resultDelete := db.Where("tracking_num = ?", "123").Delete(&guy)
	if resultDelete.Error != nil {
		fmt.Println("Delete failed:", resultDelete.Error)
	} else {
		fmt.Println("User deleted:", resultDelete.RowsAffected)
	}
}