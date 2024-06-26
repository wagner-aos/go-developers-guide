package main

import (
	"fmt"
	"sort"
)

// Flight - a struct that
// contains information about flights
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

type byPrice []Flight

func (p byPrice) Len() int {
	return len(p)
}

func (p byPrice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p byPrice) Less(i, j int) bool {
	return p[i].Price < p[j].Price
}

// SortByPrice sorts flights from highest to lowest
func SortByPrice(flights []Flight) []Flight {
	// implement
	sort.Sort(byPrice(flights))

	return flights
}

func printFlights(flights []Flight) {
	for _, flight := range flights {
		fmt.Printf("Origin: %s, Destination: %s, Price: %d \n", flight.Origin, flight.Destination, flight.Price)
	}
}

func main() {
	// an empty slice of flights
	var flights []Flight
	flights = []Flight{
		{Price: 30},
		{Price: 20},
		{Price: 50},
		{Price: 1000},
	}

	sortedList := SortByPrice(flights)
	printFlights(sortedList)
}
