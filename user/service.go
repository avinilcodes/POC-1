package user

import (
	"context"
	"taskmanager/app"
	"taskmanager/db"

	"go.uber.org/zap"
)

type Service interface {
	update(ctx context.Context, req updateRequest, userId string) (err error)
	listUsers(ctx context.Context) (users []db.User, err error)
	addUser(ctx context.Context, user db.User) (err error)
}

type userService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (cs *userService) update(ctx context.Context, c updateRequest, userID string) (err error) {
	err = c.Validate()
	if err != nil {
		cs.logger.Error("Invalid Request for user update", "err", err.Error(), "users", c)
		return
	}

	err = cs.store.UpdateUser(ctx, &db.User{
		Password: c.Password, //pass encrypt
		Name:     c.Name,
		ID:       userID,
	})
	if err != nil {
		cs.logger.Error("Error updating User", "err", err.Error(), "users", c)
		return
	}

	return
}

func (cs *userService) listUsers(ctx context.Context) (users []db.User, err error) {
	users, err = cs.store.ListUsers(ctx)
	if err != nil {
		app.GetLogger().Warn(err.Error())
		return
	}
	return
}
func (cs *userService) addUser(ctx context.Context, user db.User) (err error) {
	err = cs.store.CreateUser(ctx, user)
	if err != nil {
		app.GetLogger().Warn("Error while adding user", err.Error())
		return
	}
	return
}
func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &userService{
		store:  s,
		logger: l,
	}
}
