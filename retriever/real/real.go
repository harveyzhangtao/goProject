package reall

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type RetrieverR struct {
	UserAgent string
	TimeOut   time.Duration
}

//实现者不需要知道实现哪个接口， 只要实现方法就行
func (r RetrieverR) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(result)
}

func (r RetrieverR) Set(url string) string {
	return "SSS"
}
