package usecase

import (
	"context"

	"github.com/dimk00z/go-musthave-diploma/internal/entity"
)

type IGopherMart interface {
	RegisterUser(
		ctx context.Context,
		userName string,
		password string) (user entity.User, err error)
	Login(
		ctx context.Context,
		userName string,
		password string) (user entity.User, err error)
	GetUserToken(
		userID string) (token string, err error)
	ParseToken(tokenString string) (userID string, err error)
	NewOrder(ctx context.Context, userID string, orderNumber int) (order entity.Order, err error)
	GetOrders(ctx context.Context, userID string) (orders []entity.Order, err error)
	GetOrder(ctx context.Context, orderNumber int, userID string) (order entity.Order, err error)
	GetBalance(ctx context.Context, userID string) (balance entity.Balance, err error)
	Withdraw(ctx context.Context, userID string, orderNumber int) (err error)
}

type IGopherMartRepo interface {
	SaveUser(
		ctx context.Context,
		userID, userName, password string) (user entity.User, err error)
	GetUser(
		ctx context.Context,
		userName string) (user entity.User, err error)
	NewOrder(ctx context.Context, userID, orderID string, orderNumber int) (order entity.Order, err error)
	GetOrders(ctx context.Context, userID string) (orders []entity.Order, err error)
	GetOrder(ctx context.Context, orderNumber int) (order entity.Order, err error)
	GetBalance(ctx context.Context, userID string) (balance entity.Balance, err error)
}

type IGopherMartWebAPI interface {
	GetPasswordHash(
		password string) (passwordHash string, err error)
	VerifyPassword(
		password, hashedPassword string) (err error)
	GenerateToken(userID string) (token string, err error)
	CheckToken(tokenString string) error
	ParseToken(tokenString string) (userID string, err error)
}

type IWorkerPool interface {
	Push(task func(ctx context.Context) error)
	Run(ctx context.Context)
	Close()
}
