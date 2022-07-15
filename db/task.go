package db

import (
	"context"
	"time"
)

const ()

type Task struct {
	ID             string    `json:"id" db:"id"`
	Description    string    `json:"description" db:"descreption"` // change to description
	TaskStatusCode string    `json:"taskstatuscode" db:"taskstatuscode"`
	StartedAt      time.Time `json:"started_at" db:"started_at"`
	EndedAt        time.Time `json:"ended_at" db:"ended_at"`
}

func (s *store) CreateTask(ctx context.Context, task Task) (err error) {
	_, err = s.db.Query(`INSERT INTO tasks (id,descreption,task_status_code,started_at,ended_at) VALUES ($1,$2,$3,$4,$5)`, task.ID, task.Description, task.TaskStatusCode, task.StartedAt, task.EndedAt)
	if err != nil {
		return err
	}
	return
}
