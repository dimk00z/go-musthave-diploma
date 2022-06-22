package api

import (
	"errors"
	"net/http"
	"strconv"

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
	message := "I'm getWithdrawals"
	c.String(http.StatusOK, message)
}

type WithdrawInput struct {
	OrderNumber string `json:"order" binding:"required"`
	Sum         int    `json:"sum" binding:"required"`
}

func (h *gophermartHandlers) postWithdraw(c *gin.Context) {
	var input WithdrawInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.GetString("UserIDCtx")
	orderNumber, err := strconv.Atoi(input.OrderNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.uc.Withdraw(c.Request.Context(), userID, orderNumber)
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
	c.JSON(http.StatusOK, input)
}
