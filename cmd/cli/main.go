package main

import (
	"log"
	"task-tracker/internal/storage"
)

func main() {
	_, err := storage.Newfile()
	if err != nil {
		log.Fatal(err)
	}
}