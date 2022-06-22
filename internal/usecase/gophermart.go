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
	wp     IWorkerPool
}

func New(r IGopherMartRepo, w IGopherMartWebAPI, l logger.Interface, wp IWorkerPool) *GopherMartUseCase {
	return &GopherMartUseCase{
		repo:   r,
		webAPI: w,
		l:      l,
		wp:     wp,
	}
}

func (uc *GopherMartUseCase) RegisterUser(
	ctx context.Context,
	userName string,
	password string) (user entity.User, err error) {

	hashedPassword, err := uc.webAPI.GetPasswordHash(password)
	if err != nil {
		return
	}
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
	password string) (user entity.User, err error) {
	user, err = uc.repo.GetUser(ctx, userName)
	if err != nil {
		return
	}
	uc.l.Info("GopherMartUseCase - Login - : " + user.UserID + " " + " " + user.Login)

	err = uc.webAPI.VerifyPassword(password, user.Password)

	if err != nil {
		return
	}
	return
}

func (uc *GopherMartUseCase) GetUserToken(
	userID string) (token string, err error) {

	return uc.webAPI.GenerateToken(userID)
}

func (uc *GopherMartUseCase) ParseToken(tokenString string) (userID string, err error) {

	return uc.webAPI.ParseToken(tokenString)
}

func (uc *GopherMartUseCase) NewOrder(ctx context.Context, userID, orderNumber string) (order entity.Order, err error) {
	orderID := uuid.NewV4().String()
	order, err = uc.repo.NewOrder(ctx, userID, orderID, orderNumber)
	return
}

func (uc *GopherMartUseCase) GetOrders(ctx context.Context, userID string) (orders []entity.Order, err error) {
	return uc.repo.GetOrders(ctx, userID)
}
