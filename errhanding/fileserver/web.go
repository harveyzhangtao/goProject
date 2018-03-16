package main

import (
	"fmt"
	"goProject/errhanding/fileserver/filelisting"
	"net/http"
	"os"

	"github.com/gpmgo/gopm/modules/log"
)

func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.Handlist))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	fmt.Println(handler)
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			log.Warn("Error handing request:%s", err)
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
