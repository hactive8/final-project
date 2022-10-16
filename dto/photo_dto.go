package dto

import "time"

type CreatePhoto struct {
	ID       uint   `json:"id"`
	Title    string `json:"title" gorm:"not null" validate:"required"`
	Caption  string `json:"caption" gorm:"not null"`
	PhotoURL string `json:"photo_url" gorm:"not null" validate:"required"`
	UserId   uint   `json:"user_id" gorm:"not null" validate:"required"`
}

type GetPhoto struct {
	ID        uint      `json:"id" gorm:"column:id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

type GetAllPhoto struct {
	ID        uint      `json:"id" gorm:"column:id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
	User      GetUser   `json:"user" gorm:"foreignKey:UserID"`
}

type UpdatePhoto struct {
	Title    string `json:"title" gorm:"not null" validate:"required"`
	Caption  string `json:"caption" gorm:"not null"`
	PhotoURL string `json:"photo_url" gorm:"not null" validate:"required"`
	UserId   uint   `json:"user_id" gorm:"not null"`
}

type GetUpdatePhoto struct {
	ID        uint      `json:"id" gorm:"column:id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type GetPhotoComment struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}
