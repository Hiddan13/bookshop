package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db *gorm.DB
)

func Connect() {
	db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
	defer db.Close()

	if err != nil {
		fmt.Println(err, "Failed to connect to database")
	}
}

func GetDB() *gorm.DB {
	return db
}
