package api

import (
	"github.com/dimk00z/go-musthave-diploma/pkg/logger"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine,
	l logger.Interface) {
	// , t usecase.Translation) {
	// Options
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
	// h := handler.Group("/api")
	{
		// newTranslationRoutes(h, t, l)
	}
}
