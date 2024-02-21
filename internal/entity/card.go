package entity

import (
	"time"

	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	NameCard   string
	NameOwner  string
	NumberCard int
	// Validate time.Time
	CvvCard int
}

type CardResponse struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
	NameCard   string    `json:"nameCard"`
	NameOwner  string    `json:"nameOwner"`
	NumberCard int       `json:"numberCard"`
	CvvCard    int       `json:"cvvCard"`
}
