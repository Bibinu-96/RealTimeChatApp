package dao

import (
	"fmt"
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
		sqlDB, err := db.DB()
		if err != nil {
			fmt.Println("Error getting *sql.DB object:", err)
			return
		}

		// Configure connection pooling
		sqlDB.SetMaxOpenConns(5)    // Maximum number of open connections
		sqlDB.SetMaxIdleConns(5)    // Maximum number of idle connections
		sqlDB.SetConnMaxLifetime(0) // Maximum amount of time a connection can be reused (0 means no limit)

		// Use the `db` object to interact with your database

		// Close the connection when done
		//defer sqlDB.Close()

		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}
		log.Println("Database connection established")
	})
	return db
}
