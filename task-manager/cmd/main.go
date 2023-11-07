package main

import (
	"fmt"
	"task-manager/internal/repo"
)

func main() {
	// データベース接続
	db := repo.ConnectDb()
	defer db.Close()

	// タスクの全てが配列で返ってくる
	tasks := repo.GetTasks()

	// それぞれのタスクを表示
	for _, task := range tasks {
		fmt.Println(task)
	}
}
