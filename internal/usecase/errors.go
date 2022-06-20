package usecase

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrWrongPassword     = errors.New("wrong password")
	ErrJWT               = errors.New("wrong JWT")
)
