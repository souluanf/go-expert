package main

import (
	"fmt"
	"github.com/souluanf/go-expert/5.packaging/1-acessando-pacotes-criados/math"
)

func main() {
	m := math.NewMath(1, 2)
	fmt.Println(m.Add())
}
