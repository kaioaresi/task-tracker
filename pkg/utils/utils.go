package utils

import (
	"fmt"
)

func Help() {
	fmt.Println("Options:")
	fmt.Println("  add: add new task")
	fmt.Println("    usage: cli add 'task description'")
	fmt.Println("  list: list all tasks")
	fmt.Println("    usage: cli list")
	fmt.Println("    usage: cli list <status(todo|in-progress)>")
	fmt.Println("  update: update a task description")
	fmt.Println("    usage: cli update <taskID> 'new task description'")
	fmt.Println("  delete: delete a task")
	fmt.Println("    usage: cli delete <taskID>")
	fmt.Println("  mark-in-progress: update task status to 'in-progress'")
	fmt.Println("    usage: cli mark-in-progress <taskID>")
	fmt.Println("  mark-done: update task status to 'done'")
	fmt.Println("    usage: cli mark-done <taskID>")
}

func ErrorF(msg string, err error) error {
	if err != nil {
		return fmt.Errorf("%s %w", msg, err)
	}
	return nil
}

func Error(err error) error {
	return err
}
