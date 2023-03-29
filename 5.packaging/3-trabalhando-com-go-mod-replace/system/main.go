package main

import (
	"fmt"
	"github.com/souluanf/go-expert/5.packaging/3-trabalhando-com-go-mod-replace/math"
)

func main() {
	m := math.NewMath(1, 2)
	fmt.Println(m.Add())
}
