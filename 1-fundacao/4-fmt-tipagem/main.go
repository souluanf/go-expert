package main

import "fmt"

const a = "Hello World"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Luan"
	e float64 = 40.0
	f ID      = 10
)

func main() {
	fmt.Printf("O tipo de E é %T\n", f)
}
