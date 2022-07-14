package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"taskmanager/api"
	"taskmanager/app"
	"taskmanager/db"
	"taskmanager/utils"
	"time"

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

func ListUserHandler(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		users, err := service.listUsers(req.Context())
		if err != nil {
			app.GetLogger().Warn("error fetching users")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		respBytes, err := json.Marshal(users)
		if err != nil {
			app.GetLogger().Warn("error while marshilling users")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.Write(respBytes)
	})
}

func AddUserHandler(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		name := req.Form.Get("name")
		email := req.Form.Get("email")
		password := req.Form.Get("password")
		role_type := req.Form.Get("role_type")
		now := time.Now()
		var user db.User
		user.ID = utils.GetUniqueId()
		user.Name = name
		user.Email = email
		user.Password = password
		user.RoleType = role_type
		user.CreatedAt = now
		user.UpdatedAt = now
		fmt.Println(user)
		err := service.addUser(req.Context(), user)
		if err != nil {
			app.GetLogger().Warn("error creating user", err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Header().Add("Content-Type", "application/json")
	})
}
