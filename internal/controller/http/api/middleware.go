package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *gophermartHandlers) JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie(h.cfg.Security.CookieTokenName)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		userID, err := h.uc.ParseToken(tokenString)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Set("UserIDCtx", userID)
		c.Next()
	}
}
