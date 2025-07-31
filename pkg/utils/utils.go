package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var insufficientParameters = errors.New("insufficient parameters")
var emptyDescription = errors.New("empty task description")

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
		return fmt.Errorf("Error: %s %w", msg, err)
	}
	return nil
}

func Error(err error) error {
	return err
}

func CheckInput(input []string) error {
	if len(input) < 3 {
		Help()
		return ErrorF("bad input", insufficientParameters)
	}
	if len(os.Args[2]) == 0 {
		Help()
		return ErrorF("bad input", emptyDescription)
	}
	return nil
}

func InputToInt(input string) (int, error) {
	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, ErrorF("failed to convert input to int", err)
	}

	return number, nil
}

func GetTaskID(args []string) (int, error) {
	if err := CheckInput(args); err != nil {
		return 0, err
	}
	taskID, err := InputToInt(args[2])
	if err != nil {
		return 0, err
	}
	return taskID, nil
}
