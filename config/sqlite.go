package config

import (
	"os"

	"github.com/CanedoCompany/api-money/internal/entity"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

func InitializeSQLite() (*gorm.DB, error){
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"


	// Check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite emoney error: %v", err)
		return nil, err
	}

	// Migrate the entity
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		logger.Errorg("sqlite automigration error: %v", err)
		return nil, err
	}

	return db, nil
}