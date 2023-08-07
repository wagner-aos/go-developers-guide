package main

import "fmt"

/*
you will be tasked with finding out how many rotations
an ordered int slice has undergone and shifted by.

It has undergone = foi submetido
*/

func minRotation(array []int) int {

	size := len(array)
	if size <= 1 {
		return 0
	}

	lowIndex := 0
	middleIndex := 0
	highIndex := size - 1

	for array[lowIndex] > array[highIndex] {
		middleIndex = (lowIndex + highIndex) / 2
		if array[middleIndex] > array[highIndex] {
			lowIndex = middleIndex + 1
		} else {
			highIndex = middleIndex
		}
	}

	return lowIndex
}

func main() {

	fmt.Println("Min rotation challenge")

	//testArray := []int{1, 2, 3, 4, 5} // returns 0
	testArray := []int{3, 4, 5, 1, 2} // returns 3
	//testArray := []int{15, 18, 2, 3, 6, 12} // returns 2
	min := minRotation(testArray)
	fmt.Println(min)
}
