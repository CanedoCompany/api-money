package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	NameTransaction string
	Value           int
	// DateTransaction time.Time
	Category string
}

// TransactioGain || TransactionLose => Transactio (value - | +)

type TransactionResponse struct {
	ID              uint      `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt,omitempty"`
	NameTransaction string    `json:"nameTransaction"`
	Value           int       `json:"value"`
	// DateTransaction time.Time `jsons:"dateTransaction"`
	Category string `json:"category"`
}
