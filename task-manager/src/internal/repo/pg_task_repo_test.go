package repo_test

import (
	"task-manager/db"
	"task-manager/internal/models/task"
	"task-manager/internal/repo"
	"testing"
	"time"

	"github.com/go-pg/pg/v10"
)

func truncate_tasks(db *pg.DB) {
	_, err := db.Exec("TRUNCATE TABLE tasks")
	if err != nil {
		panic(err)
	}

	// reset sequence
	_, err = db.Exec("ALTER SEQUENCE tasks_id_seq RESTART WITH 1")
	if err != nil {
		panic(err)
	}
}

func connect(options ...string) *pg.DB {
	db := db.Connect(&db.DatabaseCredentials{
		User:     "postgres",
		Password: "postgres",
		Database: "task_manager_test",
		Port:     "8080",
	})

	if len(options) != 0 && options[0] == "--truncate" {
		truncate_tasks(db)
	}

	return db
}

var DB *pg.DB = connect()

func Test_Create_Task(t *testing.T) {
	DB = connect("--truncate")
	db := DB

	taskRepo := repo.NewPgTaskRepo(db)

	// Create a task
	task := &task.Task{
		Title:       "Test",
		Description: "Test",
		Deadline:    time.Now(),
		Completed:   false,
	}

	err := taskRepo.CreateTask(task)
	if err != nil {
		t.Fatalf("Error creating task: %v\n", err)
	}

	// Get the task
	task, err = taskRepo.GetTaskById(task.Id)
	if err != nil {
		t.Fatalf("Error getting task: %v\n", err)
	}

	// Check that the task is correct
	if task.Id != 1 {
		t.Fatalf("Expected Id to be 1, got %d\n", task.Id)
	}

	if task.Title != "Test" {
		t.Fatalf("Expected Title to be Test, got %s\n", task.Title)
	}

	if task.Description != "Test" {
		t.Fatalf("Expected Description to be Test, got %s\n", task.Description)
	}

	if task.Completed != false {
		t.Fatalf("Expected Completed to be false, got %t\n", task.Completed)
	}
}

func Test_Get_Tasks(t *testing.T) {
	db := DB

	taskRepo := repo.NewPgTaskRepo(db)

	// Get Tasks
	tasks := taskRepo.GetTasks()
	if len(tasks) != 1 {
		t.Fatalf("Expected 1 task, got %d\n", len(tasks))
	}

	// Check that the task is correct
	if tasks[0].Id != 1 {
		t.Fatalf("Expected Id to be 1, got %d\n", tasks[0].Id)
	}
}

func Test_Update_Task(t *testing.T) {
	db := DB

	taskRepo := repo.NewPgTaskRepo(db)

	// Get the task
	task, err := taskRepo.GetTaskById(1)
	if err != nil {
		t.Fatalf("Error getting task: %v\n", err)
	}

	// Update the task
	task.Title = "Updated"

	err = taskRepo.UpdateTask(task)
	if err != nil {
		t.Fatalf("Error updating task: %v\n", err)
	}

	// Get the task
	task, err = taskRepo.GetTaskById(task.Id)
	if err != nil {
		t.Fatalf("Error getting task: %v\n", err)
	}

	// Check that the task is correct
	if task.Title != "Updated" {
		t.Fatalf("Expected Title to be Updated, got %s\n", task.Title)
	}
}

func Test_Get_Where(t *testing.T) {
	db := DB

	taskRepo := repo.NewPgTaskRepo(db)

	// Get Tasks
	tasks, err := taskRepo.GetWhere("title = ?", "Updated")
	if err != nil {
		t.Fatalf("Error getting tasks: %v\n", err)
	}

	if len(tasks) != 1 {
		t.Fatalf("Expected 1 task, got %d\n", len(tasks))
	}

	if tasks[0].Title != "Updated" {
		t.Fatalf("Expected Title to be Updated, got %s\n", tasks[0].Title)
	}
}

func Test_Delete_Task(t *testing.T) {
	db := DB

	taskRepo := repo.NewPgTaskRepo(db)

	// Get the task
	task, err := taskRepo.GetTaskById(1)
	if err != nil {
		t.Fatalf("Error getting task: %v\n", err)
	}

	// Delete the task
	err = taskRepo.DeleteTask(task)
	if err != nil {
		t.Fatalf("Error deleting task: %v\n", err)
	}

	// Get the task
	task, err = taskRepo.GetTaskById(task.Id)
	if err == nil {
		t.Fatalf("Expected error getting task, got %v\n", task)
	}
}
