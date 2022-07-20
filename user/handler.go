// Package classification of Users API
//
// Documentation for users API
//
// schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// -application/json
//
// Produces:
// -application/json
// swagger:meta
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

// A list of users returns in the response
// swagger:response usersResponse
type usersResponse struct {
	// All products in the database
	// in: query
	Body []db.User
}

// swagger:response noContent
type UserAdd struct{}

// swagger:route GET /users users listUsers
// Returns list of users
// responses:
//  200: usersResponse
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

// swagger:route POST /user user addUser
// responses:
//  201: noContent
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
