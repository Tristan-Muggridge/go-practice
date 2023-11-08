package task

import (
	"fmt"
	"time"
)

// タスクの構造体
type Task struct {
	Id          int64 `pg:",pk" json:"id"`
	Title       string
	Description string
	Deadline    time.Time
	Completed   bool
}

// タスクが文字列で返ってくる
func (t *Task) String() string {
	return fmt.Sprintf(
		"Task: %s\nDescription: %s\nDeadline: %s\nCompleted: %t\n",
		t.Title, t.Description, t.Deadline, t.Completed,
	)
}
