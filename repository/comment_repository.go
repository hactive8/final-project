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
