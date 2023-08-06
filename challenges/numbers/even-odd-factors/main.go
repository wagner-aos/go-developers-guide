package main

import (
	"fmt"
	"math"
)

func OddEvenFactors(num int) string {
	factors := CheckFactors(num)
	if len(factors)%2 == 0 {
		return "even"
	}
	return "odd"
}

func CheckFactors(num int) []int {
	factors := []int{1, num}
	limit := int(math.Ceil(float64(num) / 2))
	for i := 2; i <= limit; i++ {
		if num%i == 0 {
			factors = append(factors, i)
		}
	}
	fmt.Printf("Factors: %v\n", factors)
	return factors
}

func main() {
	fmt.Println("Odd or Even Factors")

	numFactors := OddEvenFactors(23)
	fmt.Println(numFactors) // "even"

	numFactors = OddEvenFactors(36)
	fmt.Println(numFactors) // "odd"
}
