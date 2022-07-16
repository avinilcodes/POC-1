package db

import (
	"context"
)

const (
	userTaskInsert = `INSERT INTO users_tasks (user_id,task_id) VALUES ($1,$2)`
)

type UserTask struct {
	ID     string `json:"id" db:"id"`
	UserID string `json:"user_id" db:"user_id"`
	TaskID string `json:"task_id" db:"task_id"`
}

//need to add more, only 2 tasks can be assigned to a user , 1 task can be assigned to 2 users
func (s *store) AssignTask(ctx context.Context, description string, userEmail string) (err error) {
	var user User
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &user, findUserByEmailQuery, userEmail)
	})
	if err != nil {
		return err
	}
	userID := user.ID
	var task Task
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &task, findTaskIDByDescription, description)
	})
	taskID := task.ID
	if task.TaskStatusCode == "not_scoped" {
		task.TaskStatusCode = "scoped"
		_, err = s.db.Query(updateTaskStatus, task.TaskStatusCode, task.ID)
		if err != nil {
			return err
		}
	}

	s.db.Query(userTaskInsert, userID, taskID)
	if err != nil {
		return err
	}
	return
}
