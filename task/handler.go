package task

import (
	"encoding/json"
	"net/http"
	"strings"
	"taskmanager/api"
	"taskmanager/app"
	"taskmanager/db"
	"taskmanager/utils"
	"time"
)

func AddTaskHandler(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var atr AddTaskRequest
		err := json.NewDecoder(req.Body).Decode(&atr)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		now := time.Now()
		var task db.Task
		task.ID = utils.GetUniqueId()
		task.Description = atr.Description
		task.TaskStatusCode = "not_scoped"
		task.StartedAt = now
		task.EndedAt = time.Time{}
		err = service.addTask(req.Context(), task)
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
		api.Success(rw, http.StatusOK, api.Response{Message: "Task added Successfully"})
	})
}

func AssignTaskHandler(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var atr AssignTaskRequest
		err := json.NewDecoder(req.Body).Decode(&atr)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		var assignTaskRequest AssignTaskRequest
		assignTaskRequest.Description = atr.Description
		assignTaskRequest.Email = atr.Email
		err = service.assignTask(req.Context(), assignTaskRequest)
		if err != nil {
			app.GetLogger().Warn("error creating task", err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		api.Success(rw, http.StatusOK, api.Response{Message: "Task assignment Successful"})
	})
}

func ListTaskHandler(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		reqToken := req.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]
		tasks, err := service.listTasks(req.Context(), reqToken)
		if err != nil {
			app.GetLogger().Warn("error fetching tasks", err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		respBytes, err := json.Marshal(tasks)
		if err != nil {
			app.GetLogger().Warn("error while marshilling users")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.Write(respBytes)
	})
}

func UpdateTaskStatusHandler(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var utr UpdateTaskRequest
		err := json.NewDecoder(req.Body).Decode(&utr)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		reqToken := req.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]
		err = service.updateTaskStatus(req.Context(), utr.Description, utr.Status, reqToken)
		if err != nil {
			if err.Error() == "Task status invalid" {
				rw.WriteHeader(http.StatusBadRequest)
				api.Error(rw, http.StatusBadRequest, api.Response{Message: "Invalid status"})
				return
			} else if err.Error() == "Task cannot be updated previous state/states pending" {
				rw.WriteHeader(http.StatusBadRequest)
				api.Error(rw, http.StatusBadRequest, api.Response{Message: "Task Status cannot be updated as previous states are pending"})
				return
			} else if err.Error() == "Task assignee is different" {
				rw.WriteHeader(http.StatusBadRequest)
				api.Error(rw, http.StatusBadRequest, api.Response{Message: "Task Status cannot be updated as it is not assigned to you"})
				return
			} else if err.Error() == "Admin access only" {
				rw.WriteHeader(http.StatusBadRequest)
				api.Error(rw, http.StatusBadRequest, api.Response{Message: "Task Status cannot be updated ask admin to change the status"})
				return
			} else {
				app.GetLogger().Warn("Failed because", err)
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		api.Success(rw, http.StatusOK, api.Response{Message: "Task status update Successful"})
	})
}
