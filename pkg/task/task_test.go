package task

import (
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	task := NewTask("Test Task")
	_, err := task.Add()
	if err != nil {
		t.Errorf("Error adding task: %v", err)
	}

	tasks, err := task.Read()
	if err != nil {
		t.Errorf("Error reading tasks: %v", err)
	}

	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}

	if tasks[0].Description != "Test Task" {
		t.Errorf("Expected task description 'Test Task', got '%s'", tasks[0].Description)
	}

	// Clean up the tasks.json file
	if err := os.Remove(fileName); err != nil {
		t.Errorf("Error cleaning up tasks.json file: %v", err)
	}
}