package service

import (
	"hactive/final-project/config"
	"hactive/final-project/dto"
	"hactive/final-project/interfaces"
	"hactive/final-project/utils"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Conf config.Config
	repo interfaces.UserRepository
}

func NewUserService(conf config.Config, repo interfaces.UserRepository) interfaces.UserService {
	return &Service{
		Conf: conf,
		repo: repo,
	}
}

func (s *Service) Register(user *dto.Register) (dto.Register, error) {
	// bcrypt password
	pass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)

	// set password
	user.Password = string(pass)

	return s.repo.Register(user)
}

func (s *Service) Login(email, password string) (string, error) {
	login, err := s.repo.Login(email, password)
	if err != nil {
		return "", err
	}
	_ = bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(password))

	token, _ := utils.GenerateAccessToken(login.ID, login.Email)

	return token, nil
}

func (s *Service) UpdateUser(id uint, user *dto.UpdateUser) (dto.UpdateUser, error) {
	return s.repo.UpdateUser(id, user)
}

func (s *Service) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
