package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var wait sync.WaitGroup

func main() {

	wait.Add(1)
	toUpperSync("Hello Callbacks", func(v string) {
		fmt.Printf("Callback: %s\n", v)
		wait.Done()
	})
	println("Waiting async response...")
	wait.Wait()
}

func toUpperSync(word string, f func(string)) {
	go func() {
		time.Sleep(1 * time.Second)
		f(strings.ToUpper(word))
	}()
}
