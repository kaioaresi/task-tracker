package utils

import (
	"fmt"
)

func Help() {
	fmt.Println("Options:")
	fmt.Println("add: add new task")
	fmt.Println("\tusage - cli add 'task description'")
	fmt.Println("list: list all tasks")
	fmt.Println("update: update a task description")
	fmt.Println("\tusage - cli update taskID 'new task description'")

}

func ErrorF(msg string, err error) error {
	if err != nil {
		return fmt.Errorf(msg, err)
	}
	return nil
}
