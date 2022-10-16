package dto

import "time"

type CreateComment struct {
	Message string `json:"message" gorm:"not null" validate:"required"`
	PhotoID uint   `json:"photo_id" gorm:"not null"`
	UserID  uint   `json:"user_id" gorm:"not null"`
}

type GetComment struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetAllComment struct {
	ID        uint            `json:"id"`
	Message   string          `json:"message"`
	PhotoID   uint            `json:"photo_id"`
	UserID    uint            `json:"user_id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	User      GetCommentUser  `json:"user"`
	Photo     GetPhotoComment `json:"photo"`
}

type UpdateComment struct {
	Message string `json:"message" gorm:"not null" validate:"required"`
	UserID  uint   `json:"user_id" gorm:"not null"`
}

type GetPhotoUser struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
