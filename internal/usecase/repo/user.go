package repo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/dimk00z/go-musthave-diploma/internal/entity"
	"github.com/dimk00z/go-musthave-diploma/internal/usecase"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

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
	user.UserID = userID
	user.Login = userName
	return

}

func (r *GopherMartRepo) GetUser(
	ctx context.Context,
	userName string) (user entity.User, err error) {
	sql, args, err := r.Builder.
		Select("user_id, password, balance, spend").
		From("public.user").
		Where(squirrel.Eq{"login": userName}).
		Limit(1).
		ToSql()
	if err != nil {
		err = fmt.Errorf("GopherMartRepo - GetUser - r.Builder: %w", err)
		return
	}
	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		err = fmt.Errorf("GopherMartRepo - GetUser - r.Pool.Query: %w", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.UserID, &user.Password, &user.Balance.Current, &user.Balance.Spend)
		if err != nil {
			err = fmt.Errorf("GopherMartRepo - GetUser - rows.Scan: %w", err)
			return
		}
	}
	err = rows.Err()
	if err != nil {
		err = fmt.Errorf("GopherMartRepo - GetUser - rows.Err: %w", err)
		return
	}
	user.Login = userName
	return
}
