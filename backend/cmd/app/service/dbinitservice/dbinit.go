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

	err := dbinit.initDB()
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

func (dbinit DBinitService) initDB() error {
	dbinit.Log.Debug("initialising db")
	db, err := dbinit.GenericDb.GetDB()
	if err != nil {
		return err
	}

	dao.SetDb(db)

	return nil

}

func (dbinit DBinitService) Run() error {
	return dbinit.createorUpdateModelstoTables()

}
func (dbinit DBinitService) GetName() string {
	return dbinit.Name

}
