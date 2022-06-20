package repo

import (
	"context"
	"fmt"
	"log"

	"github.com/dimk00z/go-musthave-diploma/internal/entity"
	"github.com/dimk00z/go-musthave-diploma/internal/usecase"
	"github.com/dimk00z/go-musthave-diploma/pkg/postgres"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

type GopherMartRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *GopherMartRepo {
	return &GopherMartRepo{pg}
}

func (r *GopherMartRepo) SaveUser(
	ctx context.Context,
	userID, userName, password string) (user entity.User, err error) {
	sql, args, err := r.Builder.
		Insert("public.user").
		Columns("user_id, login, password").
		Values(userID, userName, password).
		ToSql()
	if err != nil {
		err = fmt.Errorf("GopherMartRepo - SaveUser - r.Builder: %w", err)
		return
	}
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err == nil {
		user.UserId = userID
		user.Login = userName
		return user, nil

	}
	if pqerr, ok := err.(*pgconn.PgError); ok {
		log.Println(pqerr.Code)
		if pgerrcode.IsIntegrityConstraintViolation(pqerr.Code) {
			err = usecase.ErrUserAlreadyExists
			return
		}
	}
	err = fmt.Errorf("GopherMartRepo - SaveUser - r.Builder: %w", err)
	return
}

func (r *GopherMartRepo) GetUser(
	ctx context.Context,
	userName string,
	password string) (user entity.User, err error) {
	return
}
