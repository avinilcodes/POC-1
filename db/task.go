package db

import (
	"context"
	"time"
)

const (
	findTaskIDByDescription = "SELECT id,description,task_status_code,started_at,ended_at FROM TASKS WHERE description=$1"
	insertTask              = `INSERT INTO tasks (id,description,task_status_code,started_at,ended_at) VALUES ($1,$2,$3,$4,$5)`
	findAllTasks            = "select * from tasks"
	updateTaskStatus        = `update tasks set task_status_code=$1 where id =$2`
	findUserFromTaskId      = `select * from users where id = (select user_id from users_tasks where task_id = $1)`
)

type Task struct {
	ID             string    `json:"id" db:"id"`
	Description    string    `json:"description" db:"description"`
	TaskStatusCode string    `json:"taskstatuscode" db:"task_status_code"`
	StartedAt      time.Time `json:"started_at" db:"started_at"`
	EndedAt        time.Time `json:"ended_at" db:"ended_at"`
}

func (s *store) CreateTask(ctx context.Context, task Task) (err error) {
	res, err := s.db.Exec(findTaskIDByDescription, task.Description)
	if err != nil {
		return
	}
	cnt, _ := res.RowsAffected()
	if cnt == 0 {
		_, err = s.db.Query(insertTask, task.ID, task.Description, task.TaskStatusCode, task.StartedAt, task.EndedAt)
		if err != nil {
			return err
		}
		return
	}
	return ErrTaskAlreadyExist

}

func (s *store) UpdateTaskStatus(ctx context.Context, description string, status string, userEmail string) (err error) {
	var task Task
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &task, findTaskIDByDescription, description)
	})
	if err != nil {
		return err
	}
	if (task.TaskStatusCode == "in_progress" || task.TaskStatusCode == "scoped" || task.TaskStatusCode == "not_scoped") && status == "mr_approved" {
		return ErrTaskStatusError
	}
	if task.TaskStatusCode == "not_scoped" && status == "in_progress" {
		return ErrTaskStatusError
	}
	if (task.TaskStatusCode == "scoped" || task.TaskStatusCode == "not_scoped") && status == "code_review" {
		return ErrTaskStatusError
	}
	var user User
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &user, findUserFromTaskId, task.ID)
	})
	if err != nil {
		return err
	}
	var currentUser User
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &currentUser, findUserByEmailQuery, userEmail)
	})
	if err != nil {
		return
	}
	if user.Email != userEmail && currentUser.Email != "admin" {
		return ErrTaskAssignedToAnotherUser
	}
	if status == "mr_approved" && currentUser.RoleType != "admin" {
		return ErrOnlyAdminAccess
	}
	flag := status != "in_progress" && status != "mr_approved" && status != "code_review"
	if !flag {
		if status == "mr_approved" {
			task.EndedAt = time.Now()
		}
		_, err = s.db.Query(updateTaskStatus, status, task.ID)
		if err != nil {
			return err
		}
		return
	}
	return ErrTaskCannotBeUpdated

}

func (s *store) ListTasks(ctx context.Context) (tasks []Task, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.SelectContext(ctx, &tasks, findAllTasks)
	})
	return
}
