package main

import (
	"fmt"
	"log"
	"task-tracker/internal/cli"
	"task-tracker/internal/task"
)

func main() {
	description := cli.Create()

	if description == "" {
		log.Fatal("Nenhum valor foi informado, utilize o -h para ver as opções disponíveis!")
	}

	t1 := task.NewTask(description)

	fmt.Println(t1)

	t1.UpdateStatus()

	fmt.Println(t1)

	t1.Done()

	fmt.Println(t1)

}
