package usecase

import (
	"context"
	"log"

	"github.com/dimk00z/go-musthave-diploma/internal/entity"
	"github.com/dimk00z/go-musthave-diploma/pkg/logger"
	uuid "github.com/satori/go.uuid"
)

type GopherMartUseCase struct {
	repo   IGopherMartRepo
	webAPI IGopherMartWebAPI
	l      logger.Interface
}

func New(r IGopherMartRepo, w IGopherMartWebAPI, l logger.Interface) *GopherMartUseCase {
	return &GopherMartUseCase{
		repo:   r,
		webAPI: w,
		l:      l,
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
	return
}

func (uc *GopherMartUseCase) Login(
	ctx context.Context,
	userName string,
	password string) (token string, err error) {
	user, err := uc.repo.GetUser(ctx, userName)
	if err != nil {
		return
	}
	uc.l.Info("GopherMartUseCase - Login - : " + user.UserId + " " + " " + user.Login)

	err = uc.webAPI.VerifyPassword(password, user.Password)

	if err != nil {
		return
	}
	token, err = uc.webAPI.GenerateToken(user.UserId)
	if err != nil {
		return "", err
	}
	uc.l.Info("GopherMartUseCase - Login - token: " + token)
	return
}

func (uc *GopherMartUseCase) GetUser(
	ctx context.Context,
	userName string) (user entity.User, err error) {
	// if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
	// 	err = ErrWrongPassword
	// 	return
	// }
	return
}
