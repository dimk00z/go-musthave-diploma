package repo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
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
	if err != nil {
		if pqerr, ok := err.(*pgconn.PgError); ok {
			if pgerrcode.IsIntegrityConstraintViolation(pqerr.Code) {
				err = usecase.ErrUserAlreadyExists
				return
			}
		}
		err = fmt.Errorf("GopherMartRepo - SaveUser - r.Builder: %w", err)
		return
	}
	user.UserId = userID
	user.Login = userName
	err = r.createBalance(ctx, userID)
	return

}

func (r *GopherMartRepo) createBalance(ctx context.Context, userID string) (err error) {
	//TODO Add balance
	if err != nil {
		err = fmt.Errorf("GopherMartRepo - createBalance - r.Builder: %w", err)
	}
	return
}
func (r *GopherMartRepo) GetUser(
	ctx context.Context,
	userName string) (user entity.User, err error) {
	sql, _, err := r.Builder.
		Select("user_id, password").
		From("public.user").
		Where(squirrel.Eq{"login": userName}).
		Limit(1).
		ToSql()
	if err != nil {
		err = fmt.Errorf("GopherMartRepo - GetUser - r.Builder: %w", err)
		return
	}
	rows, err := r.Pool.Query(ctx, sql, userName)
	if err != nil {
		err = fmt.Errorf("GopherMartRepo - GetUser - r.Pool.Query: %w", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.UserId, &user.Password)
		if err != nil {
			err = fmt.Errorf("GopherMartRepo - GetUser - rows.Scan: %w", err)
			return
		}
		break
	}
	user.Login = userName
	return
}
