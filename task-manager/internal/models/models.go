package models

import "time"

type Task struct {
	ID          string
	Title       string
	Description string
	Deadline    time.Time
	Completed   bool
}

// Making use of pointers for efficiency
func UpdateTitle(t *Task, title string) {
	t.Title = title
}

func UpdateDescription(t *Task, description string) {
	t.Description = description
}

func UpdateDeadline(t *Task, deadline time.Time) {
	t.Deadline = deadline
}

func CompleteTask(t *Task) {
	t.Completed = true
}
