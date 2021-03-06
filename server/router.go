package server

import (
	"net/http"
	"taskmanager/login"
	"taskmanager/middleware"
	"taskmanager/task"
	"taskmanager/user"

	runtimemid "github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

const (
//versionHeader = "Accept"
)

func initRouter(dep dependencies) (router *mux.Router) {
	router = mux.NewRouter()
	//Login
	router.HandleFunc("/login", login.Login(dep.UserLoginService)).Methods(http.MethodPost)

	//Add user
	router.HandleFunc("/user", middleware.AuthorizationMiddleware(user.AddUserHandler(dep.UserServices), "super_admin,admin")).Methods(http.MethodPost)

	//ListUsers
	router.HandleFunc("/users", middleware.AuthorizationMiddleware(user.ListUserHandler(dep.UserServices), "admin,super_admin")).Methods(http.MethodGet)

	//Create a task
	router.HandleFunc("/task", middleware.AuthorizationMiddleware(task.AddTaskHandler(dep.TaskService), "admin")).Methods(http.MethodPost)

	//List Tasks
	router.HandleFunc("/tasks", task.ListTaskHandler(dep.TaskService)).Methods(http.MethodGet)

	//Assign Task
	router.HandleFunc("/task/{task_id}/user/{user_id}", middleware.AuthorizationMiddleware(task.AssignTaskHandler(dep.TaskService), "admin")).Methods(http.MethodPost)

	//Update task status
	router.HandleFunc("/task/{id}", task.UpdateTaskStatusHandler(dep.TaskService)).Methods(http.MethodPut)

	//List tasks assigned to user user->task
	router.HandleFunc("/usertask", middleware.AuthorizationMiddleware(task.ListUserTaskHandler(dep.TaskService), "admin,super_admin")).Methods(http.MethodGet)

	ops := runtimemid.RedocOpts{SpecURL: "swagger.yaml"}
	sh := runtimemid.Redoc(ops, nil)

	router.Handle("/docs", sh)
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	return
}
