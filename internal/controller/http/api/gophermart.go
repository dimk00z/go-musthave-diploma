package api

import (
	"github.com/dimk00z/go-musthave-diploma/internal/usecase"
	"github.com/dimk00z/go-musthave-diploma/pkg/logger"
	"github.com/gin-gonic/gin"
)

type gophermartHandlers struct {
	uc usecase.IGopherMart
	l  logger.Interface
}

func newGopherMartRoutes(api *gin.RouterGroup, uc usecase.IGopherMart, l logger.Interface) {
	handlers := &gophermartHandlers{uc, l}

	userAPI := api.Group("/user")
	{
		userAPI.POST("/register", handlers.userRegister)
		userAPI.POST("/login", handlers.userLogin)
	}
	ordersAPI := userAPI.Group("/orders")
	{
		ordersAPI.GET("/", handlers.getOrders)
		ordersAPI.POST("/", handlers.postOrders)
	}
	balanceAPI := userAPI.Group("/balance")
	{
		balanceAPI.GET("/", handlers.getBalance)
		balanceAPI.GET("/withdrawals", handlers.getWithdrawals)
		balanceAPI.POST("/withdraw", handlers.postWithdraw)
	}
}
