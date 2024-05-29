package ltp

type LTP struct {
	Pair         string   `json:"pair"`
	Ask          []string `json:"a"` // Precio de compra
	Bid          []string `json:"b"` // Precio de venta
	Last         []string `json:"c"` // Último precio de negociación
	Volume       []string `json:"v"` // Volumen
	Weighted     []string `json:"p"` // Precio ponderado
	Transactions []int    `json:"t"` // Número de transacciones
	Lowest       []string `json:"l"` // Precio más bajo
	Highest      []string `json:"h"` // Precio más alto
	Opening      string   `json:"o"` // Precio de apertura
}
