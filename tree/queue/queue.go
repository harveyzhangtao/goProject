package queue

type Queue []interface{}

//pushes the element into the queue
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v.(int))
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
