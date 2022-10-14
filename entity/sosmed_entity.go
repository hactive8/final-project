package entity

import "gorm.io/gorm"

type Sosmed struct {
	gorm.Model
	ID             uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name           string `json:"name" gorm:"not null" validate:"required"`
	SocialMediaURL string `json:"social_media_url" gorm:"not null" validate:"required"`
	UserID         uint   `json:"user_id" gorm:"not null"`
	User           User   `json:"-" gorm:"foreignKey:UserID"`
}
