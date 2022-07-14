package useraccount

import (
	"encoding/json"
	"net/http"
	"taskmanager/api"
	"taskmanager/app"
	"taskmanager/utils"
)

func Create(service Service) http.HandlerFunc {

	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var cr createRequest
		err := json.NewDecoder(req.Body).Decode(&cr)
		if err != nil {
			app.GetLogger().Warn("Error creating user", "msg", err.Error(), "user", req.Body)
			api.Error(rw, http.StatusBadRequest, api.Response{
				Message: err.Error(),
			})
			return
		}

		resp, err := service.create(req.Context(), cr)
		if err != nil {
			if utils.CheckIfDuplicateKeyError(err) {
				app.GetLogger().Warn("Error creating user", "msg", "Email already exist", "user", req.Body)
				api.Error(rw, http.StatusConflict, api.Response{
					Message: "Email already exist",
				})
				return
			}

			app.GetLogger().Error("Error creating user", "msg", err.Error(), "user", req.Body)
			api.Error(rw, http.StatusInternalServerError, api.Response{
				Message: "Error while creating user",
			})
			return
		}

		api.Success(rw, http.StatusCreated, resp)
	})
}
