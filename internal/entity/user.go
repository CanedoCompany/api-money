package entity

import (
	"time"
	// "github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Phone    int64
}

type UserResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Phone     int64     `json:"phone"`
}

/*
func NewUser(name string, email string, password string, phone int64) *User {
 return &User{
	ID: uuid.New().String(),
	Name: name,
	Email: email,
	Password: password,
	Phone: phone,
 }
}
*/
