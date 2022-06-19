package webapi

import (
	"context"

	"github.com/dimk00z/go-musthave-diploma/internal/entity"
)

type GopherMartRepoWebAPI struct {
}

func New() *GopherMartRepoWebAPI {
	return &GopherMartRepoWebAPI{}
}

func (g *GopherMartRepoWebAPI) RegisterUser(
	ctx context.Context,
	userName string,
	password string) (user entity.User, err error) {
	return
}

func (g *GopherMartRepoWebAPI) Login(ctx context.Context,
	userName string,
	password string) (user entity.User, err error) {
	return
}
