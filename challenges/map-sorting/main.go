package main

//https://www.geeksforgeeks.org/how-to-sort-golang-map-by-keys-or-values/

import (
	"fmt"
	"sort"
)

var basket = map[string]int{
	"orange":     5,
	"apple":      7,
	"mango":      3,
	"strawberry": 9,
}

func main() {

	//sortByKeys()
	sortByValues()

}

func sortByKeys() {
	keys := make([]string, 0, len(basket))

	for k := range basket {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, basket[k])
	}
}

func sortByValues() {
	keys := make([]string, 0, len(basket))

	for k := range basket {
		keys = append(keys, k)
	}

	fmt.Println(basket)
	fmt.Println(keys)

	sort.SliceStable(keys, func(i, j int) bool {
		return basket[keys[i]] < basket[keys[j]]
	})

	fmt.Println(keys)
}
