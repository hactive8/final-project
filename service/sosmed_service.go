package service

import (
	"hactive/final-project/config"
	"hactive/final-project/dto"
	"hactive/final-project/interfaces"
)

type SosmedService struct {
	SosmedRepository interfaces.SosmedRepository
	conf             config.Config
}

func NewSosmedService(sosmedRepository interfaces.SosmedRepository, conf config.Config) interfaces.SosmedService {
	return &SosmedService{
		SosmedRepository: sosmedRepository,
		conf:             conf,
	}
}

func (s *SosmedService) CreateSosmed(sosmed *dto.CreateSosmed) (dto.CreateSosmed, error) {
	sos, err := s.SosmedRepository.CreateSosmed(sosmed)

	if err != nil {
		return sos, err
	}

	return sos, nil
}
