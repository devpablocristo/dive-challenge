package ltp

type LTP struct {
	Pair   string  `json:"pair"`
	Amount float64 `json:"amount"`
}

type InMemDB map[string]*LTP
