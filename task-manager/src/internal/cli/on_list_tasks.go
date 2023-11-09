package cli

import (
	"fmt"
	"task-manager/internal/repo"
)

func onListTasks(repo repo.TaskRepo) {
	fmt.Println("Listing tasks...")
	tasks := repo.GetTasks()
	for _, task := range tasks {
		fmt.Println(task)
	}
}
