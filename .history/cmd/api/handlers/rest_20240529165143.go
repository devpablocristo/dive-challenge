package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	pst "github.com/devpablocristo/dive-challenge/cmd/api/handlers/presenter"
	ucs "github.com/devpablocristo/dive-challenge/internal/core"
)

type RestHandler struct {
	ucs ucs.UseCasePort
}

func NewRestHandler(ucs ucs.UseCasePort) *RestHandler {
	return &RestHandler{ucs: ucs}
}

func (h *RestHandler) GetLTP(c *gin.Context) {
	// Leer los parámetros de consulta "pair"
	pairs := c.QueryArray("pair")

	// Si no se proporcionaron pares, retornar un error
	if len(pairs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'pair' query parameter"})
		return
	}

	// Llamar al caso de uso para obtener los últimos precios negociados
	ltp, err := h.ucs.GetLTP(c.Request.Context(), pairs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convertir los últimos precios negociados al formato de respuesta JSON
	response := pst.ToResponseLTPList(ltp)

	// Retornar los últimos precios negociados en formato JSON
	c.JSON(http.StatusOK, gin.H{"ltp": response})
}
