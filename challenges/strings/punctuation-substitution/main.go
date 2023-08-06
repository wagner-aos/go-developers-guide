package main

import (
	"fmt"
	"strings"
)

/*
given: "Yesterday was Saturday. Today is Sunday! and Tomorrow is Monday!!!"
wait:  "Yesterday was Saturday! Today is Sunday!! and Tomorrow is Monday!!!!"
*/

func changeCharacters(text string) string {

	slice := []string{}
	for _, v := range text {
		slice = append(slice, string(v))
	}

	result := []string{}

	isExclamation := false
	for i := 0; i < len(slice); i++ {
		if slice[i] == "." { // every ocurrency of "." is gonna be changed to "!"
			result = append(result, "!")
		} else if slice[i] == "!" && !isExclamation { //isExclamation controls if the last character added was a "!"
			result = append(result, slice[i], "!") //Append an aditional "!" to the 'result' slice
			isExclamation = true                   //change the last character added to "!", so it means if the next character is a "!", it won't be added!
		} else {
			result = append(result, slice[i]) //Adds all character
			//if the current character is not a "!", it sets isExclamation to false,
			//thus, if the next character is a "!", then it will be added to 'result' slice
			if slice[i] != "!" {
				isExclamation = false
			}
		}
	}

	return strings.Join(result, "")
}

func main() {

	text := "Yesterday was Saturday. Today is Sunday! and Tomorrow is Monday!!!"

	resultText := changeCharacters(text)

	fmt.Println(resultText)

}
