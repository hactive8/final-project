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

func (s *CommentService) GetComment() ([]dto.GetAllComment, error) {
	comments, err := s.CommentRepository.GetComment()

	if err != nil {
		return comments, err
	}

	return comments, nil
}

func (s *CommentService) UpdateComment(id uint, comment *dto.UpdateComment) (*dto.GetPhotoUser, error) {
	cmt, err := s.CommentRepository.UpdateComment(id, comment)
	if err != nil {
		return nil, err
	}

	photo, er := s.CommentRepository.GetPhotoId(id)
	if er != nil {
		return &photo, er
	}

	cmt.UserID = photo.UserID

	return &photo, nil
}

func (s *CommentService) GetPhotoId(id uint) (dto.GetPhotoUser, error) {
	photo, err := s.CommentRepository.GetPhotoId(id)

	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (s *CommentService) DeleteComment(id uint) error {
	err := s.CommentRepository.DeleteComment(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *CommentService) GetCommentId(id uint) (dto.GetComment, error) {
	cmt, err := s.CommentRepository.GetCommentId(id)

	if err != nil {
		return cmt, err
	}

	return cmt, nil
}
