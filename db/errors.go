package db

import "errors"

var (
	ErrUserNotExist              = errors.New("User does not exist in db")
	ErrTaskAlreadyExist          = errors.New("Task already exist!")
	ErrTaskCannotBeUpdated       = errors.New("Task status invalid")
	ErrTaskStatusError           = errors.New("Task cannot be updated previous state/states pending")
	ErrTaskAssignedToAnotherUser = errors.New("Task assignee is different")
	ErrOnlyAdminAccess           = errors.New("Admin access only")
)
