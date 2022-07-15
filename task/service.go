package task

import (
	"context"
	"taskmanager/app"
	"taskmanager/db"

	"go.uber.org/zap"
)

type Service interface {
	addTask(ctx context.Context, task db.Task) (err error)
}

type taskService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (cs *taskService) addTask(ctx context.Context, task db.Task) (err error) {
	err = cs.store.CreateTask(ctx, task)
	if err != nil {
		app.GetLogger().Warn("Error while adding task", err.Error())
		return
	}
	return
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &taskService{
		store:  s,
		logger: l,
	}
}
