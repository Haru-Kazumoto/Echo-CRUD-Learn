package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
	SSLMode  string
	TimeZone string
}

func (config DatabaseConfig) GetDsn() string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.Host,
		config.User,
		config.Password,
		config.DBName,
		config.Port,
		config.SSLMode,
		config.TimeZone)

	return dsn
}

func ConnectDB() (*gorm.DB, error){
	databaseConfig := DatabaseConfig{
		Host: "localhost",
		User: "postgres",
		Password: "admin",
		DBName: "echo-db",
		Port: 5432,
		SSLMode: "disable",
		TimeZone: "UTC",
	}

	dsn := databaseConfig.GetDsn()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		return nil, err
	}

	return db, nil
}