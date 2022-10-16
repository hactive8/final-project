package dto

import "time"

type CreateSosmed struct {
	Name           string    `json:"name" gorm:"not null" validate:"required"`
	SocialMediaURL string    `json:"social_media_url" gorm:"not null" validate:"required"`
	UserID         uint      `json:"user_id" gorm:"not null"`
	CreatedAt      time.Time `json:"created_at"`
}
