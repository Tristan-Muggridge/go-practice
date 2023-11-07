package models

import (
	"fmt"
	"time"
)

// タスクの構造体
type Task struct {
	ID          string
	Title       string
	Description string
	Deadline    time.Time
	Completed   bool
}

// タスク作成
func UpdateTitle(t *Task, title string) {
	// TODO: input validation
	t.Title = title
}

// タスク作成
func UpdateDescription(t *Task, description string) {
	// TODO: input validation
	t.Description = description
}

// タスク作成
func UpdateDeadline(t *Task, deadline time.Time) {
	// TODO: input validation
	t.Deadline = deadline
}

// タスク作成
func CompleteTask(t *Task) {
	// TODO: input validation
	t.Completed = true
}

// タスクが文字列で返ってくる
func (t *Task) String() string {
	return fmt.Sprintf(
		"Task: %s\nDescription: %s\nDeadline: %s\nCompleted: %t\n",
		t.Title, t.Description, t.Deadline, t.Completed,
	)
}
