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

func (r *GopherMartRepo) GetOrder(ctx context.Context, orderNumber string) (order entity.Order, err error) {
	sql, args, err := r.Builder.
		Select("user_id, order_id, status, uploaded_at, accrual").
		From("public.order").
		Where(squirrel.Eq{"order_number": orderNumber}).
		Limit(1).
		ToSql()
	if err != nil {
		err = fmt.Errorf("GopherMartRepo - GetOrder - r.Builder: %w", err)
		return
	}

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
	userID, orderID, orderNumber string) (order entity.Order, err error) {

	uploadedAt := time.Now()
	status := "NEW"
	sql, args, err := r.Builder.
		Insert("public.order").
		Columns("user_id, order_number, uploaded_at, order_id, status").
		Values(userID, orderNumber, uploadedAt, orderID, status).
		ToSql()
	if err != nil {
		err = fmt.Errorf("GopherMartRepo - NewOrder - r.Builder: %w", err)
		return
	}
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		if pqerr, ok := err.(*pgconn.PgError); ok {
			if pgerrcode.IsIntegrityConstraintViolation(pqerr.Code) {
				err = usecase.ErrOrderAlreadyGot
				return
			}
		}
		err = fmt.Errorf("GopherMartRepo - SaveUser - r.Pool.Exec: %w", err)
		return
	}
	return entity.Order{
		OrderID:     orderID,
		Status:      status,
		OrderNumber: orderNumber,
		ProcessedAt: uploadedAt}, err
}

func (r *GopherMartRepo) GetOrdersForUser(ctx context.Context, userID string) (orders []entity.Order, err error) {
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

func (r *GopherMartRepo) GetForProccessOrders(ctx context.Context) (orders []entity.Order, err error) {
	statusesForCheck := []string{"NEW", "PROCESSING", "REGISTERED", "PROCESSING"}
	sql, args, err := r.Builder.
		Select("order_id, user_id, order_number, status, uploaded_at, accrual").
		From("public.order").
		Where(squirrel.Eq{"status": statusesForCheck}).
		OrderBy("uploaded_at").ToSql()

	if err != nil {
		err = fmt.Errorf("GopherMartRepo - GetForProccessOrders - r.Builder: %w", err)
		return
	}
	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		err = fmt.Errorf("GopherMartRepo - GetForProccessOrders - r.Pool.Query: %w", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		e := entity.Order{}
		err = rows.Scan(&e.OrderID, &e.UserID, &e.OrderNumber, &e.Status, &e.ProcessedAt, &e.Accrual)
		if err != nil {
			return nil, fmt.Errorf("GopherMartRepo - GetForProccessOrders - rows.Scan: %w", err)
		}
		orders = append(orders, e)
	}
	return
}

func (r *GopherMartRepo) UpdateOrder(ctx context.Context, apiResponse entity.AccrualSystemResponse, order entity.Order) (err error) {
	sql, args, err := r.Builder.
		Update("public.order").
		Set("accrual", apiResponse.Accrual).
		Set("status", apiResponse.OrderStatus).
		Where(squirrel.Eq{"order_number": order.OrderNumber}).
		ToSql()
	if err != nil {
		return fmt.Errorf("GopherMartRepo - UpdateOrder - r.Builder: %w", err)
	}
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("GopherMartRepo - UpdateOrder - r.Pool: %w", err)
	}

	balance, err := r.GetBalance(ctx, order.UserID)
	if err != nil {
		return fmt.Errorf("GopherMartRepo - UpdateOrder - GetBalance: %w", err)
	}
	balance.Current += apiResponse.Accrual
	log.Println("!!!!!Updated balance", balance)
	err = r.UpdateBalance(ctx, order.UserID, balance)
	if err != nil {
		return fmt.Errorf("GopherMartRepo - UpdateOrder - UpdateBalance: %w", err)
	}
	return
}
