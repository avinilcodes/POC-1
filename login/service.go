package login

import (
	"context"
	"errors"
	"strings"
	"taskmanager/config"
	"taskmanager/db"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	login(ctx context.Context, req loginRequest) (tokenString string, err error)
}

type loginService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (ls *loginService) login(ctx context.Context, ul loginRequest) (tokenString string, err error) {
	ul.Email = strings.ToLower(ul.Email)
	user, err := ls.store.FindUserByEmail(ctx, ul.Email)
	// TODO: Handle the err
	if err == db.ErrUserNotExist {
		ls.logger.Warn("User Not present", "warn", err.Error(), "email ", ul.Email)
		return
	}
	if user.Email == "" {
		err = errors.New("Invalid Email or Password")
		return
	}
	// Authenticate the user
	matched := authenticateUser(user, ul.Password)
	if !matched {
		err = errors.New("Invalid Email or Password")
		return
	}
	ls.logger.Info("User is valid, generating the token")
	// Generate the Token
	tokenString, err = generateJWT(user.ID, user.Email, user.RoleType)
	return
}

func authenticateUser(user db.User, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd))
	return err == nil
}

func generateJWT(id, email, roleType string) (tokenString string, err error) {
	expirationTime := time.Now().Add(time.Minute * time.Duration(config.InitJWTConfiguration().TokenExpiry)).Unix()
	claims := &Claims{
		ID:       id,
		Email:    email,
		RoleType: roleType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(config.InitJWTConfiguration().JwtSignature))
	if err != nil {
		return
	}
	return
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &loginService{
		store:  s,
		logger: l,
	}
}
