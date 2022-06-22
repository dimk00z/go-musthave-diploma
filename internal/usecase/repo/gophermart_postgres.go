package repo

import (
	"github.com/dimk00z/go-musthave-diploma/pkg/postgres"
)

type GopherMartRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *GopherMartRepo {
	return &GopherMartRepo{pg}
}
