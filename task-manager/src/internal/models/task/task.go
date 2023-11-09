package task

import (
	"fmt"
	"time"
)

// タスクの構造体
type Task struct {
	Id          int `pg:",pk" json:"id"`
	Title       string
	Description string
	Deadline    time.Time
	Completed   bool
}

// タスクが文字列で返ってくる
func (t *Task) String() string {
	return fmt.Sprintf(
		"Task: %s | Description: %s | Deadline: %s | Completed: %t",
		t.Title, t.Description, t.Deadline, t.Completed,
	)
}
