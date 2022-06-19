package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h gophermartHandlers) getBalance(c *gin.Context) {
	message := "I'm getBalance"
	c.String(http.StatusOK, message)
}

func (h gophermartHandlers) getWithdrawals(c *gin.Context) {
	message := "I'm getWithdrawals"
	c.String(http.StatusOK, message)
}

func (h *gophermartHandlers) postWithdraw(c *gin.Context) {
	message := "I'm getWithdrawals"
	c.String(http.StatusOK, message)
}
