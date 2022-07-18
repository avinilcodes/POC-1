package user

import (
	"encoding/json"
	"net/http"
	"taskmanager/api"
	"taskmanager/app"
	"taskmanager/db"
	"taskmanager/utils"
	"time"
)

func isBadRequest(err error) bool {
	return err == errEmptyName || err == errEmptyPassword
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
		var aur AddUserRequest
		err := json.NewDecoder(req.Body).Decode(&aur)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		now := time.Now()
		var user db.User
		user.ID = utils.GetUniqueId()
		user.Name = aur.Name
		user.Email = aur.Email
		user.Password = aur.Password
		user.RoleType = aur.RoleType
		user.CreatedAt = now
		user.UpdatedAt = now
		err = service.addUser(req.Context(), user)
		if err != nil {
			app.GetLogger().Warn("error creating user", err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Header().Add("Content-Type", "application/json")
		api.Success(rw, http.StatusOK, api.Response{Message: "User added Successfully"})
	})
}
