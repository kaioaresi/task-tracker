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

	// // Write file
	// t1 := task.Task{
	// 	ID:          4,
	// 	Description: "study japones",
	// 	Status:      "TODO",
	// 	CreateAt:    time.Now(),
	// 	DeleteAt:    time.Now(),
	// 	UpdateAt:    time.Now(),
	// }
	// t2 := task.Task{
	// 	ID:          5,
	// 	Description: "study Go",
	// 	Status:      "TODO",
	// 	CreateAt:    time.Now(),
	// 	DeleteAt:    time.Now(),
	// 	UpdateAt:    time.Now(),
	// }

	// err = file.Save(t1)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// err = file.Save(t2)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// Read file
	tasks, err := file.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tasks)
}