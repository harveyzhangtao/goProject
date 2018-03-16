package main

import (
	"fmt"
	"net/http"
	"os"

	"goProject/errhanding/fileserver/filelisting"
	"log"
)

func tryRevover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("error code:", err)
		} else {
			panic(r)
		}
	}()
	//panic(errors.New("this is en error"))
	//b := 0
	//a := 5 / b
	//fmt.Println(a)
	panic(112)
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.Handlist))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
	//tryRevover()
}

type userError interface {
	error
	Message() string
}

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	fmt.Println(handler)
	return func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic : %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}

		}()

		err := handler(writer, request)
		if err != nil {
			log.Printf("Error handing request:%s", err)

			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}

}
