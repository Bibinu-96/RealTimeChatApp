package main

import (
	"backend/cmd/app/service"
	"backend/cmd/app/service/backgroundjob"
	"backend/cmd/app/service/dbinitservice"
	"backend/cmd/app/service/server"
	"backend/cmd/app/service/server/config"
	"backend/cmd/app/service/server/router"
	"backend/cmd/app/service/websocket"
	"backend/internal/channels"
	"backend/internal/database/database"
	"backend/pkg/logger"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	var log logger.Logger
	var wg sync.WaitGroup
	components := []service.Service{}

	// Create a new logger instance
	log = logger.GetLogrusLogger()
	// Set Task channel
	channels.SetTaskChannel(make(chan interface{}, 10))

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: %v", err)
	}
	// Get the values from environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")
	timezone := os.Getenv("DB_TIMEZONE")

	// Construct the DSN string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone)
	// Create a context to handle shutdown signals
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle OS signals for graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChan
		log.Info("Received shutdown signal, terminating services...")
		cancel()
	}()

	// Background Service
	bgService := backgroundjob.BackgroundJob{Log: log, Name: "BackGround Service"}

	components = append(components, bgService)

	// Database Init Service
	dbInitService := dbinitservice.DBinitService{
		Log:       log,
		Name:      "DBInitService",
		GenericDb: database.PostgressDB{DSN: dsn},
	}
	components = append(components, &dbInitService)

	// Gin Server
	serverConfig1 := config.ServerConfig{
		Addr:         ":8000",
		Router:       router.SetupGinRouter(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	ginApiServerervice := server.NewServer(serverConfig1, "GinApiServer", log)
	components = append(components, ginApiServerervice)

	// websocket server
	wsServer := websocket.Websocket{
		Addr: ":8001",
		Name: "websocket",
		Log:  log,
	}

	components = append(components, &wsServer)

	// Run all services
	for _, component := range components {
		wg.Add(1)
		go func(c service.Service) {
			defer wg.Done()

			// Panic recovery wrapper
			defer func() {
				if r := recover(); r != nil {
					log.Error("Panic occurred in service", r, c.GetName())
				}
			}()
			err := c.Run(ctx) // Pass context to allow graceful shutdown
			if err != nil {
				log.Error("Error occurred", err, c.GetName())
			}
		}(component)
	}

	// Wait for all services to complete
	wg.Wait()
	log.Info("All services stopped gracefully.")
}
