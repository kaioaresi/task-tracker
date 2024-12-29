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

	// Create file
	f1 := storage.NewFile()

	if err := f1.CreateFile(); err != nil {
		log.Fatal(err)
	}

	// Write file
	err := f1.WriteFile(task1)
	if err != nil {
		log.Fatal(err)
	}

	// Read file
	data, err := f1.ReadFile()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

}
