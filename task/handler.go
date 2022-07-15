package task

import (
	"net/http"
	"taskmanager/api"
	"taskmanager/app"
	"taskmanager/db"
	"taskmanager/utils"
	"time"
)

type AssignTaskRequest struct {
	Description string
	UserEmail   string
}

func AddTaskHandler(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		description := req.Form.Get("description")
		taskstatuscode := req.Form.Get("taskstatuscode")
		now := time.Now()
		var task db.Task
		task.ID = utils.GetUniqueId()
		task.Description = description
		task.TaskStatusCode = taskstatuscode
		task.StartedAt = now
		task.EndedAt = time.Time{}
		err := service.addTask(req.Context(), task)
		if err != nil {
			if err.Error() == "Task already exist!" {
				rw.WriteHeader(http.StatusBadRequest)
				api.Error(rw, http.StatusBadRequest, api.Response{Message: "can not duplicate Task"})
				return
			} else {
				app.GetLogger().Warn("Failed because", err)
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		rw.Header().Add("Content-Type", "application/json")
		api.Success(rw, http.StatusOK, api.Response{Message: "Task added Successfully"})
	})
}

func AssignTaskHandler(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		description := req.Form.Get("description")
		userEmail := req.Form.Get("email")
		var assignTaskRequest AssignTaskRequest
		assignTaskRequest.Description = description
		assignTaskRequest.UserEmail = userEmail
		err := service.assignTask(req.Context(), assignTaskRequest)
		if err != nil {
			app.GetLogger().Warn("error creating task", err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Header().Add("Content-Type", "application/json")
		api.Success(rw, http.StatusOK, api.Response{Message: "Task assignment Successful"})
	})
}
