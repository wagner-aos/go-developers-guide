package main

import (
	"errors"
	"fmt"
)

//We have a collection of Flights, the fields are Origin of the flight, its destination and the price of the flight
//this must be very pricey, jokes aside.
//1. the challenge is to get the min and max price, right?

// Here I define de struct Flight,
// and put the fields Origin and Destination which are the string type.
// and the Price field is the type int.
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

// Now I define the funcion getMinMax, it receive the parameter which is a slice of Flights,
// and I have put multiple returns which are int, int, error if any type of error occurs.
func getMinMax(flights []Flight) (int, int, error) {
	var err error // I defined an error var here.

	if len(flights) == 0 {
		err = errors.New("ERROR")
		return 0, 0, err
	}

	//I gonna create two variables here, one for minValue, and another for maxValue:
	//I gonna initialize both variables with the first element of the slice.
	var minValue, maxValue = flights[0].Price, flights[0].Price

	for _, f := range flights {
		if f.Price < minValue {
			minValue = f.Price
		}
		if f.Price > maxValue {
			maxValue = f.Price
		}
	}
	return minValue, maxValue, err
}

// In order to test it here, I gonna define a main method
func main() {

	//I gonna create a variable flights which is a slice of Flights.
	var flights []Flight
	//I gonna assign some flights here, and I just going to set the prices, cause I only need it for comparision.
	flights = []Flight{
		{Price: 30},
		{Price: 20},
		{Price: 5},
		{Price: 2000},
		{Price: 50},
		{Price: 1000},
	}

	min, max, error := getMinMax(flights)

	fmt.Println("Getting the Minimum and Maximum Flight Prices")
	fmt.Printf("\nMin: %d \nMax: %d \nERROR: %v \n", min, max, error)
}
