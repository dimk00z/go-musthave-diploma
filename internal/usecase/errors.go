package usecase

import "errors"

var (
	ErrUserAlreadyExists       = errors.New("user already exists")
	ErrWrongPassword           = errors.New("wrong password")
	ErrJWT                     = errors.New("wrong JWT")
	ErrOrderGotByDifferentUser = errors.New("current order has been loaded by another user")
	ErrOrderAlreadyGot         = errors.New("current order has been loaded by current user")
)
