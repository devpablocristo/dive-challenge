package ltp

type LTPdto struct {
	Pair         string   `json:"pair"` // Currency pair
	Ask          []string `json:"a"`    // Ask price
	Bid          []string `json:"b"`    // Bid price
	Last         []string `json:"c"`    // Last trade price
	Volume       []string `json:"v"`    // Volume
	Weighted     []string `json:"p"`    // Weighted price
	Transactions []int    `json:"t"`    // Number of transactions
	Lowest       []string `json:"l"`    // Lowest price
	Highest      []string `json:"h"`    // Highest price
	Opening      string   `json:"o"`    // Opening price
}
