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
func ParseJson(t task.Task) ([]byte, error) {
	j, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	}

	return j, nil
}

// CreateFile
func (f *File) CreateFile() error {

	if ok := CheckFileExists(f.Name); ok {
		log.Println("info: arquivo existe")
		return nil
	}

	file, err := os.Create(f.Name)
	if err != nil {
		return fmt.Errorf("error: can`t create file `%s`.\n %v", f.Name, err)
	}
	defer file.Close()

	log.Println("File created...", f.Name)

	return nil
}

func (f *File) ReadFile() ([]task.Task, error) {
	file, err := os.Open(f.Name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode json
	var sliceTask []task.Task
	err = json.NewDecoder(file).Decode(&sliceTask)
	if err != nil {
		return nil, err
	}

	return sliceTask, nil

}

func (f *File) WriteFile(task task.Task) error {
	file, err := os.Open(f.Name)
	if err != nil {
		return err
	}
	defer file.Close()

	slTasks, err := f.ReadFile()
	if err != nil {
		return err
	}

	slTasks = append(slTasks, task)

	fmt.Println(slTasks)

	// Encode struct Json
	b, err := json.Marshal(slTasks)
	if err != nil {
		return err
	}

	// Write a file
	_, err = file.WriteString(string(b))
	if err != nil {
		return err
	}

	return nil
}
