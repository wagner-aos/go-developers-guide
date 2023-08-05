package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// Flight struct which contains
// the origin, destination and price of a flight
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

// IsSubset checks to see if the first set of
// flights is a subset of the second set of flights.
func IsSubset(first, second []Flight) bool {
	// implement
	set := make(map[Flight]int)
	for _, value := range second {
		set[value] += 1
	}

	fmt.Printf("%v\n", set)
	//time.Sleep(10 * time.Second)

	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	fmt.Printf("%v", set)

	return true
}

func Hash(f Flight) []byte {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(f)
	return b.Bytes()
}

func main() {
	fmt.Println("Sets and Subsets Challenge")
	firstFlights := []Flight{
		{Origin: "GLA", Destination: "CDG", Price: 1000},
		{Origin: "GLA", Destination: "JFK", Price: 5000},
		{Origin: "GLA", Destination: "SNG", Price: 3000},
	}

	secondFlights := []Flight{
		{Origin: "GLA", Destination: "CDG", Price: 1000},
		{Origin: "GLA", Destination: "JFK", Price: 5000},
		{Origin: "GLA", Destination: "SNG", Price: 3000},
		{Origin: "GLA", Destination: "AMS", Price: 500},
	}

	subset := IsSubset(firstFlights, secondFlights)
	fmt.Println(subset)
}
