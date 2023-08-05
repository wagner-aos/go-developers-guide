package main

import "fmt"

func main() {

	var number int = 481
	var digitSum = getDigitsSum(number)
	fmt.Printf("Digit Sum: %d\n", digitSum)

}

//https://gist.github.com/L7C3/f6fb079c25f9a42afe8d9c0431bebfc7
//The sum of all digits of a number.
func getDigitsSum(x int) int {
	sum := 0
	module := 0
	for ok := true; ok; ok = x != 0 {
		module = x % 10
		//fmt.Printf("module: %d\n", module)
		sum += module
		//fmt.Printf("sum: %d\n", sum)
		x /= 10
		//fmt.Printf("x: %d\n", x)
	}

	fmt.Printf("Sum: %d\n", sum)
	return sum
}
