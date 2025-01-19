package storage

import (
	"fmt"
	"log"
	"os"
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
		return nil, err
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

