package repository

import (
	"errors"
	"hactive/final-project/dto"
	"hactive/final-project/entity"
	"hactive/final-project/interfaces"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Register(register *dto.Register) (dto.Register, error) {
	user := entity.User{
		Email:    register.Email,
		Password: register.Password,
		Username: register.Username,
		Age:      register.Age,
	}

	result := r.DB.Model(&user).Create(&user)

	if result.RowsAffected < 1 {
		return *register, result.Error
	}

	return *register, nil
}

func (r *repository) Login(email, password string) (dto.Login, error) {
	user := entity.User{
		Email:    email,
		Password: password,
	}

	login := dto.Login{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}

	result := r.DB.Model(&user).Where("email = ?", email).First(&login)

	if result.RowsAffected < 1 {
		return login, result.Error
	}

	return login, nil
}

func (r *repository) UpdateUser(id uint, user *dto.UpdateUser) (dto.UpdateUser, error) {
	userUpdate := entity.User{
		Email:    user.Email,
		Username: user.Username,
	}

	result := r.DB.Model(&userUpdate).Where("id = ?", id).Updates(&userUpdate)

	if result.RowsAffected < 1 {
		return *user, errors.New("user not found")
	}

	return *user, nil
}

func (r *repository) DeleteUser(id uint) error {
	user := entity.User{
		ID: id,
	}

	result := r.DB.Model(&user).Where("id = ?", id).Delete(&user)

	if result.RowsAffected < 1 {
		return errors.New("user not found")
	}

	return nil
}
