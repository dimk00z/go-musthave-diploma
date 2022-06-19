package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h gophermartHandlers) getOrders(c *gin.Context) {
	message := "I'm getOrders"
	c.String(http.StatusOK, message)
}

func (h *gophermartHandlers) postOrders(c *gin.Context) {
	message := "I'm postOrders"
	c.String(http.StatusOK, message)
}
