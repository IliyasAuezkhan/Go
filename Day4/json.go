package main
import(."fmt";"encoding/json";"log"; "maps")
type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Salary float64 `json:"Salary"`
	Job string `json:"Job"`
}
type Product struct {
	Title string `json:"title"`
	Price float64 `json:"price"`
}

func main() {
	bro := User{"Bro", 12, 123, "Rock"}
	jsonData, err := json.Marshal(bro)
	if err != nil {
		log.Fatal("Error", err)
	}
	Println(string(jsonData))

	guy := []byte(`{"title":"Shampoo", "price":13.292}`)
	var p Product
	err2 := json.Unmarshal(guy, &p)
	if err2 != nil {
		log.Fatal("Error:", err2)
	}
	Println(p)
	Printf("%+v\n", p)

	girl := []byte(`{"home":"det", "Code":123134}`)
	var result map[string]any
	err3 := json.Unmarshal(girl, &result)
	if err3 != nil {
		log.Fatal("Error", err3)
	}
	for i, v := range maps.All(result) {
		Println(i, v)
	}
}