package main
import("testing")
func TestAdd(t *testing.T) {
	num1 := 5
	num2 := 3
	expected := 8
	result := Add(num1, num2)
	if result != expected {
		t.Errorf("Add(%d, %d) сломался: получили %d, а ожидали %d", num1, num2, result, expected)
	}
}