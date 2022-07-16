package server

import (
	"net/http"
	"taskmanager/login"
	"taskmanager/middleware"
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

	//Add user
	router.HandleFunc("/user", middleware.AuthorizationMiddleware(user.AddUserHandler(dep.UserServices), "super_admin")).Methods(http.MethodPost)
	router.HandleFunc("/user/v1", middleware.AuthorizationMiddleware(user.AddUserHandler(dep.UserServices), "admin")).Methods(http.MethodPost)
	//ListUsers
	router.HandleFunc("/users", middleware.AuthorizationMiddleware(user.ListUserHandler(dep.UserServices), "admin")).Methods(http.MethodGet)
	//router.HandleFunc("/users", user.ListUserHandler(dep.UserServices)).Methods(http.MethodGet)
	//Create a task
	router.HandleFunc("/task", middleware.AuthorizationMiddleware(task.AddTaskHandler(dep.TaskService), "admin")).Methods(http.MethodPost)
	//List Tasks
	router.HandleFunc("/tasks", task.ListTaskHandler(dep.TaskService)).Methods(http.MethodGet)
	//Assign Task
	router.HandleFunc("/task/assign", middleware.AuthorizationMiddleware(task.AssignTaskHandler(dep.TaskService), "admin")).Methods(http.MethodPost)
	//Update task status
	router.HandleFunc("/task", task.UpdateTaskStatusHandler(dep.TaskService)).Methods(http.MethodPut)
	return
}
