package presenter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Presenter struct{}

func NewPresenter() *Presenter {
	return &Presenter{}
}

func (p *Presenter) Present(c *gin.Context, ltps []ltp.) {
	var ltpResponse []gin.H
	for _, ltp := range ltps {
		ltpResponse = append(ltpResponse, gin.H{
			"pair":   ltp.Pair,
			"amount": ltp.Amount,
		})
	}
	c.JSON(http.StatusOK, gin.H{"ltp": ltpResponse})
}
