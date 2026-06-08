package main
import("fmt")
func main() {
	var (
		a int
		b int
	)
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Actually its error")
	} else if a > b {
	 	fmt.Println("a is greater than b")
	} else if a < b {
		fmt.Println("b is greater than a")
	} else {
		fmt.Println("a equals to b")
	}
}