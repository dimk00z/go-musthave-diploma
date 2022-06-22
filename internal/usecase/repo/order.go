package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/dimk00z/go-musthave-diploma/internal/entity"
	"github.com/dimk00z/go-musthave-diploma/internal/usecase"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

func (r *GopherMartRepo) GetOrder(orderNumber string) (order entity.Order, err error) {

	return
}

func (r *GopherMartRepo) NewOrder(
	ctx context.Context,
	userID, orderID, orderNumber string) (order entity.Order, err error) {
	uploaded_at := time.Now()
	status := "NEW"
	sql, args, err := r.Builder.
		Insert("public.order").
		Columns("user_id, order_number, uploaded_at, order_id, status").
		Values(userID, orderNumber, uploaded_at, orderID, status).
		ToSql()
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
	return entity.Order{OrderID: orderID, Status: status, OrderNumber: orderNumber, ProcessedAt: uploaded_at.Unix()}, err
}
func (r *GopherMartRepo) GetOrders(ctx context.Context, userID string) (orders []entity.Order, err error) {
	sql, args, err := r.Builder.
		Select("order_id, order_number, status, uploaded_at, accrual").
		From("public.order").
		Where(squirrel.Eq{"user_id": userID}).
		OrderBy("uploaded_at").ToSql()
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
		e := entity.Order{}

		err = rows.Scan(&e.OrderID, &e.OrderNumber, &e.ProcessedAt, &e.Accrual)
		if err != nil {
			return nil, fmt.Errorf("TranslationRepo - GetHistory - rows.Scan: %w", err)
		}

		orders = append(orders, e)
	}
	if len(orders) == 0 {
		err = usecase.ErrNoOrderFound
	}
	return
}
