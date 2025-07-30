package task

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"task-tracker/pkg/utils"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdat"`
	UpdatedAt   time.Time `json:"updatedat"`
}

const fileName = "tasks.json"

func NewTask(description string) *Task {

	return &Task{
		Description: description,
		Status:      "TODO",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}
}

func (t *Task) Add() (string, error) {
	sliceTasks, err := t.Read()
	if err != nil {
		return "", err
	}

	t.ID = getMaxID(sliceTasks) + 1
	sliceTasks = append(sliceTasks, *t)

	err = t.Save(sliceTasks)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Task added successfully (ID: %d)", t.ID), nil
}

func (t *Task) Read() ([]Task, error) {
	bData, err := os.ReadFile(fileName)
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
		return nil, utils.ErrorF("Error to Unmarshal json file", err)
	}

	return tasks, nil
}

func (t *Task) Update(taskID int, description string) error {
	t.Description = description
	t.UpdatedAt = time.Now()

	tasks, err := t.Read()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == taskID {
			tasks[i].Description = *&t.Description
			break
		}
	}

	return t.Save(tasks)
}

func (t *Task) Delete(taskID int) error {
	tasks, err := t.Read()
	if err != nil {
		return err
	}

	indexDelete := -1
	for i, task := range tasks {
		if task.ID == taskID {
			indexDelete = i
			break
		}
	}

	if indexDelete == -1 {
		return fmt.Errorf("Task %v not found", taskID)
	}

	tasks = append(tasks[:indexDelete], tasks[indexDelete+1:]...)
	err = t.Save(tasks)
	if err != nil {
		return utils.ErrorF("Error to delete task", err)
	}

	return nil
}

func (t *Task) MarkInProgress(taskID int) error {
	t.Status = "IN-PROGRESS"
	t.UpdatedAt = time.Now()

	tasks, err := t.Read()
	if err != nil {
		return err
	}

	indexNotFount := -1
	for i := range tasks {
		if tasks[i].ID == taskID {
			tasks[i].Status = *&t.Status
			indexNotFount = i
			break
		}
	}

	if indexNotFount == -1 {
		return fmt.Errorf("Task %v not found", taskID)
	}

	return t.Save(tasks)

}

func (t *Task) MarkDone(taskID int) error {
	t.Status = "DONE"
	t.UpdatedAt = time.Now()

	tasks, err := t.Read()
	if err != nil {
		return err
	}

	indexNotFount := -1
	for i := range tasks {
		if tasks[i].ID == taskID {
			tasks[i].Status = *&t.Status
			indexNotFount = i
			break
		}
	}

	if indexNotFount == -1 {
		return fmt.Errorf("Task %v not found", taskID)
	}

	return t.Save(tasks)

}

func getMaxID(sliceTasks []Task) int {
	var maxID int
	for _, task := range sliceTasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	return maxID
}

func (t *Task) Save(tasks []Task) error {
	jsonData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return utils.ErrorF("Cannot marshal data", err)
	}

	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		return utils.ErrorF("Could not write a file", err)
	}

	return nil
}

func (t Task) List() ([]Task, error) {
	return t.Read()
}

func (t Task) ListTaskByStatus(status string) ([]Task, error) {
	tasks, err := t.Read()
	if err != nil {
		return nil, err
	}

	sliceTasks := []Task{}
	for _, task := range tasks {
		if task.Status == strings.ToUpper(status) {
			sliceTasks = append(sliceTasks, task)
		}
	}

	return sliceTasks, nil
}

func DisplayTasksTable(tasks []Task) {
	const idWidth = 5
	const statusWidth = 10
	const maxDescriptionWidth = 60

	fmt.Printf("%-*s %-*s %s\n", idWidth, "ID", statusWidth, "STATUS", "DESCRIÇÃO")
	fmt.Println("--------------------------------------")

	for _, t := range tasks {
		displayDescription := t.Description
		if len(displayDescription) > maxDescriptionWidth {
			displayDescription = displayDescription[:maxDescriptionWidth-3] + "..."
		}

		fmt.Printf("%-*d %-*s %s\n",
			idWidth, t.ID,
			statusWidth, t.Status,
			displayDescription)
	}

	fmt.Println("--------------------------------------")
}
