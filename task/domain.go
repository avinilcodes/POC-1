package task

type AddTaskRequest struct {
	Description string
}

type AssignTaskRequest struct {
	TaskId string
	UserId string
}

type UpdateTaskRequest struct {
	Id     string
	Status string `json:"status"`
}
