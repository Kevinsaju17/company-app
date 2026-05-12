package models

import "gorm.io/gorm"


// User model represents users table
type User struct {

	// gorm.Model automatically adds:
	// ID
	// CreatedAt
	// UpdatedAt
	// DeletedAt
	gorm.Model

	Name string `json:"name" validate:"required"`

	// unique prevents duplicate emails
	Email string `json:"email" gorm:"unique" validate:"required,email"`

	Password string `json:"password" validate:"required,min=6"`
}