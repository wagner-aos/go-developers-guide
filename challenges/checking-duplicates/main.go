package main

import "fmt"

type Developer struct {
	Name string
	Age  int
}

func FilterUnique(developers []Developer) []string {
	uniqueMap := make(map[string]bool)
	var uniqueNames []string

	for _, dev := range developers {
		if _, exists := uniqueMap[dev.Name]; !exists {
			uniqueMap[dev.Name] = true
			uniqueNames = append(uniqueNames, dev.Name)
		}
	}
	return uniqueNames
}

func main() {
	fmt.Println("Filter Unique Challenge")

	devs := []Developer{
		{Name: "Wagner", Age: 49},
		{Name: "John", Age: 43},
		{Name: "Robert", Age: 38},
		{Name: "Robert", Age: 40},
		{Name: "Robert", Age: 23},
	}

	fmt.Println(FilterUnique(devs))
}
