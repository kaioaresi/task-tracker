package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"task-tracker/pkg/storage"
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

const (
	fileName   = "/tmp/tasks.json"
	INPROGRESS = "IN-PROGRESS"
	DONE       = "DONE"
	TODO       = "TODO"
)

func NewTask(description string) *Task {

	return &Task{
		Description: description,
		Status:      TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}
}

func (t *Task) Add() error {
	sliceTasks, err := t.Read()
	if os.IsNotExist(err) {
		return utils.ErrorF(fmt.Sprintf("failed to read existing tasks to add new task: %s", fileName), err)
	}

	t.ID = getMaxID(sliceTasks) + 1
	sliceTasks = append(sliceTasks, *t)

	err = t.Save(sliceTasks)
	if err != nil {
		return utils.ErrorF(fmt.Sprintf("failed to save new task (ID: %d)", t.ID), err)
	}

	log.Println("Task added successfully (ID:)", t.ID)

	return nil
}

func (t *Task) Read() ([]Task, error) {

	bData, err := storage.Read(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, utils.ErrorF(fmt.Sprintf("failed to read task file '%s'", fileName), err)
	}

	if len(bData) == 0 {
		return []Task{}, nil
	}

	var tasks []Task
	err = json.Unmarshal(bData, &tasks)
	if err != nil {
		return nil, utils.ErrorF(fmt.Sprintf("failed to Unmarshal file '%s'", fileName), err)
	}

	return tasks, nil
}

func (t *Task) Save(tasks []Task) error {
	jsonData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return utils.ErrorF(fmt.Sprintf("failed to marshal data to save in file '%s'", fileName), err)
	}

	err = storage.Write(fileName, jsonData)
	if err != nil {
		return utils.ErrorF(fmt.Sprintf("failed to save task ID (%d) in file %s", t.ID, fileName), err)
	}
	return nil
}

func (t Task) List() ([]Task, error) {
	tasks, err := t.Read()
	if errors.Is(err, errors.New("no such file or directory")) {
		return nil, utils.ErrorF(fmt.Sprintf("failed to list tasks in file '%s'", fileName), err)
	}
	return tasks, nil
}

func (t *Task) Update(taskID int, description string) error {
	return t.updateTask(taskID, func(tasks []Task, index int) ([]Task, error) {
		tasks[index].Description = description
		tasks[index].UpdatedAt = time.Now()
		return tasks, nil
	})
}

func (t *Task) updateTask(taskID int, modify func(tasks []Task, index int) ([]Task, error)) error {
	tasks, err := t.Read()
	if err != nil {
		return utils.ErrorF(fmt.Sprintf("failed to read file '%s' on update task", fileName), err)
	}

	index := -1
	for i, task := range tasks {
		if task.ID == taskID {
			index = i
			break
		}
	}

	if index == -1 {
		return utils.ErrorF(fmt.Sprintf("task %d not found on file '%s'", taskID, fileName), err)
	}

	modifiedTasks, err := modify(tasks, index)
	if err != nil {
		return utils.ErrorF(fmt.Sprintf("failed to modify task %d", taskID), err)
	}

	if err := t.Save(modifiedTasks); err != nil {
		return utils.ErrorF(fmt.Sprintf("failed to save update task %d in file '%s'", taskID, fileName), err)
	}

	return nil
}

func (t *Task) Delete(taskID int) error {
	tasks, err := t.Read()
	if err != nil {
		return utils.ErrorF(fmt.Sprintf("failed to delete task id '%d' in file '%s'", taskID, fileName), err)
	}

	indexDelete := -1
	for i, task := range tasks {
		if task.ID == taskID {
			indexDelete = i
			break
		}
	}

	if indexDelete == -1 {
		return utils.ErrorF(fmt.Sprintf("task %d not found on file '%s'", taskID, fileName), err)
	}

	tasks = append(tasks[:indexDelete], tasks[indexDelete+1:]...)
	err = t.Save(tasks)
	if err != nil {
		return utils.ErrorF(fmt.Sprintf("failed to delete task %d on delete workload in file '%s'", taskID, fileName), err)
	}

	return nil
}

func (t *Task) ChangeStatus(taskID int, status string) error {
	t.Status = status
	t.UpdatedAt = time.Now()

	tasks, err := t.Read()
	if err != nil {
		return utils.ErrorF(fmt.Sprintf("failed to read file '%s' in change status task %d", fileName, taskID), err)
	}

	indexNotFount := -1
	for i := range tasks {
		if tasks[i].ID == taskID {
			tasks[i].Status = t.Status
			indexNotFount = i
			break
		}
	}

	if indexNotFount == -1 {
		return utils.ErrorF(fmt.Sprintf("failed task %d not found in file '%s'", taskID, fileName), err)
	}

	if err := t.Save(tasks); err != nil {
		return utils.ErrorF(fmt.Sprintf("failed to save status change task %d", taskID), err)
	}

	return nil
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

func (t Task) ListTaskByStatus(status string) ([]Task, error) {
	tasks, err := t.Read()
	if err != nil {
		return nil, utils.ErrorF(fmt.Sprintf("failed to list tasks by status in file '%s'", fileName), err)
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

	if len(tasks) == 0 {
		log.Println("No tasks found. Add a new task with 'cli add <task description>'.")
		return
	}

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
