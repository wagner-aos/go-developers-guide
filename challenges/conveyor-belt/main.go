package main

import "fmt"

/*
This program return the heaviest box from a conveyor belt
- by the index of array
- by the weight of the box
*/

func main() {
	conveyorBelt := []int{10, 100, 36} // The weight of each box

	fmt.Println(returnHeaviestByValue(conveyorBelt[0], conveyorBelt[1], conveyorBelt[2]))

	fmt.Println(returnHeaviestByWeight(conveyorBelt[0], conveyorBelt[1], conveyorBelt[2]))

}

func returnHeaviestByValue(weight0, weight1, weight2 int) int {

	array := [3]int{weight0, weight1, weight2}

	index := 0
	for i, v := range array {
		if v > array[index] {
			index = i
		}
	}
	return index
}

func returnHeaviestByWeight(weight0, weight1, weight2 int) int {
	array := [3]int{weight0, weight1, weight2}

	heaviest := array[0]

	for _, v := range array {
		if v > heaviest {
			heaviest = v
		}
	}

	return heaviest
}
