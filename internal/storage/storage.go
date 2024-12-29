package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task-tracker/internal/task"
)

type File struct {
	Name string `json:"name"`
}

// Construtor
func NewFile() *File {
	return &File{
		Name: "task.json",
	}
}

// CheckFileExists - exists
func CheckFileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// ParseToJson
func ParseJson(t *task.Task) (string, error) {
	j, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return string(j), nil
}

// CreateFile
func (f *File) CreateFile() error {

	ok := CheckFileExists(f.Name)
	if ok {
		log.Println("info: arquivo existe")
		return nil
	}

	var err error
	_, err = os.Create(f.Name)
	if err != nil {
		return fmt.Errorf("error: can`t create file `%s`.\n %v", f.Name, err)
	}

	log.Println("File create...", f.Name)

	return nil
}

func (f *File) WriteFile(t *task.Task) error {
	file, err := os.OpenFile(f.Name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	textJson, err := ParseJson(t)
	if err != nil {
		return err
	}

	_, err = file.WriteString(textJson)
	if err != nil {
		return err
	}

	log.Println("File writing file...", f.Name)
	return nil
}

func (f *File) ReadFile() ([]byte, error) {
	data, err := os.ReadFile(f.Name)
	if err != nil {
		return nil, err
	}
	return data, nil
}
