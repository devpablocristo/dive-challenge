package presenter

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devpablocristo/dive-challenge/internal/core"
)

type LTPPresenter struct{}

func NewLTPPresenter() *LTPPresenter {
	return &LTPPresenter{}
}

func (p *LTPPresenter) Present(c *gin.Context, ltps []*core.LTP) {
	var ltpResponse []gin.H
	for _, ltp := range ltps {
		ltpResponse = append(ltpResponse, gin.H{
			"pair":   ltp.Pair,
			"amount": ltp.Amount,
		})
	}
	c.JSON(http.StatusOK, gin.H{"ltp": ltpResponse})
}
