package user

import (
	"bankapp/api"
	"bankapp/app"
	"bankapp/db"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Update(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var c updateRequest
		userID := mux.Vars(req)["user_id"]
		if userID == "" {
			app.GetLogger().Warn(errNoUserId.Error(), "msg", userID, "user", req)
			api.Error(rw, http.StatusBadRequest, api.Response{
				Message: errNoUserId.Error(),
			})
			return
		}

		err := json.NewDecoder(req.Body).Decode(&c)
		if err != nil {
			app.GetLogger().Warn("Error updating user", "msg", err.Error(), "user", req.Body)
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		err = service.update(req.Context(), c, userID)
		if isBadRequest(err) {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, api.Response{Message: "Updated user Successfully"})
	})

}
func isBadRequest(err error) bool {
	return err == errEmptyName || err == errEmptyPassword
}

func DeleteByID(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		userID := vars["user_id"]
		if userID == "" {
			app.GetLogger().Warn(errNoUserId.Error(), "msg", "user", req)
			api.Error(rw, http.StatusBadRequest, api.Response{
				Message: errNoUserId.Error(),
			})
			return
		}

		err := service.deleteByID(req.Context(), userID)
		if err == db.ErrUserNotExist {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}

		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, api.Response{Message: "Deleted user Successfully"})
	})
}
