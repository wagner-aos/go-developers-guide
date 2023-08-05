package main

import (
	"errors"
	"fmt"
)

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func getMinMax(flights []Flight) (int, int, error) {
	var err error

	if len(flights) == 0 {
		err = errors.New("ERROR")
		return 0, 0, err
	}

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

func main() {

	var flights []Flight
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
