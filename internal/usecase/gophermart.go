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

func (uc *GopherMartUseCase) NewOrder(ctx context.Context, userID string, orderNumber string) (order entity.Order, err error) {
	checkedOrder, err := uc.repo.GetOrder(ctx, orderNumber)
	uc.l.Debug(checkedOrder)
	if checkedOrder != (entity.Order{}) {
		if checkedOrder.UserID != userID {
			err = ErrOrderGotByDifferentUser
		} else {
			err = ErrOrderAlreadyGot
		}
		return checkedOrder, err
	}
	if err != nil {
		uc.l.Error(err)
		return
	}
	orderID := uuid.NewV4().String()
	order, err = uc.repo.NewOrder(ctx, userID, orderID, orderNumber)
	return
}

func (uc *GopherMartUseCase) GetOrders(ctx context.Context, userID string) (orders []entity.Order, err error) {
	orders, err = uc.repo.GetOrders(ctx, userID)
	if err != nil {
		return
	}
	return uc.repo.GetOrders(ctx, userID)
}
func (uc *GopherMartUseCase) GetOrder(ctx context.Context,
	orderNumber string, userID string) (order entity.Order, err error) {
	order, err = uc.repo.GetOrder(ctx, orderNumber)
	if err != nil {
		return
	}
	return
}

func (uc *GopherMartUseCase) GetBalance(ctx context.Context, userID string) (balance entity.Balance, err error) {
	return uc.repo.GetBalance(ctx, userID)
}

func (uc *GopherMartUseCase) Withdraw(ctx context.Context, userID string, orderNumber string, sum float32) (err error) {
	balance, err := uc.repo.GetBalance(ctx, userID)
	if err != nil {
		uc.l.Error(err)
	}
	if balance.Current < float32(sum) {
		err = ErrNotEnoughFunds
		return
	}
	_, err = uc.GetOrder(ctx, orderNumber, userID)
	if err != nil {
		err = ErrWrongOrder
		return
	}
	balance.Current = balance.Current - sum
	balance.Spend += sum
	err = uc.repo.UpdateBalance(ctx, userID, balance)
	// TODO add saving Withdraw
	// err=uc.repo.SaveWithdraw()

	return
}

func (uc *GopherMartUseCase) GetWithdrawals(
	ctx context.Context,
	userID string) (withdrawals []entity.Withdrawal, err error) {

	return uc.repo.GetWithdrawals(ctx, userID)
}
