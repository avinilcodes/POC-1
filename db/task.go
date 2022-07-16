package db

import (
	"context"
	"time"
)

const (
	findTaskIDByDescription = "SELECT id,descreption,task_status_code,started_at,ended_at FROM TASKS WHERE descreption=$1"
	insertTask              = `INSERT INTO tasks (id,descreption,task_status_code,started_at,ended_at) VALUES ($1,$2,$3,$4,$5)`
	findAllTasks            = "select * from tasks"
	updateTaskStatus        = `update tasks set task_status_code=$1 where id =$2`
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

func (s *store) UpdateTaskStatus(ctx context.Context, description string, status string) (err error) {
	var task Task
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &task, findTaskIDByDescription, description)
	})
	task.TaskStatusCode = status
	flag := status != "in_progress" && status != "mr_approved" && status != "code_review"
	if !flag {
		if status == "mr_approved" {
			task.EndedAt = time.Now()
		}
		_, err = s.db.Query(updateTaskStatus, task.TaskStatusCode, task.ID)
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
