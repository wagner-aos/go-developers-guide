package main

import "fmt"

func main() {

	var array = []int{3, 10, 20, 30, 2, 45, 40}
	calculateChecksum(array) //result [60 85]

	var array2 = []int{4, 15, 20, 30, 10, 3, 45, 40, 30}
	calculateChecksum(array2) //result [75 115]

	var array3 = []int{4, 15, 20, 30, 10, 4, 45, 40, 30, 15}
	calculateChecksum(array3) //result [75 130]

}

func calculateChecksum(fileContent []int) []int {

	var checksumArray []int
	var sum = 0
	var headerValue int

	for i := 0; i < len(fileContent); i++ {
		//fmt.Println("i: ", i)

		if headerValue == 0 {
			headerValue = fileContent[i]
		} else {
			headerValue += fileContent[i] + 1
		}
		//fmt.Println("headerValue: ", headerValue)

		for j := i + 1; j <= headerValue; j++ {
			//fmt.Println("Somando: ", fileContent[j])
			sum += fileContent[j]
			//fmt.Println("sum: ", sum)
			//fmt.Println("j: ", j)
			if j == headerValue {
				i = j
				//headerValue += fileContent[j+1]
				//fmt.Println("BREAK")
				//fmt.Println("j: ", j)
				//fmt.Println("headerValue: ", headerValue)
				break
			}
		}

		//i = headerValue

		checksumArray = append(checksumArray, sum)
		sum = 0
	}

	fmt.Printf("array: %v - checkSumArray: %v \n", fileContent, checksumArray)
	return checksumArray
}
