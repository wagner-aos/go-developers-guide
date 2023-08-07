package main

/*
you are tasked with decoding the secret message.
You will need to implement the DecodeSecret function which will take in a string
that has been encoded using base64 encoding and decode this string.

This decoded string will be the result of a Caesar cipher which has shifted
all of the characters of the string up by 1 place. So you will have to ensure
that when you return the result, it also decodes this cipher.
*/

import (
	"encoding/base64"
	"fmt"
)

func DecodeSecret(message string) string {
	// I will decode base 64 'message' string
	data, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	var secret []rune //in this array of rune I will append every char decode with Ceasar's chiper
	//Traverse all characters of 'data' base64 decoded to decode one by one.
	for _, char := range data {
		fmt.Println(rune(char + 1))
		secret = append(secret, rune(char-1))
	}

	fmt.Println(secret)
	//and finally i will return the array of rune 'secret' to a string decoded!
	return string(secret)
}

func main() {
	fmt.Println("Decode the Secret")

	message := "VEZEU0ZVVFVTSk9I"
	result := DecodeSecret(message)
	fmt.Println(result)

}
