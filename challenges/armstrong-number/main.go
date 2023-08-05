package main

import (
	"fmt"
	"math"
)

/*
Armstrong Numbers - An Armstrong number is a 3-digit number
such that each of the sum of the cubes of its digits equal the number itself:
*/

type MyInt int

func (nptr *MyInt) isArmstrongNumber() bool {
	n := *nptr
	if n > 999 && n < 100 {
		return false
	}

	n1 := float64(n / 100)
	n2 := float64((n % 100) / 10)
	n3 := float64(n % 10)

	fmt.Printf("n1 %f\n", n1)
	fmt.Printf("n2 %f\n", n2)
	fmt.Printf("n3 %f\n", n3)

	calculated := math.Pow(n1, 3) + math.Pow(n2, 3) + math.Pow(n3, 3)
	return float64(n) == calculated
}

func main() {

	var number MyInt = 371
	fmt.Printf("\n%d is Armostrong Number: %v\n\n", number, number.isArmstrongNumber())

}
