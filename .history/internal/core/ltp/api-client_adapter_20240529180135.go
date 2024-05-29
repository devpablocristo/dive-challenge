package ltp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
		// Obtener la hora actual antes de la solicitud
		startTime := time.Now()

		url := fmt.Sprintf("%s/0/public/Ticker?pair=%s", a.BaseURL, pair)
		resp, err := a.Client.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		// Obtener la hora actual después de recibir la respuesta
		endTime := time.Now()

		var tickerResp KrakenResponse
		if err := json.NewDecoder(resp.Body).Decode(&tickerResp); err != nil {
			return nil, err
		}

		if len(tickerResp.Error) > 0 {
			return nil, fmt.Errorf("api error: %v", tickerResp.Error)
		}

		quote, exists := tickerResp.Result[pair]
		if !exists {
			return nil, fmt.Errorf("pair %s not found in response", pair)
		}

		// Verificar si la respuesta cumple con la precisión temporal requerida
		timeDiff := endTime.Sub(startTime)
		if timeDiff.Minutes() > 1 {
			return nil, fmt.Errorf("data is older than one minute for pair %s", pair)
		}

		amount, err := strconv.ParseFloat(quote.Close[0], 64)
		if err != nil {
			return nil, err
		}

		lastTradePrice := LTP{
			Pair:   pair,
			Amount: amount,
		}
		lastTradedPrices = append(lastTradedPrices, lastTradePrice)
	}

	return lastTradedPrices, nil
}
