package task

import (
	"context"
	"taskmanager/app"
	"taskmanager/db"

	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

type Service interface {
	addTask(ctx context.Context, task db.Task) (err error)
	assignTask(ctx context.Context, assignTaskRequest AssignTaskRequest) (err error)
	listTasks(ctx context.Context, token string) (tasks []db.Task, err error)
	updateTaskStatus(ctx context.Context, id string, status string, token string) (err error)
	listUserTask(ctx context.Context) (usertask []db.NameUserTask, err error)
}

type taskService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func ReturnUserEmail(token string) string {
	tokenString := token
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})
	if err != nil {
		return err.Error()
	}
	var email string
	for key, val := range claims {
		if key == "email" {
			email = val.(string)
		}
	}
	return email
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
	err = ts.store.AssignTask(ctx, assignTaskRequest.UserId, assignTaskRequest.TaskId)
	if err != nil {
		app.GetLogger().Warn("Error while assigning task", err.Error())
		return
	}
	return
}
func (ts *taskService) listTasks(ctx context.Context, token string) (tasks []db.Task, err error) {
	email := ReturnUserEmail(token)
	tasks, err = ts.store.ListTasks(ctx, email)
	if err != nil {
		app.GetLogger().Warn("Error while fetching tasks", err.Error())
		return
	}
	return
}

func (ts *taskService) listUserTask(ctx context.Context) (usertask []db.NameUserTask, err error) {
	usertask, err = ts.store.ListUserTask(ctx)
	if err != nil {
		app.GetLogger().Warn("Error while fetching tasks", err.Error())
		return
	}
	return
}

func (ts *taskService) updateTaskStatus(ctx context.Context, id string, status string, token string) (err error) {
	email := ReturnUserEmail(token)
	err = ts.store.UpdateTaskStatus(ctx, id, status, email)
	if err != nil {
		app.GetLogger().Warn("Error while updating task", err.Error())
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
