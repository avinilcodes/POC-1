package login

import "github.com/golang-jwt/jwt/v4"

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	RoleType string `json:"roleType"`
	jwt.StandardClaims
}
