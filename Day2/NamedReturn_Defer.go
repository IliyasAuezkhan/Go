package main
import("fmt")
func area_peremeter(length int, width int) (area int, p int) {
	area = length * width
	p = 2*length + 2*width
	return
}
func need(a int, b string) (result1 int, result2 string) {
	result1 += a * a
	result2 = "Hello "+ b
	return result1, result2
}
func main() {
	fmt.Println("What the length and width?")
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(area_peremeter(a, b))

	defer fmt.Println("First") // will execute after all fucntions
	defer fmt.Println(need(5, "Bro")) // execution of several defers will execute in order from bottom to up
	fmt.Println("Second")
}