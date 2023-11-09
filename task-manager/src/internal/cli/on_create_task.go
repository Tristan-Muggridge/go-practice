package cli

import (
	"fmt"
	"log"
	"task-manager/internal/models/task"
	"task-manager/internal/repo"
)

func onCreateTask(repo repo.TaskRepo) {
	fmt.Println("Please enter a title:")
	title := captureInput()
	fmt.Println("Please enter a description:")
	description := captureInput()

	task := &task.Task{
		Title:       title,
		Description: description,
		Completed:   false,
	}

	fmt.Println("Creating task...")

	err := repo.CreateTask(task)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(task)
}
