package dbinitservice

import (
	"backend/internal/database/dao"
	"backend/internal/database/database"
	"backend/internal/database/models"
	"backend/pkg/logger"
	"context"
)

type DBinitService struct {
	Name      string
	Log       logger.Logger
	GenericDb database.Database
}

func (dbinit DBinitService) createorUpdateModelstoTables() error {

	db := dao.GetDB()
	// // Create the enum type first
	// err := db.Exec("CREATE TYPE message_type_enum AS ENUM ('direct', 'group')").Error
	// if err != nil {
	// 	return err
	// }

	// Migrate in order of dependencies
	err := db.AutoMigrate(
		&models.User{},  // 1. Users
		&models.Group{}, // 2. Groups
		&models.Message{},
	)
	if err != nil {
		return err
	}

	dbinit.Log.Info("Database migration completed successfully!")
	return err
}

func (dbinit DBinitService) initDB() error {
	dbinit.Log.Debug("initialising db")
	db, err := dbinit.GenericDb.GetDB()
	if err != nil {
		return err
	}

	dao.SetDb(db)

	return nil

}

func (dbinit DBinitService) Run(ctx context.Context) error {
	err := dbinit.initDB()
	if err != nil {
		return err
	}
	return dbinit.createorUpdateModelstoTables()

}
func (dbinit DBinitService) GetName() string {
	return dbinit.Name

}
