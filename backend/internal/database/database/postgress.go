package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type PostgressDB struct {
	DSN string
}

func (postgressDb PostgressDB) GetDB() (*gorm.DB, error) {
	var err error
	db, err := gorm.Open(postgres.Open(postgressDb.DSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // Use plural table names
		},
	})
	if err != nil {
		fmt.Println("Error getting *sql.DB object:", err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Error getting *sql.DB object:", err)
		return nil, err
	}
	// Configure connection pooling
	sqlDB.SetMaxOpenConns(5)    // Maximum number of open connections
	sqlDB.SetMaxIdleConns(5)    // Maximum number of idle connections
	sqlDB.SetConnMaxLifetime(0) // Maximum amount of time a connection can be reused (0 means no limit)

	// Use the `db` object to interact with your database

	// Close the connection when done
	//defer sqlDB.Close()

	return db, nil
}
