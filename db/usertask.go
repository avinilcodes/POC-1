package db

import (
	"context"
	"time"
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
func (s *store) AssignTask(ctx context.Context, userID string, taskID string) (err error) {
	_, err = s.db.Query(updateTaskStatus, "scoped", time.Time{}, taskID)
	if err != nil {
		return err
	}
	s.db.Query(userTaskInsert, userID, taskID)
	if err != nil {
		return err
	}
	return
}
