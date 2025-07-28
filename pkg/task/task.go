package task

import (
	"encoding/json"
	"task-tracker/pkg/storage/file"
	"time"
)

type Task struct {
	ID          uint      `json:"id"`
	Description string    `json:"Description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdat"`
	UpdatedAt   time.Time `json:"updatedat"`
}

const fileName = "tasks.json"

func NewTask(description string) *Task {

	return &Task{
		ID:          0,
		Description: description,
		Status:      "TODO",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

}

func (t *Task) Save() (string, error) {
	file, err := file.ProvideFile(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	jsonData, err := json.Marshal(t)
	if err != nil {
		return "", err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return "", err
	}

	return "Task saved!", nil
}
