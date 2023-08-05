package main

import (
	"fmt"
	"time"
)

func main() {
	print := func(message string) {
		fmt.Println(message)
	}

	go print("hello go routine")

	time.Sleep(1 * time.Second)
}

func print() {
	fmt.Println("Hello World")
}
