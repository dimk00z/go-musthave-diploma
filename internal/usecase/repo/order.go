package repo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/dimk00z/go-musthave-diploma/internal/entity"
	"github.com/dimk00z/go-musthave-diploma/internal/usecase"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

func (r *GopherMartRepo) GetOrder(ctx context.Context, orderNumber int) (order entity.Order, err error) {
	sql, args, err := r.Builder.
		Select("user_id, order_id, status, uploaded_at, accrual").
		From("public.order").
		Where(squirrel.Eq{"order_number": orderNumber}).
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
		e := entity.Order{}

		err = rows.Scan(&e.UserID, &e.OrderID, &e.Status, &e.ProcessedAt, &e.Accrual)
		if err != nil {
			return order, fmt.Errorf("GopherMartRepo - GetOrder - rows.Scan: %w", err)
		}
		e.OrderNumber = orderNumber
		order = e
	}
	return
}

func (r *GopherMartRepo) NewOrder(
	ctx context.Context,
	userID, orderID string, orderNumber int) (order entity.Order, err error) {

	uploadedAt := time.Now()
	status := "NEW"
	sql, args, err := r.Builder.
		Insert("public.order").
		Columns("user_id, order_number, uploaded_at, order_id, status").
		Values(userID, orderNumber, uploadedAt, orderID, status).
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
	return entity.Order{
		OrderID:     orderID,
		Status:      status,
		OrderNumber: orderNumber,
		ProcessedAt: uploadedAt}, err
}

func (r *GopherMartRepo) GetOrders(ctx context.Context, userID string) (orders []entity.Order, err error) {
	sql, args, err := r.Builder.
		Select("order_id, order_number, status, uploaded_at, accrual").
		From("public.order").
		Where(squirrel.Eq{"user_id": userID}).
		OrderBy("uploaded_at").ToSql()
	if err != nil {
		err = fmt.Errorf("GopherMartRepo - GetOrders - r.Builder: %w", err)
		return
	}
	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		err = fmt.Errorf("GopherMartRepo - GetOrders - r.Pool.Query: %w", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		e := entity.Order{}

		err = rows.Scan(&e.OrderID, &e.OrderNumber, &e.Status, &e.ProcessedAt, &e.Accrual)
		if err != nil {
			return nil, fmt.Errorf("GopherMartRepo - GetOrders - rows.Scan: %w", err)
		}
		e.UserID = userID
		orders = append(orders, e)
	}
	if len(orders) == 0 {
		err = usecase.ErrNoOrderFound
	}
	return
}
