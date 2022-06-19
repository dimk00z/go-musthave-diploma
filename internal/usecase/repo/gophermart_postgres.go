package repo

import (
	"context"

	"github.com/dimk00z/go-musthave-diploma/internal/entity"
	"github.com/dimk00z/go-musthave-diploma/pkg/postgres"
)

type GopherMartRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *GopherMartRepo {
	return &GopherMartRepo{pg}
}

func (r *GopherMartRepo) RegisterUser(
	ctx context.Context,
	userName string,
	password string) (user entity.User, err error) {
	return
}

func (r *GopherMartRepo) GetUser(
	ctx context.Context,
	userName string,
	password string) (user entity.User, err error) {
	return
}
