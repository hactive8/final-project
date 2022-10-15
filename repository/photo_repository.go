package repository

import (
	"hactive/final-project/dto"
	"hactive/final-project/entity"
	"hactive/final-project/interfaces"

	"gorm.io/gorm"
)

type photorepository struct {
	DB *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) interfaces.PhotoRepository {
	return &photorepository{
		DB: db,
	}
}

func (r *photorepository) CreatePhoto(photo *dto.CreatePhoto) (dto.CreatePhoto, error) {
	photos := entity.Photo{
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoURL: photo.PhotoURL,
		UserID:   photo.UserId,
	}

	result := r.DB.Model(&photos).Create(&photos)

	if result.RowsAffected < 1 {
		return *photo, result.Error
	}

	return *photo, nil
}
