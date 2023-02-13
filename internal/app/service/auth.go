package service

import (
	"errors"
	"strings"

	"github.com/muratovdias/diplom/internal/app/repository"
	"github.com/muratovdias/diplom/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var (
	err             error
	hashPassword    []byte
	id              int
	user            models.User
	ErrInvalidPwd   = errors.New("invalid password")
	ErrInvalidEmail = errors.New("invalid email")
	ErrUserExist    = errors.New("user already exist")
	ErrUnautorized  = errors.New("unautorized")
	ErrEmailAndPwd  = errors.New("invalid email and password")
)

type AuthService interface {
	CreateClient(*models.User) (int, error)
	CreateTrainer(*models.User) (int, error)
	Authentication(string, string, string) error
}

type Auth struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *Auth {
	return &Auth{
		repo: repo,
	}
}

func (s *Auth) CreateClient(client *models.User) (int, error) {
	client.Pwd, err = generatePassword(client.Pwd)
	if err != nil {
		return 0, err
	}
	if err = validate(*client); err != nil {
		return 0, err
	}
	id, err = s.repo.CreateClient(client)
	if strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint") {
		return 0, ErrUserExist
	}
	return id, err
}
func (s *Auth) CreateTrainer(trainer *models.User) (int, error) {
	trainer.Pwd, err = generatePassword(trainer.Pwd)
	if err != nil {
		return 0, err
	}
	if err = validate(*trainer); err != nil {
		return 0, err
	}
	id, err = s.repo.CreateTrainer(trainer)
	if err != nil {
		if strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint") {
			return 0, ErrUserExist
		}
	}
	return id, err
}

func (s *Auth) Authentication(role, email, password string) error {
	user, err = s.repo.Authentication(role, email)
	if err != nil {
		return ErrUnautorized
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(password)); err != nil {
		return ErrInvalidPwd
	}
	return nil
}

func generatePassword(password string) (string, error) {
	hashPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}
