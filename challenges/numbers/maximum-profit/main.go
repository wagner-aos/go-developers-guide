package main

import (
	"fmt"
)

/*
You are given an array. Each element represents the price of a stock on that particular day.
Calculate and return the maximum profit you can make from buying and selling that stock only once.

For example: [9, 11, 8, 5, 7, 10]

Here, the optimal trade is to buy when the price is 5, and sell when it is 10,
so the return value should be 5 (profit = 10 - 5 = 5).

Here's your starting point:

def buy_and_sell(arr):

	#Fill this in.

print buy_and_sell([9, 11, 8, 5, 7, 10])
# 5
*/
func main() {

	//stocks := []int{9, 11, 8, 5, 7, 10}
	stocks := []int{9, 11, 8, 5, 7, 12, 11, 14}
	profit := buyAndSell(stocks)

	fmt.Printf("The biggest profit is: %d\n", profit)
}

func buyAndSell(stocks []int) int {
	//I gonna declare to variables here, one to smallest price of stock,
	//and another to the biggest one.
	var smallestPrice = stocks[0]
	var biggestPrice int

	// I gonna iterate the array of stocks to find the smallest price.
	for _, v := range stocks {
		if v < smallestPrice {
			smallestPrice = v
		}
	}

	//Now I have the smallest price of the stock, I gonna find the biggest value of the day.
	for i := 0; i < len(stocks); i++ {
		if stocks[i] > smallestPrice && stocks[i] > biggestPrice {
			biggestPrice = stocks[i]
		}
	}

	fmt.Printf("smallestPrice: %d\n", smallestPrice)
	fmt.Printf("biggestPrice: %d\n", biggestPrice)

	//Here I return the profit
	return biggestPrice - smallestPrice

}
