package main
import("fmt")
func main() {
	fmt.Println("From what value should we start? the boundary that included? How fast increase?")
	var (
		a int
		b int
		c int
		sum int = 0
	)
	fmt.Scan(&a, &b, &c)
	for i := a; i <= b; i += c {
		fmt.Println(i)
		sum += i
	}
	fmt.Println(sum)
	fmt.Println("\n")
	for i := 0; i <= 10; i++ {
		if i % 2 == 0 {
			continue
		} 
		fmt.Println(i)
	}
	fmt.Println("\n")
	v := 0
	for i := 10; i >= 0; i--{
		if i % 4 == 0 {
			break
		}
		v += i
	}
	fmt.Println(v)
	for i := range 4 {
		fmt.Println(i)
	}
}