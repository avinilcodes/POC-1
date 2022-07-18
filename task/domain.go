package task

type AddTaskRequest struct {
	Description string
}

type AssignTaskRequest struct {
	Description string
	Email       string
}

type UpdateTaskRequest struct {
	Description string
	Status      string `json:"status"`
}
