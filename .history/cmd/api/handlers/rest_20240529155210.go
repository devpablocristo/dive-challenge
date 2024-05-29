package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ucs "github.com/devpablocristo/dive-challenge/internal/core"
)

type RestHandler struct {
	ucs ucs.UseCasePort
}

func NewRestHandler(ucs ucs.UseCasePort) *RestHandler {
	return &RestHandler{ucs: ucs}
}

func (h *RestHandler) HelloWorld(c *gin.Context) {
	str := "Hello, World!!! Olá Mundo!!! Hola Mundo!!!"
	c.JSON(http.StatusOK, str)
}

func (h *RestHandler) GetLTP(c *gin.Context) {

	// // Simulación de datos de ejemplo
	// prices := map[string]float64{
	// 	"BTC/USD": 52000.12,
	// 	"BTC/CHF": 49000.12,
	// 	"BTC/EUR": 50000.12,
	// }

	// // Formato de la respuesta JSON
	// resp := gin.H{
	// 	"ltp": []gin.H{
	// 		{"pair": "BTC/CHF", "amount": prices["BTC/CHF"]},
	// 		{"pair": "BTC/EUR", "amount": prices["BTC/EUR"]},
	// 		{"pair": "BTC/USD", "amount": prices["BTC/USD"]},
	// 	},
	// }

	// Escribir la respuesta como JSON
	//c.JSON(http.StatusOK, resp)

	ltp, err := h.ucs.GetLTP(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

}
