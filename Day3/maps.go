package main
import(."fmt"; "maps"; "slices"; "strings")
func main() {
	guy := make(map[string]int)
	for i := 0; i < 5; i++ {
		var a string
		var b int
		Scan(&a, &b)
		_, ok := guy[a]
		if ok {
			guy[a] += b
		} else {
			guy[a] = b
		}
	}
	for i, j := range guy {
		Println(i, j)
	}

	bro := make(map[string]int)
	bro["h"] = 4
	bro["g"] = 7
	delete(bro, "h") // delete key
	Println(bro)

	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	keys := make([]string, 0, len(a))
	for k := range a {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, k := range keys {
		Printf("%v : %v, ", k, a[k])
	}
	Println()
	Println(len(a))

	e := maps.Clone(a)
	Println(e)

	gg := make(map[string]int)
	maps.Copy(gg, e)
	Println(gg)
	if maps.Equal(e, gg) {
		Println("They are equal")
	}

	m1 := map[string]float64{"pi":3.14}
	m2 := map[string]float64{"pi":3.1434394}
	close := func(v1, v2 float64) bool {
		return v1 - v2 < 0.01
	}
	Println(maps.EqualFunc(m1, m2, close))

	hey := map[string]int{"Ford":1, "Hi":90, "Soldier boy":500}
	del := func(k string, v int) bool {
		return strings.HasPrefix(k, "H") || v < 25
	}
	maps.DeleteFunc(hey, del)
	Println(hey)

	joker := map[int]string{}
	qwerty := []string{"hey", "sorry", "but"}
	maps.Insert(joker, slices.All(qwerty))
	Println(joker)
	m3 := map[int]string{400:"Lol", 1:"KO"}
	maps.Insert(joker, maps.All(m3))
	Println(joker)

	for k, v := range maps.All(joker) {
		Println(k, v)
	}
	for k := range maps.Keys(joker) {
		Println(k)
	}
	for k := range maps.Values(joker) {
		Println(k)
	}

	clear(a)
	Println(a)
}