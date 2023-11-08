package repo

import (
	"task-manager/internal/models/task"

	"github.com/go-pg/pg/v10"
)

type PgTaskRepo struct {
	db *pg.DB
}

func NewPgTaskRepo(db *pg.DB) *PgTaskRepo {
	return &PgTaskRepo{db}
}

func (repo *PgTaskRepo) CreateTask(task *task.Task) error {
	_, err := repo.db.Model(task).Insert()
	return err
}

func (repo *PgTaskRepo) GetTasks() []*task.Task {
	var tasks []*task.Task
	repo.db.Model(&tasks).Select()
	return tasks
}

func (repo *PgTaskRepo) GetTaskById(id int64) (*task.Task, error) {
	task := &task.Task{Id: id}
	err := repo.db.Model(task).WherePK().Select()
	return task, err
}

func (repo *PgTaskRepo) UpdateTask(task *task.Task) error {
	_, err := repo.db.Model(task).WherePK().Update()
	return err
}

func (repo *PgTaskRepo) DeleteTask(task *task.Task) error {
	_, err := repo.db.Model(task).WherePK().Delete()
	return err
}

func (repo *PgTaskRepo) GetWhere(where string, params ...interface{}) ([]*task.Task, error) {
	var tasks []*task.Task
	err := repo.db.Model(&tasks).Where(where, params...).Select()
	return tasks, err
}
