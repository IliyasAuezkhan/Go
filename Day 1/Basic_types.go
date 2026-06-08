package main
import "fmt"

func main() {
	var bro int = 5
	fmt.Println(bro)
	var root string = "Go"
	fmt.Println(root)
	var home float32 = 23.94795
	fmt.Println(home)
	var a, b  = 5, true
	fmt.Println(a, b)
	c, d := 23.4, false
	fmt.Println(c, d)
	home = 45.5459
	fmt.Println(home)
	var who bool = false
	fmt.Println(who)
	f := 32
	fmt.Println(f)
	f = 50
	fmt.Println(f)
	var (
		x = "string"
		y = 5
		z = 34.6
	)
	fmt.Println(x, y, z)
}