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
	DeleteAt    *time.Time `json:"deleteAt"`
	UpdateAt    *time.Time `json:"updateAt"`
}

func NewTask(description string) *Task {
	now := time.Now()
	return &Task{
		ID:          uuid.New(),
		Description: description,
		Status:      "TODO",
		CreateAt:    &now,
		DeleteAt:    nil,
		UpdateAt:    nil,
	}
}

func (t *Task) UpdateStatus() {
	now := time.Now()
	t.Status = "in-progress"
	t.UpdateAt = &now
}

func (t *Task) Done() {
	now := time.Now()
	t.Status = "Done"
	t.DeleteAt = &now
}
