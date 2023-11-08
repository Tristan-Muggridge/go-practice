package main

import (
	"fmt"
	"task-manager/db"
	"task-manager/internal/models/task"
	"task-manager/internal/repo"
)

func main() {
	db := db.Connect(&db.DatabaseCredentials{
		User:     "postgres",
		Password: "postgres",
		Database: "task_manager",
		Port:     "5432",
	})
	defer db.Close()

	taskRepo := repo.NewPgTaskRepo(db)

	task := &task.Task{
		Id:          1,
		Title:       "Test",
		Description: "Test",
	}

	err := taskRepo.CreateTask(task)
	if err != nil {
		fmt.Printf("Error creating task: %v\n", err)
	}

	tasks := taskRepo.GetTasks()
	for _, task := range tasks {
		fmt.Printf("Task: %v\n", task)
	}
}
