package database

import (
	"devcode-todolist/entities"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	connection *gorm.DB
	once       sync.Once
}

func (database *Database) lazyInit() {
	database.once.Do(func() {
		host := os.Getenv("MYSQL_HOST")
		port := os.Getenv("MYSQL_PORT")
		username := os.Getenv("MYSQL_USER")
		password := os.Getenv("MYSQL_PASSWORD")
		dbname := os.Getenv("MYSQL_DBNAME")
		dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("Cannot connect database")
		}

		db.AutoMigrate(
			&entities.Activity{},
			&entities.Todo{},
		)

		database.connection = db
	})
}

func (database *Database) GetConnection() *gorm.DB {
	database.lazyInit()
	return database.connection
}

var DB = &Database{}
