package dao

import (
	"sync"

	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func SetDB(database *gorm.DB) {
	db = database
}

// GetDB provides a singleton database connection
func GetDB() *gorm.DB {

	return db
}
