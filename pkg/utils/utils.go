package utils

import (
	"fmt"
)

func Help() {
	fmt.Println("Options:")
	fmt.Println("add: add new task")
	fmt.Println("\tusage - cli add 'task description'")
}

func ErrorF(msg string, err error) error {
	if err != nil {
		return fmt.Errorf(msg, err)
	}
	return nil
}
