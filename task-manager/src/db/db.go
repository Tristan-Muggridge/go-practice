package db

import (
	"fmt"
	"task-manager/internal/models/task"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type DatabaseCredentials struct {
	User     string
	Password string
	Database string
	Port     string
}

var db *pg.DB

func loadSchema(db *pg.DB) error {
	models := []interface{}{
		(*task.Task)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func Connect(credentials *DatabaseCredentials) *pg.DB {
	if db == nil {
		db = pg.Connect(&pg.Options{
			User:     credentials.User,
			Password: credentials.Password,
			Database: credentials.Database,
			Addr:     fmt.Sprintf("localhost:%s", credentials.Port),
		})
	}

	err := loadSchema(db)
	if err != nil {
		fmt.Printf("Error loading schema: %v\n", err)
	}

	return db
}

func Close() {
	db.Close()
}
