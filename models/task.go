package models

type Status string

const (
	Todo       Status = "TODO"
	InProgress Status = "IN_PROGRESS"
	Completed  Status = "COMPLETED"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}
