package file

import (
	"fmt"
	"log"
	"os"
)

func CreateFile(fileName string) (*os.File, error) {
	return os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}

func ProvideFile(fileName string) (*os.File, error) {
	if _, err := os.Stat(fileName); err != nil {
		if !os.IsNotExist(err) {
			log.Println("File exist")
			return nil, err
		}
	}

	file, err := CreateFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error to save task on file, %v\n", err)
	}

	return file, nil
}
