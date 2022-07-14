package db

const ()

type UserTask struct {
	ID     string `json:"id" db:"id"`
	UserID string `json:"user_id" db:"user_id"`
	TaskID string `json:"task_id" db:"task_id"`
}
