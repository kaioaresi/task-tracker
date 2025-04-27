package storage

import "task-tracker/internal/task"

type Storage interface {
	Save(t task.Task) error
}
