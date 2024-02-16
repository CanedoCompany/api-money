package config

import (
	"gorm.io/gorm"
)

var (
	bd *gorm.DB
	logger *Logger
)

func Init() error {
  return nil
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}