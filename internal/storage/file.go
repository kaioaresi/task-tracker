package storage

import (
	"encoding/json"
	"fmt"
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

// Newfile - create a file
func Newfile() (*FileStorage, error) {

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
		Name: fileName,
	}, nil
}

// Save - save task on file
func (f *FileStorage) Save(t task.Task) error {
	log.Print("Start saving task.....")

	file, err := os.OpenFile("task.json", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("error to open file!\n%v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(t)
	if err != nil {
		return fmt.Errorf("error to encode task!\n%v", err)
	}

	log.Print("task saved!")

	return nil
}

// Read - read file
func (f *FileStorage) Read() (*task.Task, error) {
	log.Println("Start read file.....")

	file, err := os.Open("task.json")
	if err != nil {
		return nil, fmt.Errorf("erro to open file %f", err)
	}
	log.Println("Successfully Opened task.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	tasks := task.Task{}

	err = decoder.Decode(&tasks)
	if err != nil {
		return nil, fmt.Errorf("error to decode file\n%v", err)
	}

	log.Println("Finished read file.....")

	return &tasks, nil
}

