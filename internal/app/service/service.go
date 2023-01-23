package service

import "github.com/muratovdias/diplom/internal/app/repository"

type Service struct {
	AuthService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AuthService: NewAuthService(repo.AuthRepository),
	}
}
