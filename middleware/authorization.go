package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"taskmanager/api"
	"taskmanager/app"
	"taskmanager/config"
	"taskmanager/login"

	"github.com/golang-jwt/jwt"
)

func AuthorizationMiddleware(handler http.Handler, roleType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//routeName := r.URL.RequestURI()
		token, err := readToken(r)
		claims, err := validateToken(token)
		if err != nil {
			app.GetLogger().Warn("Unauthorized user for %v", r.URL.RequestURI(), claims)
			api.Error(w, http.StatusUnauthorized, api.Response{Message: err.Error()})
			return
		}

		if !strings.EqualFold(claims.RoleType, roleType) {
			app.GetLogger().Warn("Access denied for %v for %v", claims.Email, r.URL.RequestURI())
			api.Error(w, http.StatusForbidden, api.Response{Message: "Access Denied"})
			return
		}

		handler.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "claims", &claims)))
	}
}

//readToken method will read the Authorization header and will return the token string or error
func readToken(r *http.Request) (token string, err error) {
	//TODO: header missing error
	authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
	if len(authHeader) != 2 {
		err = errors.New("Malform header")
		return
	}
	token = authHeader[1]
	return
}

//validateToken will validate the given token, and it will return the claims or error
func validateToken(jwtToken string) (login.Claims, error) {
	if jwtToken == "" {
		err := errors.New("Authorization token is missing")
		return login.Claims{}, err
	}
	// Parse the token
	token, err := jwt.ParseWithClaims(jwtToken, &login.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.InitJWTConfiguration().JwtSignature), nil
	})

	if err != nil {
		err = errors.New("Token is invalid")
		return login.Claims{}, err
	}
	claims := token.Claims.(*login.Claims)
	return *claims, err
}
