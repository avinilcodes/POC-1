package server

import (
	"net/http"
	"taskmanager/login"
	"taskmanager/task"
	"taskmanager/user"

	"github.com/gorilla/mux"
)

const (
//versionHeader = "Accept"
)

func initRouter(dep dependencies) (router *mux.Router) {
	//v1 := fmt.Sprintf("application/vnd.%s.v1", config.AppName())
	router = mux.NewRouter()
	// user
	/*
		router.HandleFunc("/user", middleware.AuthorizationMiddleware(useraccount.Create(dep.UserAccountService), "accountant")).Methods(http.MethodPost).Headers(versionHeader, v1)
		router.HandleFunc("/user/{user_id}", middleware.AuthorizationMiddleware(user.DeleteByID(dep.UserServices), "accountant")).Methods(http.MethodDelete).Headers(versionHeader, v1)
		router.HandleFunc("/user/{user_id}", middleware.AuthorizationMiddleware(user.Update(dep.UserServices), "accountant")).Methods(http.MethodPut)
	*/
	//Login
	router.HandleFunc("/login", login.Login(dep.UserLoginService)).Methods(http.MethodPost)

	//ListUsers
	router.HandleFunc("/users", user.ListUserHandler(dep.UserServices)).Methods(http.MethodGet)

	//Add user
	router.HandleFunc("/user", user.AddUserHandler(dep.UserServices)).Methods(http.MethodPost)
	//Create a task
	router.HandleFunc("/task", task.AddTaskHandler(dep.TaskService)).Methods(http.MethodPost)
	//Assign Task
	router.HandleFunc("/task/assign", task.AssignTaskHandler(dep.TaskService)).Methods(http.MethodPost)
	return
}
