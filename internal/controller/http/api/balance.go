package api

import (
	"errors"
	"net/http"

	"github.com/dimk00z/go-musthave-diploma/internal/usecase"
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
	userID := c.GetString("UserIDCtx")
	responseStatus := http.StatusOK
	withdrawals, err := h.uc.GetWithdrawals(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(withdrawals) == 0 {
		responseStatus = http.StatusNoContent
	}
	c.JSON(responseStatus, withdrawals)
}

type WithdrawInput struct {
	OrderNumber string  `json:"order" binding:"required"`
	Sum         float32 `json:"sum" binding:"required"`
}

func (h *gophermartHandlers) postWithdraw(c *gin.Context) {
	var input WithdrawInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.GetString("UserIDCtx")

	err := h.uc.Withdraw(c.Request.Context(), userID, input.OrderNumber, input.Sum)
	if errors.Is(err, usecase.ErrNotEnoughFunds) {
		c.JSON(http.StatusPaymentRequired, gin.H{"error": err.Error()})
		return
	}
	if errors.Is(err, usecase.ErrWrongOrder) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//update logic
	c.JSON(http.StatusOK, input)
}
