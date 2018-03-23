package main

import (
	"fmt"
	"sync"
)

func main() {
	chanDemo()
	////bufferedChannel()
	//channelClose()
}

//
//func channelClose() {
//	c := make(chan int, 3)
//	go worker(0, c)
//
//	c <- 'a'
//	c <- 'b'
//	c <- 'c'
//	c <- 'd'
//	close(c)
//	time.Sleep(time.Millisecond)
//
//}
//func bufferedChannel() {
//	c := make(chan int, 3)
//	go worker(0, c)
//
//	c <- 'a'
//	c <- 'b'
//	c <- 'c'
//	c <- 'd'
//	time.Sleep(time.Millisecond)
//}

func doWorker(id int, w worker) {
	for n := range w.in {
		fmt.Printf("work %d received %c\n", id, n)
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

func createworker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w)
	return w
}

func chanDemo() {
	var workers [10]worker
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i] = createworker(i, &wg)
	}
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()
}
