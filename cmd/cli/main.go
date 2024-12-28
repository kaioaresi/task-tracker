package main

import (
	"fmt"
	"log"
	"task-tracker/internal/storage"
	"task-tracker/internal/task"
)

func main() {
	// criando uma task
	task1 := task.NewTask("clean the house")

	fmt.Println(task1)

	f1 := storage.NewFile("task1")

	if err := f1.CreateFile(); err != nil {
		log.Println(err)
	}

}
