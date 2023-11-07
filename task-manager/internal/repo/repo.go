package repo

import (
	"task-manager/internal/models"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var db *pg.DB

// データベース接続、db参照を返す
func ConnectDb() *pg.DB {
	db = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "Tasks",
	})
	return db
}

// テーブル作成
func CreateSchema(db *pg.DB) error {
	models := []interface{}{
		(*models.Task)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// タスク作成
func CreateTask(task *models.Task) error {
	_, err := db.Model(task).Insert()
	return err
}

// タスク取得、配列参照が本物より能率的になる
func GetTasks() []*models.Task {
	var tasks []*models.Task
	err := db.Model(&tasks).Select()
	if err != nil {
		panic(err)
	}
	return tasks
}
