package main

import (
	"fmt"
	"log"
	"os"
	"task-tracker/pkg/task"
	"task-tracker/pkg/utils"
)

func main() {

	if len(os.Args) < 2 {
		utils.Help()
		return
	}

	argTask := os.Args[1]

	switch argTask {
	case "help":
		utils.Help()
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: you need to info description")
			utils.Help()
			return
		}

		if len(os.Args[2]) == 0 {
			fmt.Println("Empty task description")
			utils.Help()
			return
		}

		taskDescription := os.Args[2]

		task := task.NewTask(taskDescription)

		status, err := task.Save()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(status)

	default:
		fmt.Printf("Invalide option, %q\n", argTask)
		utils.Help()
	}
}
