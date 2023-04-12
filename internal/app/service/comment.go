package service

import (
	"github.com/muratovdias/diplom/internal/app/repository"
	"github.com/muratovdias/diplom/internal/models"
)

type Commet interface {
	Create(comment models.Comment) error
	Delete(authorID, trainerID int) error
	GetCommentsByTrainerID(userID, trainerID int) ([]models.Comment, error)
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

func (c *CommetnService) Delete(authorID, trainerID int) error {
	return c.repo.Delete(authorID, trainerID)
}

func (c *CommetnService) GetCommentsByTrainerID(userID, trainerID int) ([]models.Comment, error) {
	comments, err := c.repo.GetCommentsByTrainerID(trainerID)
	if err != nil {
		return comments, err
	}
	for i := 0; i < len(comments); i++ {
		// fmt.Println(comment.AuthorID == userID)
		if comments[i].AuthorID == userID {
			comments[i].Edit = true
		}
		// fmt.Println(comment.Edit)
	}
	return comments, nil
}
