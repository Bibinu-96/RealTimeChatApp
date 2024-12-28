package dao

import (
	"sync"

	"backend/internal/database/database"

	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func SetDb(database database.Database) error {
	var err error
	db, err = database.GetDB()
	if err != nil {
		return err
	}
	return nil

}

// GetDB provides a singleton database connection
func GetDB() *gorm.DB {

	return db
}
