package service

import (
	"hactive/final-project/config"
	"hactive/final-project/dto"
	"hactive/final-project/interfaces"
	"time"
)

type CommentService struct {
	CommentRepository interfaces.CommentRepository
	Conf              config.Config
}

func NewCommentService(commentRepository interfaces.CommentRepository, conf config.Config) interfaces.CommentService {
	return &CommentService{
		CommentRepository: commentRepository,
		Conf:              conf,
	}
}

func (s *CommentService) CreateComment(comment *dto.CreateComment) (dto.GetComment, error) {
	cmt, err := s.CommentRepository.CreateComment(comment)

	data := dto.GetComment{
		Message:   cmt.Message,
		PhotoID:   cmt.PhotoID,
		UserID:    cmt.UserID,
		CreatedAt: time.Now(),
	}

	if err != nil {
		return data, err
	}

	return data, nil
}
