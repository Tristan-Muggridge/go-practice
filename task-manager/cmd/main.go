package main

import (
	"fmt"

	"task-manager/internal/repo"
)

func main() {
	fmt.Println("Hello, World!")

	db := repo.ConnectDb()
	defer db.Close()

	// err := repo.CreateSchema(db)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// task := &models.Task{
	// 	ID:          "first-task",
	// 	Title:       "Test Task",
	// 	Description: "This is a test task.",
	// 	Deadline:    time.Now(),
	// }

	// _, err = db.Model(task).Insert()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// print all tasks in db

	tasks := repo.GetTasks()

	for _, task := range tasks {
		fmt.Printf("Task: %s\n", task.Title)
	}
}
