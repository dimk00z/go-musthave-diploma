package api

import (
	"net/http"

	"github.com/dimk00z/go-musthave-diploma/pkg/logger"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine,
	l logger.Interface) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// POST /api/user/register — регистрация пользователя;
	// POST /api/user/login — аутентификация пользователя;

	// POST /api/user/orders — загрузка пользователем номера заказа для расчёта;
	// GET /api/user/orders — получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях;

	// GET /api/user/balance — получение текущего баланса счёта баллов лояльности пользователя;
	// POST /api/user/balance/withdraw — запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа;
	// GET /api/user/balance/withdrawals — получение информации о выводе средств с накопительного счёта пользователем.
	// Routers
	api := handler.Group("/api")
	userAPI := api.Group("/user")
	{
		userAPI.POST("/register", func(c *gin.Context) {
			message := "test"
			c.String(http.StatusOK, message)
		})
		userAPI.POST("/login", func(c *gin.Context) {
			message := "test"
			c.String(http.StatusOK, message)
		})
	}
	ordersAPI := userAPI.Group("/orders")
	{
		ordersAPI.GET("/", func(c *gin.Context) {
			message := "test"
			c.String(http.StatusOK, message)
		})
		ordersAPI.POST("/", func(c *gin.Context) {
			message := "test"
			c.String(http.StatusOK, message)
		})
	}
	balanceAPI := userAPI.Group("/balance")
	{
		balanceAPI.GET("/", func(c *gin.Context) {
			message := "test"
			c.String(http.StatusOK, message)
		})
		balanceAPI.GET("/withdrawals", func(c *gin.Context) {
			message := "test"
			c.String(http.StatusOK, message)
		})
		balanceAPI.POST("/withdraw", func(c *gin.Context) {
			message := "test"
			c.String(http.StatusOK, message)
		})
	}
}
