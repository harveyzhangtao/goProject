package main

import (
	"bufio"
	"fmt"
	"goProject/functional/fib"
	"os"
)

func main() {
	//tryDefer()
	writefile("fib.txt")
}

func tryDefer() {
	//defer 先进后出
	defer fmt.Println(1)
	defer fmt.Println(2)
	panic("111")
}

func writefile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibona()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}
