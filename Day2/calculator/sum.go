package calculator
import("fmt")
func Sum(nums ...int) {
	total := 0
	for _, a := range nums {
		total += a
	}
	fmt.Println(total)
}