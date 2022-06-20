package webapi

import (
	"context"

	"github.com/dimk00z/go-musthave-diploma/config"
	"github.com/dimk00z/go-musthave-diploma/internal/entity"
)

type GopherMartWebAPI struct {
	cfg *config.Config
}

func New(cfg *config.Config) *GopherMartWebAPI {
	return &GopherMartWebAPI{
		cfg: cfg,
	}
}

func (g *GopherMartWebAPI) RegisterUser(
	ctx context.Context,
	userName string,
	password string) (user entity.User, err error) {
	return
}

func (g *GopherMartWebAPI) Login(ctx context.Context,
	userName string,
	password string) (user entity.User, err error) {
	return
}
