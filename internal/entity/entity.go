package entity

import "github.com/google/uuid"

type User struct {
 ID string
 Name string
 Email string
 Password string
 Phone number
}

func NewUser(name string, email string, password string, phone number) *User {
 return &User{
	ID: uuid.New().String(),
	Name: name,
	Email: email,
	Password: password,
	Phone: phone,
 }
}