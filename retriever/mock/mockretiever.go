package mock

type Retriever struct {
	Contents string
}

func (r Retriever) Get(url string) string {
	return r.Contents
}

func (r Retriever) Set(url string) string {
	return "sss"
}
