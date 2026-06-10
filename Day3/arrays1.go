package main
import(."fmt")
func main() {
	array := []int{0, 1, 2, 3, 4}
	Println(array)
	array[0] = 10
	Println(array)
	Println(len(array)) //length of the array
	for i, value := range array {
		Println(i, value)
	}
	for i := 0; i < len(array); i++ {
		Println(i, array[i])
	}


	array1 := [5]int{1, 2}
	Println(array1)
	

	array2 := [6]int{1:90, 4:-12} // specify by index other ones will remain by default zero
	Println(array2)

	array3 := [][]int{
		{1,2},
		{3,5},
		{6, 90},
	}
	Println(array3)
	Println(len(array3)) //only first value by x axis if we consider is as a table
	Println(array3[0])
	Println(array3[1][1])
	for i := 0; i < len(array3); i++ {
		for j := 0; j < len(array3[i]); j++ {
			Println(array3[i][j])
		}
	}
}