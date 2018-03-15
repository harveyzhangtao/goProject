package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	//a := adder()
	//fmt.Println(a)
	//fmt.Println(a(5))
	//fmt.Println(a(6))

	//for i := 0; i < 10; i++ {
	//	fmt.Println(a(i))
	//}

	f := fibona()
	PrintFileContents(f)
}

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		fmt.Println(v)
		sum += v
		return sum
	}
}

func fibona() IntGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func test(a int, b int) (int, int) {
	return b, a
}

type IntGen func() int

func (g IntGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 1000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func PrintFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
