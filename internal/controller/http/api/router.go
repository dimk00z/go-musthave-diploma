package api

import (
	"github.com/dimk00z/go-musthave-diploma/internal/usecase"
	"github.com/dimk00z/go-musthave-diploma/pkg/logger"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine,
	l logger.Interface, uc usecase.IGopherMart) {

	router.Use(gzip.Gzip(gzip.DefaultCompression))
	// Swagger
	// POST /api/user/register — регистрация пользователя;
	// POST /api/user/login — аутентификация пользователя;

	// POST /api/user/orders — загрузка пользователем номера заказа для расчёта;
	// GET /api/user/orders — получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях;

	// GET /api/user/balance — получение текущего баланса счёта баллов лояльности пользователя;
	// POST /api/user/balance/withdraw — запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа;
	// GET /api/user/balance/withdrawals — получение информации о выводе средств с накопительного счёта пользователем.
	// Routers
	apiRouter := router.Group("/api")
	{
		newGopherMartRoutes(apiRouter, uc, l)
	}
}
