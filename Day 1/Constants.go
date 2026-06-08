package main
import("fmt")
const root = 777
func main() {
	const bro string = "hello"
	const a, b = 78, "Cool"
	const freak = true
	fmt.Println(bro, a, b, freak)
	fmt.Print(root)
	// we cannot change constants. If we do it, it will raise error
}