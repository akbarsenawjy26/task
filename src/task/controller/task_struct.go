package controller

type CreateTaskRequest struct {
	Description string  `json:"description"`
	Status      *string `json:"status"`
}

type UpdateTaskRequest struct {
	Description string  `json:"description"`
	Status      *string `json:"status"`
	ID          int32   `json:"id"`
}
