package task

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID  `json:"id"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreateAt    *time.Time `json:"createAt"`
	deleteAt    *time.Time `json:"deleteAt"`
	updateAt    *time.Time `json:"updateAt"`
}

func NewTask(description string) *Task {
	now := time.Now()
	return &Task{
		ID:          uuid.New(),
		Description: description,
		Status:      "TODO",
		CreateAt:    &now,
		deleteAt:    nil,
		updateAt:    nil,
	}
}
