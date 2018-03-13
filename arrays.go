package main

import "fmt"

func main() {
	//var arr1 [5] int
	//arr2 := [3]int{1,3,5}
	//arr3 := [...]int{2,4,6,8,10}
	arr := [...]int{0,1,2,3,4,5,6,7}
	s := arr[2:8]
	fmt.Println(s);
	fmt.Println(arr[:6])
	fmt.Println(arr[2:])
	fmt.Println(arr[:])
	printArray(arr[:])
	updateSlice(s)
	fmt.Println(append(s, 10))
	fmt.Println(s)
	testSlice()
}

func updateSlice(s[]int)  {
	s[0] = 100
}

func printArray( arr []int)  {
	arr[0] = 100
	for i,v:= range  arr{
		fmt.Println(i, v)
	}
}

func testSlice()  {
	var s[]int
	for i:=0;i<100;i++{
		s = append(s,2*i+1)
	}
	fmt.Println(s)
	s0 := []int{2,3,4,5}
	s1 := make([]int, 16)
	copy(s1,s0)
	s2 := make([]int, 10, 32)
	fmt.Println(s1,cap(s1),s2,cap(s2))
	s4 := append(s1[:3], s1[4:]... )
	fmt.Println(s4)

}
