package main

import (
	"fmt"
	"log"
	"task-tracker/internal/storage"
)

func main() {
	file, err := storage.Newfile()
	if err != nil {
		log.Println(err)
	}

	// Read file
	tasks, err := file.Read()
	if err != nil {
		log.Println(err)
	}

	for _, task := range tasks {
		fmt.Println(task.Description)
	}

}