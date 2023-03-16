package main

func main() {
	// memória > endereço > valor
	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	b := &a

	println(&a)
	println(a)
	
	println(ponteiro)
	println(*ponteiro)

	println(b)
	println(*b)
}
