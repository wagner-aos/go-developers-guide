package main

import (
	"fmt"
	"math"
)

func main() {

	//temperatures := []int{-5, -4, -3, -2, -1, 1, 2, 4, 5}
	temperatures := []float64{7, -10, 13, 8, 4, -7.2, -12, -3.7, 3.5, -9.6, 6.5, -1.7, -6.2, 7}
	fmt.Println(closeToZero(temperatures))

}

func closeToZero(numbers []float64) float64 {

	if len(numbers) == 0 {
		return 0
	}

	closest := numbers[0]

	for _, t := range numbers {

		number := t
		absNumber := math.Abs(number)
		absClosest := math.Abs(closest)
		fmt.Println("ABSNumber: ", absNumber)
		fmt.Println("ABSClosest: ", absClosest)

		fmt.Println("Number: ", t)

		if absNumber < absClosest {
			fmt.Println("<")
			closest = number
		} else if absNumber == absClosest && closest < 0 {
			fmt.Println("<=")
			closest = number
		}

		fmt.Println("closest: ", closest)
		//fmt.Println(t)

	}
	return closest
}

//func closeToZero(numbers []float64) float64 {
// for _, t := range numbers {

// 	fmt.Println("Number: ", t)

// 	if len(numbers) == 0 {
// 		return 0
// 	} else if t > 0 && t <= math.Abs(closest) {
// 		fmt.Println("<=")
// 		closest = t
// 	} else if t < 0 && t < math.Abs(closest) {
// 		fmt.Println("<")
// 		closest = t
// 	}

// 	fmt.Println("closest: ", closest)
// 	//fmt.Println(t)

// }

// return closest
// }
