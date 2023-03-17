package main

func sumComPonteiro(a, b *int) int {
	*a = 50
	return *a + *b
}

func sumSemPonteiro(a, b int) int {
	a = 50
	return a + b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 10
	soma := sumComPonteiro(&minhaVar1, &minhaVar2)
	println(minhaVar1)
	println(soma)
}
