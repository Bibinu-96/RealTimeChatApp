package database

import (
	"fmt"

	"gorm.io/gorm"
)

type Database interface {
	InitDB() (*gorm.DB, error)
	RunMigrations(*gorm.DB) error
}

type DatabaseType string

const (
	PostgreSQL DatabaseType = "postgres"
	MySQL      DatabaseType = "mysql"
	SQLite     DatabaseType = "sqlite"
	SQLServer  DatabaseType = "sqlserver"
)

type DatabaseConfig struct {
	Type     DatabaseType // Database type (e.g., postgres, mysql, sqlite)
	Host     string       // Hostname or IP
	Port     int          // Port number
	User     string       // Username
	Password string       // Password
	DBName   string       // Database name
	Options  string       // Additional options (e.g., sslmode, charset)
}

func GenerateDSN(cfg DatabaseConfig) (string, error) {
	switch cfg.Type {
	case PostgreSQL:
		if cfg.Host == "" {
			cfg.Host = "localhost"
		}
		if cfg.Port == 0 {
			cfg.Port = 5432
		}
		return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s %s",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.Options), nil

	case MySQL:
		if cfg.Host == "" {
			cfg.Host = "127.0.0.1"
		}
		if cfg.Port == 0 {
			cfg.Port = 3306
		}
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.Options), nil

	case SQLite:
		if cfg.DBName == "" {
			return "", fmt.Errorf("SQLite requires a database name (file path)")
		}
		return cfg.DBName, nil

	case SQLServer:
		if cfg.Host == "" {
			cfg.Host = "localhost"
		}
		if cfg.Port == 0 {
			cfg.Port = 1433
		}
		return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&%s",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.Options), nil

	default:
		return "", fmt.Errorf("unsupported database type: %s", cfg.Type)
	}
}
