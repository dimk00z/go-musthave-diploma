package usecase

import (
	"context"

	"github.com/dimk00z/go-musthave-diploma/internal/entity"
)

type GopherMartUseCase struct {
	repo   IGopherMartRepo
	webAPI IGopherMartWebAPI
}

func New(r IGopherMartRepo, w IGopherMartWebAPI) *GopherMartUseCase {
	return &GopherMartUseCase{
		repo:   r,
		webAPI: w,
	}
}

func (uc *GopherMartUseCase) RegisterUser(
	ctx context.Context,
	userName string,
	password string) (user entity.User, err error) {
	return
}

func (uc *GopherMartUseCase) GetUser(
	ctx context.Context,
	userName string,
	password string) (user entity.User, err error) {
	return
}
