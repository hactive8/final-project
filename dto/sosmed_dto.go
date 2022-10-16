package dto

import "time"

type CreateSosmed struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name" gorm:"not null" validate:"required"`
	SocialMediaURL string    `json:"social_media_url" gorm:"not null" validate:"required"`
	UserID         uint      `json:"user_id" gorm:"not null"`
	CreatedAt      time.Time `json:"created_at"`
}

type UpdateSosmed struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name" gorm:"not null" validate:"required"`
	SocialMediaURL string    `json:"social_media_url" gorm:"not null" validate:"required"`
	UserID         uint      `json:"user_id" gorm:"not null"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type GetAllSosmed struct {
	ID             uint      `json:"id" gorm:"column:id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           GetUser   `json:"user"`
}
