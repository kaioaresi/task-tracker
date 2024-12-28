package storage

import (
	"fmt"
	"log"
	"os"
)

type File struct {
	Name string
	Info *os.File
}

// Construtor
func NewFile(name string) *File {
	return &File{
		Name: name,
	}
}

// CheckFileExists - exists
func CheckFileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// CreateFile
func (f *File) CreateFile() error {

	ok := CheckFileExists(f.Name)
	if !ok {
		log.Println("info: arquivo existe")
		return nil
	}

	var err error
	f.Info, err = os.Create(f.Name + ".json")
	if err != nil {
		return fmt.Errorf("error: can`t create file `%s`.\n %v", f.Name, err)
	}

	log.Println("File create...", f.Name)

	return nil
}

func (f *File) WriteFile(msg string) error {
	return nil
}
