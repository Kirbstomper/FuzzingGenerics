package main

import (
	"fmt"

	"github.com/kirbstomper/GoGenerics/genericmath"
)

func main() {

	var a int
	var b float64

	a = 1
	b = 2.6

	println(genericmath.Add(a, a))
	println(genericmath.Add(b, b))

	fmt.Println("Welcome to the Generic Jungle")
}
