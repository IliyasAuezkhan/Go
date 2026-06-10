package main
import(."fmt")

type Person struct {
	name string
	age int
	job string
	salary float64
}

type area struct {
	width int
	length int
}

func (p *Person) UpdateAge(newage int) {
	p.age = newage
}

func (a *area) calculate() int {
	return a.width * a.length
}

func main() {
	bro := Person{"John", 32, "Engineer", 32000}
	Println(bro.name, bro.age, bro.job, bro.salary)

	bro.UpdateAge(45)
	Println(bro.age)

	block := area{5, 10}
	Println(block.calculate())
}