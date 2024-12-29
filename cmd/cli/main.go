package main

import (
	"log"
	"task-tracker/internal/storage"
)

func main() {

	// Create file
	f1 := storage.NewFile()
	err := f1.CreateFile()
	if err != nil {
		log.Fatal(err)
	}

	// Read file
	// jsonFile, err := f1.ReadFile()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Atual: ", jsonFile)

	// criando uma task
	// task1 := task.NewTask("clean the house")
	// task2 := task.NewTask("study")

	// // jsonFile = append(jsonFile, *task1, *task2)

	// // fmt.Println(len(jsonFile))

	// err = f1.WriteFile(*task1)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
