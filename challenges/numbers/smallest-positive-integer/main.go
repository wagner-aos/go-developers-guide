package main

import (
	"fmt"
)

/*

You are given an array of integers. Return the smallest positive integer that is not present in the array.
The array may contain duplicate entries.
For example, the input [3, 4, -1, 1] should return 2 because it is the smallest positive integer that doesn't exist in the array.

Your solution should run in linear time and use constant space.
Here's your starting point:

def first_missing_positive(nums):
  # Fill this in.

print first_missing_positive([3, 4, -1, 1])
*/

// func firstMissingPositive(array []int) int {
// 	//sort.Ints(array)
// 	fmt.Printf("array: %v", array)

// 	var smallestPositiveNumber = 1
// 	//[-1,1,3,4]

// 	for i := 1; i < 100; i++ {
// 		exists := false
// 		for _, v := range array {
// 			if v == i {
// 				exists = true
// 				break
// 			}
// 		}
// 		if !exists {
// 			smallestPositiveNumber = i
// 			break
// 		}
// 	}
// 	return smallestPositiveNumber
// }

// // Finding the smallest positive number using a sorted slice
// func firstMissingPositive(numbers []int) int {
// 	var smallestPositiveNumber int

// 	sort.Ints(numbers)
// 	smallestPositiveNumber = numbers[len(numbers)-1]

// 	for _, n := range numbers {
// 		if n > 0 && n < smallestPositiveNumber {
// 			smallestPositiveNumber = n
// 		}
// 	}
// 	return smallestPositiveNumber
// }

// Another approach could be finding the first positive and try to compare other number
// between it and zero.
func firstMissingPositive(numbers []int) int {
	var smallestPositiveNumber int

	//I gonna find the first positive number in the slice
	for _, n := range numbers {
		if n > 0 {
			smallestPositiveNumber = n
			break
		}
	}

	fmt.Printf("The first posivive number which was found in the slice was: %d", smallestPositiveNumber)

	for _, n := range numbers {
		if n > 0 && n < smallestPositiveNumber {
			smallestPositiveNumber = n
		}
	}
	return smallestPositiveNumber
}

func main() {

	//var arrayOfInts = []int{4, -1, 2, 5, 1, -8, 7}
	//var arrayOfInts = []int{4, -1, 3, 5, -8, 7}
	var arrayOfInts = []int{-7, 9, -1, 3, 5, -8, 7}
	result := firstMissingPositive(arrayOfInts)

	//fmt.Printf("Array: %v\n", arrayOfInts)
	fmt.Printf("\nSmallest Positive Number is: %d\n", result)
}
