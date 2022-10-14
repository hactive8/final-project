package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string    `json:"username" gorm:"not null;unique" validate:"required,unique"`
	Email     string    `json:"email" gorm:"not null;unique" validate:"required,email,unique"`
	Password  string    `json:"password" gorm:"not null" validate:"required,min=6"`
	Age       int       `json:"age" gorm:"not null" validate:"required,min=8"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
