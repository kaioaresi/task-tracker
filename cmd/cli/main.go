package main

import (
	"log"
	"task-tracker/internal/storage"
	"task-tracker/internal/task"
)

func main() {
	// criando uma task
	task1 := task.NewTask("clean the house")

	f1 := storage.NewFile()

	if err := f1.CreateFile(); err != nil {
		log.Fatal(err)
	}

	err := f1.WriteFile(task1)
	if err != nil {
		log.Fatal(err)
	}
}
