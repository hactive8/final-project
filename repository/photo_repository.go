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

func (r *photorepository) GetAllPhoto() ([]dto.GetAllPhoto, error) {
	res := dto.GetUser{}

	photo := entity.Photo{}
	data := []dto.GetAllPhoto{}

	result := r.DB.Model(&photo).Find(&data)

	if result.RowsAffected < 1 {
		return data, result.Error
	}

	// user
	for i, v := range data {
		result := r.DB.Model(&entity.User{}).Where("id = ?", v.UserID).First(&res)
		if result.RowsAffected < 1 {
			return data, result.Error
		}
		v.User = res

		data = append(data[:i], v)
	}

	return data, nil
}

func (r *photorepository) UpdatePhoto(id int, photo *dto.UpdatePhoto) (dto.UpdatePhoto, error) {
	photos := entity.Photo{
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoURL: photo.PhotoURL,
		UserID:   photo.UserId,
	}

	result := r.DB.Model(&photos).Where("id = ?", id).Updates(&photos)

	if result.RowsAffected < 1 {
		return *photo, result.Error
	}

	return *photo, nil
}

func (r *photorepository) DeletePhoto(id int) error {
	photo := entity.Photo{}

	result := r.DB.Model(&photo).Where("id = ?", id).Delete(&photo)

	if result.RowsAffected < 1 {
		return result.Error
	}

	return nil
}
