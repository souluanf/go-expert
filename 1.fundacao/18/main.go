package main

type MyNumber int

type Number interface {
	~int | float64 // ˜ significa que ele considera que podemos usar int ou float64
}

func Sum[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	return a == b
}
func main() {
	m := map[string]int{"João": 200, "Maria": 700, "Pedro": 254}
	m2 := map[string]float64{"João": 30.0, "Maria": 2500.0, "Pedro": 305.0}
	m3 := map[string]MyNumber{"João": 300, "Maria": 400, "Pedro": 500}
	println(Sum(m))
	println(Sum(m2))
	println(Sum(m3))
	println(Compara(10, 10))
}
