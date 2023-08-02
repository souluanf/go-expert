package main

func main() {
	var minhaVar interface{} = "Luan Fernandes"
	println(minhaVar.(string)) // afirmando para sistema que é uma String
	resultsdo, ok := minhaVar.(int)
	if ok {
		println(resultsdo)
	} else {
		println("Não é um inteiro")
	}
}
