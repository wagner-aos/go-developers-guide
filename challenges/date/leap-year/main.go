package main

import "fmt"

func CheckLeapYear(year int) bool {
	// If a year is multiple of 400,
	// then it is a leap year
	if year%400 == 0 {
		return true
	}

	// Else If a year is muliplt of 100,
	// then it is not a leap year
	if year%100 == 0 {
		return false
	}

	// Else If a year is muliplt of 4,
	// then it is a leap year
	if year%4 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println("Check Leap Year")

	year := 1988
	fmt.Println(CheckLeapYear(year))
	year = 1992
	fmt.Println(CheckLeapYear(year))
	year = 1996
	fmt.Println(CheckLeapYear(year))
	year = 2000
	fmt.Println(CheckLeapYear(year))
	year = 2016
	fmt.Println(CheckLeapYear(year))
	year = 2020
	fmt.Println(CheckLeapYear(year))
	year = 2023
	fmt.Println(CheckLeapYear(year))
}
