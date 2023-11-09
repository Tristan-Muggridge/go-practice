package repo

import (
	"task-manager/internal/models/task"
)

type TaskRepo interface {
	CreateTask(task *task.Task) error
	GetTasks() []*task.Task
	GetTaskById(id int) (*task.Task, error)
	UpdateTask(task *task.Task) error
	DeleteTask(task *task.Task) error
	GetWhere(where string, params ...interface{}) ([]*task.Task, error)
}
