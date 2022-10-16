package repository

import (
	"hactive/final-project/dto"
	"hactive/final-project/entity"
	"hactive/final-project/interfaces"

	"gorm.io/gorm"
)

type sosmedRepository struct {
	DB *gorm.DB
}

func NewSosmedRepository(db *gorm.DB) interfaces.SosmedRepository {
	return &sosmedRepository{
		DB: db,
	}
}

func (r *sosmedRepository) CreateSosmed(sosmed *dto.CreateSosmed) (dto.CreateSosmed, error) {
	sos := entity.Sosmed{
		Name:           sosmed.Name,
		SocialMediaURL: sosmed.SocialMediaURL,
		UserID:         sosmed.UserID,
	}

	result := r.DB.Model(&sos).Create(&sos)

	if result.RowsAffected < 1 {
		return *sosmed, result.Error
	}

	return *sosmed, nil
}
