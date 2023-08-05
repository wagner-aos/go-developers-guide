package main

import "fmt"

//An algorithm of a function that can return a point of the join of two sequences

//https://gist.github.com/L7C3/f6fb079c25f9a42afe8d9c0431bebfc7
func getDigitsSum(x int) int {
	sum := 0
	for ok := true; ok; ok = x != 0 {
		sum += x % 10
		fmt.Printf("sum: %d\n", sum)
		x /= 10
		fmt.Printf("x: %d\n", x)
	}

	fmt.Printf("Sum: %d\n", sum)
	return sum
}

func findJoinPoint(seq1, seq2 int) int {
	for seq1 != seq2 {
		fmt.Printf("seq1 = %d - seq2 %d \n", seq1, seq2)

		if seq1 < seq2 {
			if seq1 == 0 {
				return -1 // no chance to join, seq1 is 0
			}
			seq1 += getDigitsSum(seq1)
			fmt.Printf("Seq 1 calculated: %d\n", seq1)
		} else if seq2 < seq1 {
			if seq2 == 0 {
				return -2 // no chance to join, seq2 is 0
			}
			seq2 += getDigitsSum(seq2)
			fmt.Printf("Seq 2 calculated: %d\n", seq2)
		}
	}

	return seq1
}

func main() {

	result := findJoinPoint(471, 480)

	fmt.Printf("\nJointing Point: %d\n", result)
}
