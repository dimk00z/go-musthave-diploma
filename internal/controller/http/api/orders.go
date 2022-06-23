package api

import (
	"errors"
	"net/http"

	"github.com/ShiraazMoollatjie/goluhn"
	"github.com/dimk00z/go-musthave-diploma/internal/usecase"
	"github.com/gin-gonic/gin"
)

func (h gophermartHandlers) getOrders(c *gin.Context) {

	userID := c.GetString("UserIDCtx")
	orders, err := h.uc.GetOrders(c.Request.Context(), userID)
	responseStatus := http.StatusOK
	if err != nil {
		if errors.Is(err, usecase.ErrNoOrderFound) {
			responseStatus = http.StatusNoContent
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(responseStatus, orders)
}

func (h *gophermartHandlers) postOrder(c *gin.Context) {
	contentType := c.Request.Header.Get("Content-Type")
	if contentType != "text/plain" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong header"})
		return
	}
	userID := c.GetString("UserIDCtx")
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orderNumber := string(body)
	err = goluhn.Validate(orderNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = goluhn.Validate(orderNumber)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	order, err := h.uc.NewOrder(c.Request.Context(), userID, orderNumber)
	if errors.Is(err, usecase.ErrOrderAlreadyGot) {
		c.JSON(http.StatusOK, order)
		return
	}
	if errors.Is(err, usecase.ErrOrderGotByDifferentUser) {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.l.Debug(order)
	c.JSON(http.StatusAccepted, order)
}

type OrderURI struct {
	OrderNumber string `json:"order" uri:"order"`
}

// func (h gophermartHandlers) getOrder(c *gin.Context) {
// 	orderURI := OrderURI{}
// 	if err := c.BindUri(&orderURI); err != nil {
// 		c.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}
// 	// ostrconv.Atoi(s)
// 	userID := c.GetString("UserIDCtx")
// 	order, err := h.uc.GetOrder(c.Request.Context(), orderURI.OrderNumber, userID)
// 	if errors.Is(err, )
// 	c.JSON(http.StatusOK, order)

// }
