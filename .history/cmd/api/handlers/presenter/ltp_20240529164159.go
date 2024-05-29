package presenter

import (
	"github.com/devpablocristo/dive-challenge/internal/core/ltp"
)

type responseLTP struct {
	Pair   string  `json:"pair"`
	Amount float64 `json:"amount"`
}

func ToResponseLTP(l *ltp.LTP) responseLTP {
	return responseLTP{
		Pair:   l.Pair,
		Amount: l.Amount,
	}
}

func ToResponseLTPList(ltps []ltp.LTP) []responseLTP {
	responseLTPs := make([]responseLTP, len(ltps))
	for i, ltp := range ltps {
		responseLTPs[i] = ToResponseLTP(&ltp)
	}
	return responseLTPs
}
