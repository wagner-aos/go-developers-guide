package main

import (
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex
	value int
}

//You must test race condition of the variable counte using the command line:
// go run --race main.go

func main() {
	counter := Counter{}

	for i := 0; i < 10; i++ {
		go func(i int) {
			counter.Lock() // you must comment here to test race condition
			counter.value++
			defer counter.Unlock() // you must comment here to test race condition
		}(i)
	}

	time.Sleep(time.Second)
	counter.Lock()
	defer counter.Unlock()

	println(counter.value)

}
