package main

import "fmt"

/*Welcome Gophers! In this challenge, you will be tasked with finding out
how many rotations an ordered int slice has undergone and shifted by.
*/

func MinRotations(array []int) int {
	// Implement me :)
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
	fmt.Println("Min Rotation Challenge")

	testArr := []int{15, 18, 2, 3, 6, 12}
	min := MinRotations(testArr) // returns 2
	fmt.Println(min)

	testArr2 := []int{7, 9, 11, 12, 5}
	min2 := MinRotations(testArr2) // return 4
	fmt.Println(min2)

}
