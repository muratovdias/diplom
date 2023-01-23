package service

import "errors"

var (
	EmailError    = errors.New("wrong email")
	PasswordError = errors.New("wrong password")
)
