package entity

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Message   string    `json:"message" gorm:"not null" validate:"required"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	User      User      `json:"-" gorm:"foreignKey:UserID"`
	PhotoID   uint      `json:"photo_id" gorm:"not null"`
	Photo     Photo     `json:"-" gorm:"foreignKey:PhotoID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
