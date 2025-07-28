package task

import (
	"encoding/json"
	"fmt"
	"os"
	"task-tracker/pkg/storage/file"
	"task-tracker/pkg/utils"
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
		return "", utils.ErrorF("Error to provide file", err)
	}
	defer file.Close()

	sliceTasks, err := t.ReadTasks()
	if err != nil {
		return "", err
	}
	fmt.Println("Slice vazio", sliceTasks)
	sliceTasks = append(sliceTasks, *t)

	jsonData, err := json.Marshal(sliceTasks)
	if err != nil {
		return "", utils.ErrorF("Cannot marshal data", err)
	}

	fmt.Println(string(jsonData))
	_, err = file.Write(jsonData)
	if err != nil {
		return "", utils.ErrorF("Could not write a file", err)
	}

	return "Task saved!", nil
}

func (t Task) ReadTasks() ([]Task, error) {
	bData, err := file.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, utils.ErrorF("Error to read task file", err)
	}

	if len(bData) == 0 {
		return []Task{}, nil
	}

	var tasks []Task
	err = json.Unmarshal(bData, &tasks)
	if err != nil {
		return nil, utils.ErrorF("Error to unmarshal json file", err)
	}

	return tasks, nil
}
