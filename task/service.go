package task

import (
	"context"
	"taskmanager/app"
	"taskmanager/db"

	"go.uber.org/zap"
)

type Service interface {
	addTask(ctx context.Context, task db.Task) (err error)
	assignTask(ctx context.Context, assignTaskRequest AssignTaskRequest) (err error)
	listTasks(ctx context.Context) (tasks []db.Task, err error)
	updateTaskStatus(ctx context.Context, description string, status string) (err error)
}

type taskService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (ts *taskService) addTask(ctx context.Context, task db.Task) (err error) {
	err = ts.store.CreateTask(ctx, task)
	if err != nil {
		app.GetLogger().Warn("Error while adding task", err.Error())
		return
	}
	return
}

func (ts *taskService) assignTask(ctx context.Context, assignTaskRequest AssignTaskRequest) (err error) {
	err = ts.store.AssignTask(ctx, assignTaskRequest.Description, assignTaskRequest.UserEmail)
	if err != nil {
		app.GetLogger().Warn("Error while adding task", err.Error())
		return
	}
	return
}
func (ts *taskService) listTasks(ctx context.Context) (tasks []db.Task, err error) {
	tasks, err = ts.store.ListTasks(ctx)
	if err != nil {
		app.GetLogger().Warn("Error while adding task", err.Error())
		return
	}
	return
}

func (ts *taskService) updateTaskStatus(ctx context.Context, description string, status string) (err error) {
	err = ts.store.UpdateTaskStatus(ctx, description, status)
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
