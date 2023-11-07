package models

import (
	"testing"
	"time"
)

var result *Task;

func Test_Constructor(t *testing.T) {

	now := time.Now()

	result = Task{
		ID:          "1",
		Title:       "Test Task",
		Description: "This is a test task.",
		Deadline:    now,
		Completed:   false,
	}

	if result.ID != "1" {
		t.Errorf("Task.ID should be 1, got %s", result.ID)
	}

	if result.Title != "Test Task" {
		t.Errorf("Task.Title should be 'Test Task', got %s", result.Title)
	}

	if result.Description != "This is a test task." {
		t.Errorf("Task.Description should be 'This is a test task.', got %s", result.Description)
	}

	if result.Deadline != now {
		t.Errorf("Task.Deadline should be nil, got %s", result.Deadline)
	}

	if result.Completed != false {
		t.Errorf("Task.Completed should be false, got %t", result.Completed)
	}
}

func Test_Update_Title(t *testing.T) {
	newTitle := "new-title"
	result.UpdateTitle(newTitle)

	if result.Title != newTitle {
		t.Errorf("Task.UpdateTitle was not successful. Got %s, but should be %s", result.Title, newTitle)
	}
}