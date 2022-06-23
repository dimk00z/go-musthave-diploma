package repo

import (
	"context"
	"fmt"
	"log"

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

	log.Println(sql, args)

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		err = fmt.Errorf("GopherMartRepo - GetOrder - r.Pool.Query: %w", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&balance.Current, &balance.Spend)
		if err != nil {
			err = fmt.Errorf("GopherMartRepo - GetUser - rows.Scan: %w", err)
			return
		}
	}
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
	_, err = r.Pool.Query(context.Background(), sql, args...)
	return

}

func (r *GopherMartRepo) SaveWithdraw(ctx context.Context, userID string, sum int) (err error) {
	//TODO add Withdraw logic

	return
}
func (r *GopherMartRepo) GetWithdrawals(ctx context.Context, userID string) (withdrawals []entity.Withdrawal, err error) {
	//TODO add Withdraw logic

	return
}
