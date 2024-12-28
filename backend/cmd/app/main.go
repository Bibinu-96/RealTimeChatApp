package main

import (
	"backend/cmd/app/service"

	"sync"
	"time"

	"backend/cmd/app/service/dbinitservice"
	"backend/cmd/app/service/server"
	"backend/cmd/app/service/server/config"
	"backend/cmd/app/service/server/router"
	"backend/internal/database/database"
	"backend/pkg/logger"
)

func main() {
	var log logger.Logger
	var wg sync.WaitGroup
	components := []service.Service{}

	// Create a new logger instance
	log = logger.NewLogrusLogger()

	// db init service
	var dbInitService service.Service
	dbInitService = dbinitservice.DBinitService{
		Log:       log,
		Name:      "DBInitService",
		GenericDb: database.PostgressDB{DSN: ""},
	}

	components = append(components, dbInitService)
	//var service1 service.Service

	// run gin server
	var service1 service.Service
	serverConfig1 := config.ServerConfig{
		Addr:         ":8000",
		Router:       router.SetupGinRouter(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	service1 = server.NewServer(serverConfig1, "GinApiServer", log)

	components = append(components, service1)

	//second service
	var service2 service.Service
	serverConfig2 := config.ServerConfig{
		Addr:         ":8001",
		Router:       router.SetupChiRouter(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	service2 = server.NewServer(serverConfig2, "ChiApiServer", log)

	components = append(components, service2)

	for _, component := range components {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := component.Run()
			if err != nil {
				log.Error("Error occured", err, component.GetName())
			}
		}()

	}

	wg.Wait()
}
