package db

import "errors"

var (
	ErrAccountNotExist    = errors.New("Account Id does not exist in db")
	ErrUserNotExist       = errors.New("User does not exist in db")
	ErrCreatingAccountant = errors.New("Error while creating Accountant")
	ErrEmailAlreadyExist  = errors.New("Email already exist!")
)
