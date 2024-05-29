package ltp

type LTP struct {
	Pair   string      `json:"pair"`
	Amount interface{} `json:"amount"`
}

type InMemDB map[string]*LTP
