package main

import "fmt"

func FindPairsWithGivenDifference(arr []int, k int) [][]int {
	pairs := make([][]int, 0)

	mapp := map[int]int{}

	if k == 0 {
		return pairs
	}

	for _, x := range arr {
		diff := x - k
		mapp[diff] = x
		print(mapp[diff])
	}

	for _, y := range arr {
		for _, v := range mapp {
			if mapp[y] == v {
				p := []int{mapp[y], y}
				pairs = append(pairs, p)
			}
		}
	}

	return pairs
}

func main() {
	input := []int{0, -1, -2, 2, 1}
	result := FindPairsWithGivenDifference(input, 1)
	fmt.Printf("\n%v\n", result)

}

//PSEUDO CODE:
// # since we don't allow duplicates, no pair can satisfy x - 0 = y
// if k == 0:
// 	return []

// map = {}
// answer = []

// for (x in arr):
// 	map[x - k] = x

// for (y in arr):
// 	if y in map:
// 		answer.push([map[y], y])

// return answer

// import "fmt"

// func FindPairsWithGivenDifference(arr []int, k int) [][]int {
// 	var pairs [][]int

// 	for i := 0; i < len(arr); i++ {
// 		for j := 0; j < len(arr); j++ {
// 			diff := arr[i] - arr[j]
// 			if diff == k {
// 				//fmt.Printf("i: %d - j: %d = %d\n", arr[i], arr[j], diff)
// 				p := make([]int, 0, 2)
// 				p = append(p, arr[i])
// 				p = append(p, arr[j])
// 				fmt.Printf("p: %v\n", p)
// 				pairs = append(pairs, p)
// 			}
// 		}
// 	}

// 	return pairs
// }

// func main() {
// 	input := []int{0, -1, -2, 2, 1}
// 	result := FindPairsWithGivenDifference(input, 1)

// 	fmt.Printf("%v\n", result)
// 	//[[1, 0], [0, -1], [-1, -2], [2, 1]]
// }

// Pairs with Specific Difference
// A naive approach is is to run two loops. The outer loop picks the first element (smaller element)
//and the inner loop looks up for the element picked by the outer loop plus k.
//While this solution is done in O(1) space complexity, its time complexity is O(N^2),
//which isn’t asymptotically optimal.

// We can use a hash map to improve the time complexity to O(N⋅log(N))
//for the worst case and O(N) for the average case. We rely on the fact
//that if x - y = k then x - k = y.

// The first step is to traverse the array, and for each element arr[i],
//we add a key-value pair of (arr[i] - k, arr[i]) to a hash map.
//Once the map is populated, we traverse the array again, and check
//for each element if a match exists in the map.

// Both the first and second steps take O(N⋅log(N)) for the worst case and O(N)
//for the average case. So the overall time complexity is O(N) for the average case.
