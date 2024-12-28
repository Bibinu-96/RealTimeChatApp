package dbinitservice

import (
	"backend/cmd/app/service/logger"
	"backend/internal/database/dao"
	"backend/internal/database/database"
	"backend/internal/database/models"
)

type DBinitService struct {
	Name      string
	Log       logger.Logger
	GenericDb database.Database
}

func (dbinit DBinitService) createorUpdateModelstoTables() error {

	dbinit.Log.Debug("initialising db")

	// var genericDB database.Database
	// dsn := "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable"
	// genericDB = database.PostgressDB{DSN: dsn}

	err := dao.SetDb(dbinit.GenericDb)
	if err != nil {
		return err
	}

	db := dao.GetDB()
	err = db.AutoMigrate(&models.User{}, &models.Group{}, &models.Message{}, &models.GroupMember{})
	if err != nil {
		dbinit.Log.Error("error creating or updating db")
	}
	return err
}

func (dbinit DBinitService) Run() error {
	return dbinit.createorUpdateModelstoTables()

}
func (dbinit DBinitService) GetName() string {
	return dbinit.Name

}
