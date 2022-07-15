package db

import (
	"context"
	"fmt"
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
	userID := user.ID
	var task Task
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &task, findTaskIDByDescription, description)
	})
	fmt.Println(user)
	fmt.Println(task)
	taskID := task.ID
	if task.TaskStatusCode == "not_scoped" {
		task.TaskStatusCode = "scoped"
		_, err = s.db.Query(`update tasks set task_status_code=$1 where id =$2`, task.TaskStatusCode, task.ID)
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
