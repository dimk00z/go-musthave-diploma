package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/dimk00z/go-musthave-diploma/internal/entity"
)

func (r *GopherMartRepo) SaveWithdraw(
	ctx context.Context,
	userID string,
	orderNumber string,
	sum float32,
	withdrawalID string) (err error) {

	processedAt := time.Now()
	sql, args, err := r.Builder.
		Insert("public.withdrawal").
		Columns("withdrawal_id, user_id, order_number, processed_at, withdraw_sum").
		Values(withdrawalID, userID, orderNumber, processedAt, sum).
		ToSql()
	if err != nil {
		err = fmt.Errorf("GopherMartRepo - SaveWithdraw - r.Builder: %w", err)
		return
	}
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		err = fmt.Errorf("GopherMartRepo - SaveWithdraw - r.Pool.Exec: %w", err)
		return
	}
	return
}
func (r *GopherMartRepo) GetWithdrawals(ctx context.Context, userID string) (withdrawals []entity.Withdrawal, err error) {
	sql, args, err := r.Builder.
		Select("order_number, processed_at, withdraw_sum").
		From("public.withdrawal").
		Where(squirrel.Eq{"user_id": userID}).
		OrderBy("processed_at").ToSql()
	if err != nil {
		err = fmt.Errorf("GopherMartRepo - GetWithdrawals - r.Builder: %w", err)
		return
	}
	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		err = fmt.Errorf("GopherMartRepo - GetWithdrawals - r.Pool.Query: %w", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		e := entity.Withdrawal{}
		err = rows.Scan(&e.OrderNumber, &e.ProcessedAt, &e.Sum)
		if err != nil {
			return nil, fmt.Errorf("GopherMartRepo - GetWithdrawals - rows.Scan: %w", err)
		}
		withdrawals = append(withdrawals, e)
	}
	return
}
