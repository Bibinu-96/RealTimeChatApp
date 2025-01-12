package postgres

import (
	"backend/internal/database/models"
	"backend/pkg/logger"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type PostgresDb struct {
	DSN string
	Log logger.Logger
}

func (postgresDb *PostgresDb) InitDB() (*gorm.DB, error) {
	var err error
	db, err := gorm.Open(postgres.Open(postgresDb.DSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // Use plural table names
		},
	})
	if err != nil {
		postgresDb.Log.Error("err getting postgress db", err)
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

func (postgresDb *PostgresDb) RunMigrations(db *gorm.DB) error {

	// Migrate in order of dependencies
	err := db.AutoMigrate(
		&models.User{},  // 1. Users
		&models.Group{}, // 2. Groups
		&models.Message{},
		&models.UserInteraction{},
	)
	if err != nil {
		return err
	}

	postgresDb.Log.Info("Database migration completed successfully!")
	return nil

}
