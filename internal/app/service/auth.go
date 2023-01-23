package service

import (
	"fmt"

	"github.com/kamalshkeir/argon"
	"github.com/muratovdias/diplom/internal/app/repository"
	"github.com/muratovdias/diplom/internal/models"
)

var err error

type AuthService interface {
	CreateClient(*models.Client) (int, error)
	CreateTrainer(*models.Trainer) (int, error)
}

type Auth struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *Auth {
	return &Auth{
		repo: repo,
	}
}

func (s *Auth) CreateClient(client *models.Client) (int, error) {
	client.Password, err = argon.Hash(client.Password)
	// fmt.Println(client.Password)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	return s.repo.CreateClient(client)
}
func (s *Auth) CreateTrainer(trainer *models.Trainer) (int, error) {
	return s.repo.CreateTrainer(trainer)
}
