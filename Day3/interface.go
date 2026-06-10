package main
import(."fmt")
type Speaker interface {
	Speak() string
}
type Renamer interface {
	Getname() string
	Setname(string) 
}
type Animal struct {
	name string
}
func (d *Animal) Getname() string {
	return d.name
}
func (a *Animal) Setname(newname string) {
	a.name = newname
}
func Change(r Renamer,name string) {
	r.Setname(name)
}


type Dog struct {
	name string
}
func (b Dog) Speak() string {
	return "Gaf"
}
type Cat struct {
	name string
}
func (b Cat) Speak() string {
	return "Meow"
}
func MakeThemSpeak(s Speaker) {
	Println(s.Speak())
}
func AreYouThere(s Speaker) {
	dog, ok := s.(Dog)
	if ok {
		Println(dog.name)
	}
}

func main() {
	sharik := Dog{"Sharik"}
	mars := Cat{"Mars"}
	MakeThemSpeak(sharik)
	MakeThemSpeak(mars)
	AreYouThere(sharik)

	bro := Animal{"Qw"}
	Println(bro.Getname())
	Change(&bro, "Sharik")
	Println(bro.name)
}