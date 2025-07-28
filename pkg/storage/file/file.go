package file

import (
	"log"
	"os"
	"task-tracker/pkg/utils"
)

func CreateFile(fileName string) (*os.File, error) {
	return os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
}

func CheckFile(fileName string) error {
	if _, err := os.Stat(fileName); err != nil {
		if !os.IsNotExist(err) {
			log.Println("File exist")
			return utils.ErrorF("Cannot check if file existe!", err)
		}
	}

	return nil
}

func ProvideFile(fileName string) (*os.File, error) {
	err := CheckFile(fileName)
	if err != nil {
		return nil, err
	}

	file, err := CreateFile(fileName)
	log.Println("Create file...")
	if err != nil {
		return nil, utils.ErrorF("error to save task on file", err)
	}

	return file, nil
}

func ReadFile(fileName string) ([]byte, error) {
	bData, err := os.ReadFile(fileName)
	if err != nil {
		return nil, utils.ErrorF("error could not read file", err)
	}

	return bData, nil
}
