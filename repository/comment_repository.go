package repository

import (
	"hactive/final-project/dto"
	"hactive/final-project/entity"
	"hactive/final-project/interfaces"

	"gorm.io/gorm"
)

type commentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) interfaces.CommentRepository {
	return &commentRepository{
		DB: db,
	}
}

func (r *commentRepository) CreateComment(comment *dto.CreateComment) (dto.CreateComment, error) {
	cmt := entity.Comment{
		Message: comment.Message,
		PhotoID: comment.PhotoID,
		UserID:  comment.UserID,
	}

	result := r.DB.Model(&cmt).Create(&cmt)

	if result.RowsAffected < 1 {
		return *comment, result.Error
	}

	return *comment, nil
}

func (r *commentRepository) GetComment() ([]dto.GetAllComment, error) {
	var comments []dto.GetAllComment

	result := r.DB.Model(&entity.Comment{}).Find(&comments)

	if result.RowsAffected < 1 {
		return comments, result.Error
	}

	// get user and photo data
	for i, v := range comments {
		// get user data
		user := dto.GetCommentUser{}
		result := r.DB.Model(&entity.User{}).Where("id = ?", v.UserID).Find(&user)
		if result.RowsAffected < 1 {
			return comments, result.Error
		}

		v.User = user

		// get photo data
		photo := dto.GetPhotoComment{}
		result = r.DB.Model(&entity.Photo{}).Where("id = ?", v.PhotoID).Find(&photo)
		if result.RowsAffected < 1 {
			return comments, result.Error
		}

		v.Photo = photo

		comments = append(comments[:i], v)
	}

	return comments, nil
}
