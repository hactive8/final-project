package dto

type CreatePhoto struct {
	ID       uint   `json:"id"`
	Title    string `json:"title" gorm:"not null" validate:"required"`
	Caption  string `json:"caption" gorm:"not null"`
	PhotoURL string `json:"photo_url" gorm:"not null" validate:"required"`
	UserId   uint   `json:"user_id" gorm:"not null" validate:"required"`
}

type GetPhoto struct {
	ID        uint   `json:"id" gorm:"column:id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoURL  string `json:"photo_url"`
	UserID    uint   `json:"user_id"`
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
}
