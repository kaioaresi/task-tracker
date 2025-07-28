package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task-tracker/pkg/storage/file"
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

		status, err := task.AddTask()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(status)
	case "list":
		bData, err := file.ReadFile("tasks.json")
		if err != nil {
			log.Fatalf("Error to read task file", err)
			return
		}

		if len(bData) == 0 {
			return
		}

		var tasks []task.Task
		err = json.Unmarshal(bData, &tasks)
		if err != nil {
			log.Fatalf("Error to Unmarshal json file", err)
		}

		for _, task := range tasks {
			fmt.Println("ID", task.ID)
			fmt.Println("Description", task.Description)
		}

	default:
		fmt.Printf("Invalide option, %q\n", argTask)
		utils.Help()
	}
}
