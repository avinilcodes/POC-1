package user

import "errors"

var (
	errEmptyPassword = errors.New("User password must be presnet")
	errEmptyName     = errors.New("User Name  must be presnet")
	errNoUserId      = errors.New("User is not present")
)
