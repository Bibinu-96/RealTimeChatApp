package dao

import (
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// GetDB provides a singleton database connection
func GetDB() *gorm.DB {
	once.Do(func() {
		dsn := "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable"
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}
		log.Println("Database connection established")
	})
	return db
}
