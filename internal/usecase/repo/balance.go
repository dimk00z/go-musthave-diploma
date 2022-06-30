package repo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/dimk00z/go-musthave-diploma/internal/entity"
)

func (r *GopherMartRepo) GetBalance(ctx context.Context, userID string) (balance entity.Balance, err error) {
	sql, args, err := r.Builder.
		Select("balance, spend").
		From("public.user").
		Where(squirrel.Eq{"user_id": userID}).
		Limit(1).
		ToSql()

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		err = fmt.Errorf("GopherMartRepo - GetBalance - r.Pool.Query: %w", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&balance.Current, &balance.Spend)
		if err != nil {
			err = fmt.Errorf("GopherMartRepo - GetBalance - rows.Scan: %w", err)
			return
		}
	}
	// err = rows.Err()
	// if err != nil {
	// 	err = fmt.Errorf("GopherMartRepo - GetBalance - rows.Err: %w", err)
	// 	return
	// }
	return
}

func (r *GopherMartRepo) UpdateBalance(ctx context.Context, userID string, balance entity.Balance) (err error) {
	sql, args, err := r.Builder.
		Update("public.user").
		Set("balance", balance.Current).
		Set("spend", balance.Spend).
		Where(squirrel.Eq{"user_id": userID}).
		ToSql()

	if err != nil {
		return fmt.Errorf("GopherMartRepo - UpdateBalance - r.Builder: %w", err)
	}
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("GopherMartRepo - UpdateBalance - r.Pool: %w", err)
	}
	return

}
