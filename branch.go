package main

import (
	"io/ioutil"
	"fmt"
	"strconv"
	"os"
	"bufio"
	"reflect"
	"runtime"
	"math"
)

func main() {
	const filename="abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil{
		fmt.Println(err)
	}else{
		fmt.Printf("%s\n",contents)
	}

	fmt.Println(grade(60))
	sum := 0
	for i:=1;i<=100;i++{
		sum +=i
	}
	fmt.Println(sum, conver(12))
	conver(5)
	printFile("abc.txt")
	fmt.Println(divc(13,3))

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(sumD(1,2,3,4,5))
	a,b := 3,4
	a,b = swap(a,b)
	fmt.Println(a, b)
}

func swap( a, b int)(int, int) {
	return b,a
}

func pow( a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
func apply( op func(int, int) int, a, b int) int  {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling %s with args"+"(%d, %d)", opName, a, b)
	return op(a, b)
}

func sumD(numbers ...int) int  {
	s := 0
	for i := range numbers{
		s += numbers[i]
	}
	return s
}

func divc( a, b int)(int, int)  {
	return a/b, a%b
}

func printFile(filename string)  {
	file, err := os.Open(filename)
	if(err != nil){
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func conver(n int) string  {
	result := ""
	for ; n>0; n/=2{
		lsb := n%2
		result = strconv.Itoa(lsb) + result
	}
	return result
}
func grade(score int) string  {
	g := ""
	switch  {
	case score<60:
		g="f"
	case score<90:
		g="G"
	default:
		panic(fmt.Sprintf("wrong %d", score))
	}
	return g
}

func eval(a,b int , op string) int  {
	var result int
	switch op {//switch 默认break
		case "+":
			result = a + b
		case "-":
			result = a - b
		case "*":
			result = a * b
		default:
			panic("undefind:"+op)
	}
	return result;

}