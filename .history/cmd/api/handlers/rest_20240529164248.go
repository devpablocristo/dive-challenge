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

func (h *RestHandler) HelloWorld(c *gin.Context) {
	str := "Hello, World!!! Olá Mundo!!! Hola Mundo!!!"
	c.JSON(http.StatusOK, str)
}

func (h *RestHandler) GetLTP(c *gin.Context) {
	ltp, err := h.ucs.GetLTP(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, pst.ToResponseLTPList(ltp))

}
