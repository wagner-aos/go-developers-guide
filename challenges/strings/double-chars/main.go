package main

/*
In this challenge, you are tasked with implementing a function DoubleChars
which will take in a string and then return another string which has every letter in the word doubled.

https://tutorialedge.net/challenges/go/repeating-letters/
*/

import (
	"fmt"
)

func DoubleChars(original string) string {
	endString := make([]rune, 0, len(original)*2)
	for _, c := range original {
		endString = append(endString, c, c)
	}
	return string(endString)
}

////Alternative function
// func DoubleChars(original string) string {
// 	//Initially, we declare an array with double the size of the original string that is being passed as the function parameter.
// 	result := make([]string, len(original)*2)

// 	//I gonna traverse the original string to obtain every character of the string.
// 	for _, v := range original {
// 		//I gonna append every character in a double way til the end of the original string.
// 		result = append(result, string(v), string(v))
// 	}
// 	return strings.Join(result, "")
// }

func main() {
	fmt.Println("Smallest Difference Challenge")

	original := "gophers"
	doubled := DoubleChars(original)
	fmt.Println(doubled) // ggoopphheerrss
}
