package service

import (
	"hactive/final-project/config"
	"hactive/final-project/dto"
	"hactive/final-project/interfaces"
	"time"
)

type PhotoService struct {
	PhotoRepository interfaces.PhotoRepository
	conf            config.Config
}

func NewPhotoService(repo interfaces.PhotoRepository, conf config.Config) interfaces.PhotoService {
	return &PhotoService{
		PhotoRepository: repo,
		conf:            conf,
	}
}

func (s *PhotoService) CreatePhoto(photo *dto.CreatePhoto) (dto.GetPhoto, error) {
	photos, err := s.PhotoRepository.CreatePhoto(photo)
	data := dto.GetPhoto{
		Title:     photos.Title,
		Caption:   photos.Caption,
		PhotoURL:  photos.PhotoURL,
		UserID:    photos.UserId,
		CreatedAt: time.Now(),
	}
	if err != nil {
		return data, err
	}

	return data, nil
}

func (s *PhotoService) GetAllPhoto() ([]dto.GetAllPhoto, error) {
	return s.PhotoRepository.GetAllPhoto()
}

func (s *PhotoService) UpdatePhoto(id int, photo *dto.UpdatePhoto) (dto.GetUpdatePhoto, error) {
	photos, err := s.PhotoRepository.UpdatePhoto(id, photo)
	data := dto.GetUpdatePhoto{
		Title:     photos.Title,
		Caption:   photos.Caption,
		PhotoURL:  photos.PhotoURL,
		UserID:    photos.UserId,
		UpdatedAt: time.Now(),
	}

	if err != nil {
		return data, err
	}

	return data, nil
}
