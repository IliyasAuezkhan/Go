package main
import(
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)//10 is a default cost
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "Cool"
	hashed, err := HashPassword(password)
	if err != nil {
		panic("Error of hashing: " + err.Error())
	}
	fmt.Println("Hashed:", hashed)
	fmt.Println("Clean password:", password)
	//password = "123"
	if CheckPasswordHash(password, hashed) {
		fmt.Println("Доступ разрешен")
	} else {
		fmt.Println("Incorrect password")
	}
}