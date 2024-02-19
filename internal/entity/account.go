package entity

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	OwnerName string
}
