package main

import (
	"fmt"
	"goProject/retriever/mock"
	"goProject/retriever/real"
	"time"
)

type Ret interface {
	Get(url string) string
}

//使用者来使用
func download(r Ret) string {
	return r.Get("http://www.baidu.com")
}
func main() {
	var r Ret
	r = mock.Retriever{"this contents"}
	fmt.Printf("%T %v\n", r, r)
	r = reall.RetrieverR{UserAgent: "asas", TimeOut: time.Minute}
	fmt.Printf("%T %v\n", r, r)
	realRet := r.(reall.RetrieverR)
	fmt.Println(realRet.TimeOut)

	//fmt.Println(download(r))
}
