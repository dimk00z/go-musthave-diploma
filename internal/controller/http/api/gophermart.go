package api

import (
	"github.com/dimk00z/go-musthave-diploma/config"
	"github.com/dimk00z/go-musthave-diploma/internal/usecase"
	"github.com/dimk00z/go-musthave-diploma/pkg/logger"
	"github.com/gin-gonic/gin"
)

type gophermartHandlers struct {
	uc  usecase.IGopherMart
	l   logger.Interface
	cfg *config.Config
}

// POST /api/user/register — регистрация пользователя;
// POST /api/user/login — аутентификация пользователя;

// POST /api/user/orders — загрузка пользователем номера заказа для расчёта;
// GET /api/user/orders — получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях;

// GET /api/user/balance — получение текущего баланса счёта баллов лояльности пользователя;
// POST /api/user/balance/withdraw — запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа;
// GET /api/user/balance/withdrawals — получение информации о выводе средств с накопительного счёта пользователем.
// Routers
func newGopherMartRoutes(api *gin.RouterGroup, uc usecase.IGopherMart, l logger.Interface, cfg *config.Config) {
	handlers := &gophermartHandlers{uc, l, cfg}

	userAPI := api.Group("/user")
	{
		userAPI.POST("/register", handlers.userRegister)
		userAPI.POST("/login", handlers.userLogin)
	}
	ordersAPI := userAPI.Group("/orders")
	ordersAPI.Use(handlers.JwtAuthMiddleware())
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
