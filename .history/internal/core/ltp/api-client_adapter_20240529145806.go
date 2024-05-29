package ltp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ApiClient struct {
	BaseURL string
	Client  *http.Client
}

func NewExternalAPI(baseURL string) APIClientPort {
	return &ApiClient{
		BaseURL: baseURL,
		Client:  &http.Client{Timeout: 10 * time.Second},
	}
}

func (a *ApiClient) GetKrakenLTPs(ctx context.Context, pairs []string) ([]LTP, error) {
	var lastTradedPrices []LTP

	for _, pair := range pairs {
		url := fmt.Sprintf("%s/0/public/Ticker?pair=%s", a.BaseURL, pair)
		resp, err := a.Client.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var tickerResp struct {
			Error  []string `json:"error"`
			Ticker map[string]struct {
				C []float64 `json:"c"`
			} `json:"result"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&tickerResp); err != nil {
			return nil, err
		}

		if len(tickerResp.Error) > 0 {
			return nil, fmt.Errorf("api error: %v", tickerResp.Error)
		}

		lastTradePrice := LTP{
			Pair:   pair,
			Amount: tickerResp.Ticker[pair].C[0],
		}
		lastTradedPrices = append(lastTradedPrices, lastTradePrice)
	}

	return lastTradedPrices, nil
}
