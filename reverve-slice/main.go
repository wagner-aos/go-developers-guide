package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rev_slc := []int{}

	for i := range slice {
		rev_slc = append(rev_slc, slice[len(slice)-1-i])
	}

	fmt.Println(rev_slc)
}
