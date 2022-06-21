package api

import (
	"net/http"
	"strconv"

	"github.com/dimk00z/go-musthave-diploma/internal/utils"
	"github.com/gin-gonic/gin"
)

func (h gophermartHandlers) getOrders(c *gin.Context) {
	message := "I'm getOrders"

	userID := c.GetString("UserIDCtx")
	if len(userID) > 0 {
		message += "userid " + userID
	}
	c.String(http.StatusOK, message)
}

func (h *gophermartHandlers) postOrders(c *gin.Context) {
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
	orderID, err := strconv.Atoi(string(body))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = utils.LuhnValid(orderID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	// TODO add saving
	c.JSON(http.StatusOK, gin.H{"user": userID, "body": orderID})
}

type OrderURI struct {
	OrderID int `json:"order" uri:"order"`
}

func (h gophermartHandlers) getOrder(c *gin.Context) {
	orderURI := OrderURI{}
	if err := c.BindUri(&orderURI); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// ostrconv.Atoi(s)
	userID := c.GetString("UserIDCtx")

	c.JSON(http.StatusOK, gin.H{"user": userID, "body": orderURI.OrderID})

}
