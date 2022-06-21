package api

import (
	"net/http"

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
	message := "I'm postOrders"
	c.String(http.StatusOK, message)
}
