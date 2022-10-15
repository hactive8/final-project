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
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err != nil {
		return data, err
	}

	return data, nil
}
