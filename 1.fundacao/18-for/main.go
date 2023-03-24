package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}
	numeros := []string{"um", "dois", "três"}
	for key, value := range numeros {
		println(key, value)
	}
	for _, value := range numeros { //Blank identifier -> _
		println(value)
	}
	i := 0
	for i < 10 {
		println(i)
		i++
	}
}
