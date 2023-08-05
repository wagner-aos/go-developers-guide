package main

import "fmt"

type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func filter(s []student, f func(student) bool) []student {
	var result []student
	for _, v := range s {
		if f(v) == true {
			result = append(result, v)
		}
	}
	return result
}

func praticalUseOfFirstClassFunction() {
	s1 := student{
		firstName: "Wagner",
		lastName:  "AOS",
		grade:     "A",
		country:   "Brazil",
	}

	s2 := student{
		firstName: "John",
		lastName:  "Doe",
		grade:     "B",
		country:   "USA",
	}

	s := []student{s1, s2}

	f := filter(s, func(s student) bool {
		if s.grade == "A" {
			return true
		}
		return false
	})

	fmt.Println(f)
}
