package task

import "time"

type Task struct {
	ID          int       `json:"id,omitempty"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreateAt    time.Time `json:"createAt"`
	DeleteAt    time.Time `json:"deleteAt,omitempty"`
	UpdateAt    time.Time `json:"updateAt,omitempty"`
}

func NewTask(description string) *Task {
	return &Task{
		ID:          0,
		Description: description,
		Status:      "todo",
		CreateAt:    time.Now(),
	}
}
