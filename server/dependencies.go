package server

import (
	"taskmanager/app"
	"taskmanager/db"
	"taskmanager/login"
	"taskmanager/task"
	"taskmanager/user"
	"taskmanager/useraccount"
	"taskmanager/utils"
)

type dependencies struct {
	UserLoginService   login.Service
	UserServices       user.Service
	UserAccountService useraccount.Service
	TaskService        task.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)

	// call new service
	userService := user.NewService(dbStore, logger)

	taskService := task.NewService(dbStore, logger)
	loginService := login.NewService(dbStore, logger)

	userAccountService := useraccount.NewService(dbStore, logger)
	err := db.CreateSuperAdmin(dbStore)
	if err != nil && !utils.CheckIfDuplicateKeyError(err) {
		return dependencies{}, err
	}

	return dependencies{
		UserServices:       userService,
		UserLoginService:   loginService,
		UserAccountService: userAccountService,
		TaskService:        taskService,
	}, nil
}
