package basic

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 33
	bb = 44
)

func main() {
	euler()
	var a, b int = 3, 4
	fmt.Println(triang(a, b))
	consts()
}

func triang(a int, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

const a, b int = 3, 4

func consts() {
	const (
		cpp = 1 << iota
		java
		python
		golang
	)

	fmt.Println(cpp, java, python, golang)
}

func euler() {
	fmt.Printf("%3.f", cmplx.Exp(1i*math.Pi)+1)
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

func variableType() {
	a, b, c := 3, 4, "aaa"
	fmt.Println(a, b, c)
}
func variable() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b)
	fmt.Printf("%q", s)
}
