package presenter

import (
	"github.com/devpablocristo/dive-challenge/internal/core/ltp"
)

type responseLTP struct {
	Pair   string  `json:"pair"`
	Amount float64 `json:"amount"`
}

func LTP(l *ltp.LTP) *responseLTP {
	return &responseLTP{
		Pair:   l.Pair,
		Amount: l.Amount,
	}
}

func LTPList(ltps []ltp.LTP) []responseLTP {
	var responseLTPs []responseLTP
	for _, ltp := range ltps {
		responseLTPs = append(responseLTPs, LTP(&ltp))
	}
	return responseLTPs
}
