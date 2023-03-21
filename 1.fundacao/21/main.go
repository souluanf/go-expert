package main

func main() {
	a := 6
	b := 2
	if a > b {
		println("a é maior que b")
	} else {
		println("a é menor que b")
	}

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	default:
		println("não sei")
	}
}
