package api

import (
	"github.com/dimk00z/go-musthave-diploma/config"
	"github.com/dimk00z/go-musthave-diploma/internal/usecase"
	"github.com/dimk00z/go-musthave-diploma/pkg/logger"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine,
	l logger.Interface, uc usecase.IGopherMart, cfg *config.Config) {

	router.Use(gzip.Gzip(gzip.DefaultCompression))

	apiRouter := router.Group("/api")
	{
		newGopherMartRoutes(apiRouter, uc, l, cfg)
	}
}
