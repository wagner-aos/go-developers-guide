package main

import (
	"fmt"
	"sort"
)

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func CheckPermutations(str1, str2 string) bool {
	// if lengths are not equal - false
	if len(str1) != len(str2) {
		return false
	}

	// sort both strings/runes
	var rune1 RuneSlice = []rune(str1)
	var rune2 RuneSlice = []rune(str2)

	sort.Sort(rune1)
	sort.Sort(rune2)

	//fmt.Println(string(rune1[:]))
	//fmt.Println(string(rune2[:]))

	// compare rune1 and rune 2 by indexes
	for i := 0; i < len(rune1); i++ {
		if rune1[i] != rune2[i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("Check Permutations Challenge")

	str1 := "adcme"
	str2 := "medac"

	isPermutation := CheckPermutations(str1, str2)
	fmt.Println(isPermutation)

}
