package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h gophermartHandlers) getBalance(c *gin.Context) {
	userID := c.GetString("UserIDCtx")
	balance, err := h.uc.GetBalance(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, balance)
}

func (h gophermartHandlers) getWithdrawals(c *gin.Context) {
	message := "I'm getWithdrawals"
	c.String(http.StatusOK, message)
}

func (h *gophermartHandlers) postWithdraw(c *gin.Context) {
	message := "I'm getWithdrawals"
	c.String(http.StatusOK, message)
}
