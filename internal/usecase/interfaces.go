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
	NewOrder(userID, orderNumber string) (order entity.Order, err error)
}
type IGopherMartRepo interface {
	SaveUser(
		ctx context.Context,
		userID, userName, password string) (user entity.User, err error)
	GetUser(
		ctx context.Context,
		userName string) (user entity.User, err error)
}

type IGopherMartWebAPI interface {
	RegisterUser(
		ctx context.Context,
		userName string,
		password string) (user entity.User, err error)
	Login(ctx context.Context,
		userName string,
		password string) (user entity.User, err error)
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
