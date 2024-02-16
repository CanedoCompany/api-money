package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
 gorm.Model
 ID string
 Name string
 Email string
 Password string
 Phone int64
}

func NewUser(name string, email string, password string, phone int64) *User {
 return &User{
	ID: uuid.New().String(),
	Name: name,
	Email: email,
	Password: password,
	Phone: phone,
 }
}