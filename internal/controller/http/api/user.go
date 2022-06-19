package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *gophermartHandlers) userLogin(c *gin.Context) {
	message := "I'm login"
	c.String(http.StatusOK, message)
}

func (h *gophermartHandlers) userRegister(c *gin.Context) {
	message := "I'm register"
	c.String(http.StatusOK, message)
}
