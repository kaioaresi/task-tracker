package storage

import (
	"fmt"
	"os"
	"task-tracker/pkg/utils"
)

func CreateFile(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, utils.ErrorF("failed to open file", err)
	}
	return file, nil
}

func CheckFile(fileName string) error {
	if _, err := os.Stat(fileName); err != nil {
		if !os.IsNotExist(err) {
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
	if err != nil {
		return nil, utils.ErrorF("error to save task on file", err)
	}

	return file, nil
}

func Read(fileName string) ([]byte, error) {
	bData, err := os.ReadFile(fileName)
	if err != nil {
		return nil, utils.ErrorF(fmt.Sprintf("error could not read file '%s'", fileName), err)
	}

	return bData, nil
}

func Write(fileName string, data []byte) error {
	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		return utils.ErrorF(fmt.Sprintf("failed to write a file '%s'", fileName), err)
	}

	return nil
}
