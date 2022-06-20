package usecase

import (
	"context"
	"log"

	"github.com/dimk00z/go-musthave-diploma/internal/entity"
	uuid "github.com/satori/go.uuid"
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
	hashedPassword, err := uc.webAPI.GetPasswordHash(password)
	log.Println(userName, hashedPassword)
	userID := uuid.NewV4().String()
	user, err = uc.repo.SaveUser(ctx, userID, userName, hashedPassword)
	if err != nil {
		return
	}
	// TODO add balace
	return
}

func (uc *GopherMartUseCase) GetUser(
	ctx context.Context,
	userName string,
	password string) (user entity.User, err error) {
	// if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
	// 	err = ErrWrongPassword
	// 	return
	// }
	return
}
