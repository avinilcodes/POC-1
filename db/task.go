package db

import "time"

const ()

type Task struct {
	ID             string    `json:"id" db:"id"`
	Description    string    `json:"description" db:"descreption"` // change to description
	TaskStatusCode string    `json:"taskstatuscode" db:"taskstatuscode"`
	StartedAt      time.Time `json:"started_at" db:"started_at"`
	EndedAt        time.Time `json:"ended_at" db:"ended_at"`
}
