package main

import (
	"fmt"
	//"unicode/utf8"
)

func main() {
	s := "yes一二三四"
	//for i,v := range s{
	//	fmt.Printf("%d %X",i,v)
	//}
	//bytes := []byte(s)
	//fmt.Println(len(bytes))
	//for  len(bytes)>0  {
	//	ch, size := utf8.DecodeRune(bytes)
	//	fmt.Println(ch, size)
	//	bytes = bytes[size:]
	//	fmt.Printf("%c ", ch)
	//}
	fmt.Println([]rune(s))
	for i, ch := range []rune(s){
		fmt.Printf("%d%c ", i, ch)
	}
}
