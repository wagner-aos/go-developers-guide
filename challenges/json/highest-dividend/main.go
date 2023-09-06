package main

import (
	"encoding/json"
	"fmt"
)

/*
In this Go challenge, you are going to be working with JSON data and implementing a function
that takes in a JSON string full of stock quotes and returns the one with the highest dividend return.

You will need to implement the HighestDividend function so that it takes in a JSON string.
You will then need to somehow convert this JSON string into something you can traverse in Go using the encoding/json package.

Finally, you will need to return the string ticker for the stock that has the highest dividend.
Example:
[
  {"ticker": "APPL", "dividend": 0.5},
  {"ticker": "MSFT", "dividend": 0.2}
]

## HighestDividend returns "APPL"
*/

type Stocks struct {
	Stocks []Stocks
}

type Stock struct {
	Ticker   string  `json:"ticker"`
	Dividend float64 `json:"dividend"`
}

// HighestDividend iterates through a JSON string of stocks
// unmarshalls them into a struct and returns the Stock with
// the highest dividend
func HighestDividend(stocksJson string) (string, error) {
	var stockSlice []Stock

	//print(stockSlice)

	//Decode JSON (json to struct)
	if err := json.Unmarshal([]byte(stocksJson), &stockSlice); err != nil {
		panic(err)
	}
	higuest := stockSlice[0]
	//fmt.Println(stockSlice)
	for _, s := range stockSlice {
		if s.Dividend > higuest.Dividend {
			higuest = s
		}
	}

	return higuest.Ticker, nil
}

func main() {
	fmt.Println("Stock Price AI")

	stocks_json := `[
    {"ticker":"APPL","dividend":0.5},
    {"ticker":"GOOG","dividend":0.2},
    {"ticker":"MSFT","dividend":0.3}
  ]`

	highestDividend, _ := HighestDividend(stocks_json)
	fmt.Println(highestDividend)
}
