package main

import (
	"fmt"
	"math"
)

func DiffSquares(n, m int) int {
	x := math.Pow(float64(n), 2)
	y := math.Pow(float64(m), 2)
	return int(x) - int(y)
}

func main() {
	fmt.Println("Calculate The Difference of Squares")
	result := DiffSquares(5, 4)
	fmt.Println(result)
}
