package webapi

import (
	"github.com/dimk00z/go-musthave-diploma/config"
)

type GopherMartWebAPI struct {
	cfg *config.Config
}

func New(cfg *config.Config) *GopherMartWebAPI {
	return &GopherMartWebAPI{
		cfg: cfg,
	}
}
