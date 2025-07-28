package file

import (
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
		return nil, err
	}

	return file, nil
}
