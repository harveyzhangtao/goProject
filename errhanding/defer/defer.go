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
	//	file, err := os.Create(filename)
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	if err != nil {
		//panic(err)
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibona()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}
