package main

import (
	"task-manager/db"
	"task-manager/internal/cli"
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

	cli.NewCli(taskRepo)
}
