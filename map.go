package main

import "fmt"

func main() {
	//map 是无序的
	//m := map[string]string{
	//	"name":"test",
	//	"name2":"test2",
	//	"name3":"test3",
	//}
	//var m2 map[string]int
	//for k,v := range m{
	//	fmt.Println(k, v)
	//}
 //delete(m,"name2")
	//fmt.Println(m,m2)

	fmt.Println(getString("abcdca一二二三"))
}

func getString(str string) int  {
	last  := make(map[rune]int)
	start := 0
	maxLeng := 0
	c := []rune(str);
	fmt.Println(c)

	for i, ch := range []rune(str){
		lastI,ok  := last[ch];
		if  ok && lastI >= start {
			start = lastI+1
		}
		if i - start + 1 > maxLeng {
			maxLeng = i - start + 1
		}
		last[ch] = i
		//fmt.Println(i,lastI,start, maxLeng)
	}
	//fmt.Println(last)

	return  maxLeng
}
