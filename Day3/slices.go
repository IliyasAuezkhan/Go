package main
import(."fmt"; "slices"; "cmp")
func main() {
	array_s := []int{}
	for i := 0; i < 5; i++ {
		var y int
		Scan(&y)
		array_s = append(array_s, y)
	}
	Println(array_s)
	Println(len(array_s))
	array_s = append(array_s, 1, 70, -67)
	Println(array_s)
	Println(cap(array_s))

	num := make([]int, 0, 5) //data type, length, capacity
	num = append(num, 0, 1, 2, 3, 4, 5, 6, 7, 8)
	Println(num)

	my_slice := num[2:4]
	Println(my_slice)
	Println(num[:4])
	Println(num[6:])
	Println(len(my_slice)) //changed
	Println(cap(my_slice)) //saved

	bro := append(my_slice, num...)
	Println(bro)

	les := []int{1, 2, 3}
	guy := make([]int, len(les))
	copy(guy, les)
	Println(guy)

	if slices.Contains(guy, 1) {
		Println("Yes, he contains 1!")
	}

	Println("Index of the 3 is", slices.Index(guy, 3))

	numbers := []int{1, 2, 3, 4, 5, -67}
	hasNegative := slices.ContainsFunc(numbers, func(n int) bool {
		return n < 0
	})
	Println(hasNegative)
	index := slices.IndexFunc(numbers, func(n int) bool {
		return n < 0
	})
	Println(index)  //first occurence of the negative number

	numbers = slices.Insert(numbers, 4, 777)
	Println(numbers)

	numbers = slices.Delete(numbers, 2, 5) //end doesnt included 
	Println(numbers)

	qwe := []int{1,1,2,2,3,3,4,4}
	qwe = slices.Compact(qwe) //delete duplicates also slice must be sorted to work
	Println(qwe)

	freak := []int{5, 7, 9, 10}
	freak = slices.Replace(freak, 1, 3, 67, 67, 67, 67, 67) // end doesnt included
	Println(freak)

	freak2 := []int{7, -9, 0, 12, -777}
	slices.Sort(freak2)
	Println(freak2)
	if slices.IsSorted(freak2) {
		Println("Yes freak 2 is sorted!")
	}
	slices.Reverse(freak2)
	Println(freak2)
	
	type User struct {
		Name string
		Age int
	}
	users := []User {
		{"Bro", 25},
		{"A-train", 89},
		{"Soldier boy", 17},
	}
	slices.SortFunc(users, func(a, b User) int {
		return cmp.Compare(b.Age, a.Age)
	})
	Println(users)
	slices.SortFunc(users, func(a, b User) int {
		return cmp.Compare(b.Name, a.Name)
	})
	Println(users)

	a := []int{1,2,3,4}
	b := slices.Clone(a)
	Println(b)
	if slices.Equal(a, b) {
		Println("They are equal")
	}
}