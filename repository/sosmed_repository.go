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

func (r *sosmedRepository) GetSosmed() ([]dto.GetAllSosmed, error) {
	res := dto.GetUser{}

	sos := entity.Sosmed{}
	data := []dto.GetAllSosmed{}

	result := r.DB.Model(&sos).Find(&data)

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

func (r *sosmedRepository) UpdateSosmed(id uint, sosmed *dto.UpdateSosmed) (dto.UpdateSosmed, error) {
	sos := entity.Sosmed{
		Name:           sosmed.Name,
		SocialMediaURL: sosmed.SocialMediaURL,
		UserID:         sosmed.UserID,
	}

	result := r.DB.Model(&sos).Where("id = ?", id).Updates(&sos)

	if result.RowsAffected < 1 {
		return *sosmed, result.Error
	}

	return *sosmed, nil
}

func (r *sosmedRepository) DeleteSosmed(id uint) error {
	sos := entity.Sosmed{}

	result := r.DB.Model(&sos).Where("id = ?", id).Delete(&sos)

	if result.RowsAffected < 1 {
		return result.Error
	}

	return nil
}

func (r *sosmedRepository) GetSosmedByID(id uint) (dto.GetAllSosmed, error) {
	sos := entity.Sosmed{}
	data := dto.GetAllSosmed{}

	result := r.DB.Model(&sos).Where("user_id = ?", id).First(&data)

	if result.RowsAffected < 1 {
		return data, result.Error
	}

	return data, nil
}
