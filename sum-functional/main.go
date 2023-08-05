package main

import "fmt"

func main() {

	f := sum

	result := executor(2, 3, f)
	fmt.Println("result: ", result)
}

func sum(a, b int) int {
	return a + b
}

func executor(a, b int, f func(int, int) int) int {
	return f(a, b)
}
