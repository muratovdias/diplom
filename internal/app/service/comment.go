package service

import (
	"github.com/muratovdias/diplom/internal/app/repository"
	"github.com/muratovdias/diplom/internal/models"
)

type Commet interface {
	Create(comment models.Comment) error
	GetCommentsByTrainerID(id int) ([]models.Comment, error)
}

type CommetnService struct {
	repo repository.Comment
}

func NewCommentService(repo repository.Comment) *CommetnService {
	return &CommetnService{
		repo: repo,
	}
}

func (c *CommetnService) Create(comment models.Comment) error {
	return c.repo.Create(comment)
}

func (c *CommetnService) GetCommentsByTrainerID(id int) ([]models.Comment, error) {
	return c.repo.GetCommentsByTrainerID(id)
}
