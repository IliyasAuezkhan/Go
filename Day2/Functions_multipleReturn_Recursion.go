package main
import("fmt")
func need(a int, b string) (result1 int, result2 string) {
	result1 += a * a
	result2 = "Hello "+ b
	return result1, result2
}
func factorial(a float64) (result float64) {
	if a > 0 {
		result = a * factorial(a - 1)
	} else {
		result = 1
	}
	return result
}
func variadic_sum(nums ...int) (total int){
	total = 0
	for _, i := range nums {
		total += i
	}
	return total
}
func closure(limit int) func() (int, bool) {
	current := 0
	return func() (int, bool) {
		current += 2
		if current > limit {
			return 0, false
		}
		return current, true
	}
}
func main() {
	a, b := need(5, "Bro")
	fmt.Println(a, b)
	_, c := need(6, "Hugie")
	fmt.Println(c)
	var e float64
	fmt.Scan(&e)
	fmt.Println(factorial(e))

	var array []int
	fmt.Println("Whats the length of the array?")
	var l int
	fmt.Scan(&l)
	for i := 0; i < l; i++ {
		var element int
		fmt.Scan(&element)
		array = append(array, element)
	}
	fmt.Println(array)
	fmt.Println(variadic_sum(1, 2, 3, 4))
	fmt.Println(variadic_sum(array...))

	var limit int
	fmt.Scan(&limit)
	generator := closure(limit)
	for {
		val, ok := generator()
		if !ok {
			break
		}
		fmt.Print(val, " ")
	}
}