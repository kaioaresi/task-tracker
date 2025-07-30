package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
	case "help", "-h", "--help":
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

		status, err := task.Add()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(status)
	case "list":

		t := task.Task{}

		if len(os.Args) == 2 {

			tasks, err := t.List()
			if err != nil {
				log.Fatal(err)
			}
			task.DisplayTasksTable(tasks)
			return
		}

		listOptions := os.Args[2]

		fmt.Println("List option:", listOptions)
		fmt.Println("Exibindo list tasks todo")
		tasksTodo, err := t.ListTaskByStatus(listOptions)
		if err != nil {
			log.Fatal(err)
		}

		task.DisplayTasksTable(tasksTodo)

	case "update":
		if len(os.Args) < 3 {
			fmt.Println("Error: you need to info description")
			utils.Help()
			return
		}

		if len(os.Args[2]) == 0 {
			fmt.Println("Task list is empty")
			utils.Help()
			return
		}

		taskID, _ := strconv.Atoi(os.Args[2])
		taskDescription := os.Args[3]

		t := task.Task{}
		if err := t.Update(taskID, taskDescription); err != nil {
			log.Fatal(err)
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: you need to info task 1")
			utils.Help()
			return
		}

		if len(os.Args[2]) == 0 {
			utils.Help()
			return
		}

		taskID, _ := strconv.Atoi(os.Args[2])
		t := task.Task{}
		if err := t.Delete(taskID); err != nil {
			log.Fatal(err)
		}
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Error: you need to info task 1")
			utils.Help()
			return
		}

		if len(os.Args[2]) == 0 {
			utils.Help()
			return
		}
		taskID, _ := strconv.Atoi(os.Args[2])
		t := task.Task{}
		if err := t.MarkInProgress(taskID); err != nil {
			log.Fatal(err)
		}
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Error: you need to info task 1")
			utils.Help()
			return
		}

		if len(os.Args[2]) == 0 {
			utils.Help()
			return
		}
		taskID, _ := strconv.Atoi(os.Args[2])
		t := task.Task{}
		if err := t.MarkDone(taskID); err != nil {
			log.Fatal(err)
		}

	default:
		fmt.Printf("Invalide option, %q\n", argTask)
		utils.Help()
	}
}
