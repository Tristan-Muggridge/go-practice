package cli

import (
	"fmt"
	"log"
	"strconv"
	"task-manager/internal/repo"
)

func onDeleteTask(repo repo.TaskRepo) {
	fmt.Println("Please enter the ID of the task you want to delete:")
	id := captureInput()
	id_int, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	task, err := repo.GetTaskById(id_int)
	if err != nil {
		log.Fatalln(err)
	}

	if task == nil {
		fmt.Println("Task not found")
		return
	}

	fmt.Println("Deleting task...")
	err = repo.DeleteTask(task)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Task deleted")
}
