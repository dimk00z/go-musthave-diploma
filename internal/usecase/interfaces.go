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
	GetUser(
		ctx context.Context,
		userName string,
		password string) (user entity.User, err error)
}
type IGopherMartRepo interface {
	SaveUser(
		ctx context.Context,
		userID, userName, password string) (user entity.User, err error)
	GetUser(
		ctx context.Context,
		userName string,
		password string) (user entity.User, err error)
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
}
