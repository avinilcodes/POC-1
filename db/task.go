package db

import (
	"context"
	"fmt"
	"time"
)

const (
	findTaskIDByDescription = "SELECT id,descreption,task_status_code,started_at,ended_at FROM TASKS WHERE descreption=$1"
	insertTask              = `INSERT INTO tasks (id,descreption,task_status_code,started_at,ended_at) VALUES ($1,$2,$3,$4,$5)`
)

type Task struct {
	ID             string    `json:"id" db:"id"`
	Description    string    `json:"description" db:"descreption"` // change to description
	TaskStatusCode string    `json:"taskstatuscode" db:"task_status_code"`
	StartedAt      time.Time `json:"started_at" db:"started_at"`
	EndedAt        time.Time `json:"ended_at" db:"ended_at"`
}

func (s *store) CreateTask(ctx context.Context, task Task) (err error) {
	res, err := s.db.Exec(findTaskIDByDescription, task.Description)
	if err != nil {
		return
	}
	fmt.Println(res)
	cnt, _ := res.RowsAffected()
	fmt.Println(cnt)
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
	if status == "mr_approved" {
		task.EndedAt = time.Now()
	}
	_, err = s.db.Query(`INSERT INTO tasks (id,descreption,task_status_code,started_at,ended_at) VALUES ($1,$2,$3,$4,$5)`, task.ID, task.Description, task.TaskStatusCode, task.StartedAt, task.EndedAt)
	if err != nil {
		return err
	}
	return
}
