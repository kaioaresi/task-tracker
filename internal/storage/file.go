package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"task-tracker/internal/task"
)

type FileStorage struct {
	Name string
}

// Check if file existe
func CheckFileExist(fileName string) (bool, error) {
	if _, err := os.Stat(fileName); err == nil {
        return true, fmt.Errorf("file exist '%s'", fileName)
	}

	return false, nil
}

// Create a file
func Newfile() (*FileStorage, error){

	fileName := "task.json"

	if ok, err := CheckFileExist(fileName); ok {
		log.Println(err)
		return nil, nil
	}

	file, err := os.Create(fileName)
	if err != nil {
		return nil, fmt.Errorf("error to create file %e", err)
	}
	defer file.Close()
	
	log.Println("Create file...")

	return &FileStorage{
		Name:      fileName,
	}, nil
}

// Read file
func (f *FileStorage) Read () (task.Task,error) {
	log.Println("Start read file.....")

	file, err := os.Open("task.json")
	if err != nil {
		return nil, fmt.Errorf("erro to open file %f", err)
	}
	log.Println("Successfully Opened users.json")
	defer file.Close()

	// Parsing with Struct
	bFile, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("erro to read file %f", err)
	}

	var tasks task.Task

	err = json.Unmarshal(bFile, &tasks)
	if err != nil {
		return nil, fmt.Errorf("error to parsing file %e", err)
	}

	log.Println("Finished read file.....")
	
	return tasks, nil
}
