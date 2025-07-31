package main

import (
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
	t := task.Task{}

	switch argTask {
	case "help", "-h", "--help":
		utils.Help()
	case "add":
		if err := utils.CheckInput(os.Args); err != nil {
			log.Fatal(err)
		}

		task := task.NewTask(os.Args[2])

		if err := task.Add(); err != nil {
			log.Fatalln(err)
		}
	case "list":
		if len(os.Args) == 2 {

			tasks, err := t.List()
			if err != nil {
				log.Fatal(err)
			}
			task.DisplayTasksTable(tasks)
			return
		}

		taskByStatus, err := t.ListTaskByStatus(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		task.DisplayTasksTable(taskByStatus)

	case "update":
		if len(os.Args) < 3 {
			log.Println("Error: you need to info description")
			utils.Help()
			return
		}

		if len(os.Args[2]) == 0 || len(os.Args[3]) == 0 {
			log.Println("Task list is empty")
			utils.Help()
			return
		}

		taskID, err := utils.InputToInt(os.Args[2])
		if err != nil {
			log.Fatalln(err)
		}

		if err := t.Update(taskID, os.Args[3]); err != nil {
			log.Fatalln(err)
		}
	case "delete":
		if err := utils.CheckInput(os.Args); err != nil {
			log.Fatalln(err)
		}

		taskID, err := utils.InputToInt(os.Args[2])
		if err != nil {
			log.Fatalln(err)
		}
		if err := t.Delete(taskID); err != nil {
			log.Fatalln(err)
		}
	case "mark-in-progress":
		if err := utils.CheckInput(os.Args); err != nil {
			log.Fatalln(err)
		}
		taskID, err := utils.InputToInt(os.Args[2])
		if err != nil {
			log.Fatalln(err)
		}
		if err := t.ChangeStatus(taskID, task.INPROGRESS); err != nil {
			log.Fatalln(err)
		}
	case "mark-done":
		if err := utils.CheckInput(os.Args); err != nil {
			log.Fatalln(err)
		}

		taskID, err := utils.InputToInt(os.Args[2])
		if err != nil {
			log.Fatalln(err)
		}

		if err := t.ChangeStatus(taskID, task.DONE); err != nil {
			log.Fatal(err)
		}

	default:
		log.Fatalf("Invalide option, %q\n", argTask)
		utils.Help()
	}
}
