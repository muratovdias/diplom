package service

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/muratovdias/diplom/internal/app/repository"
	"github.com/muratovdias/diplom/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidPwd   = errors.New("invalid password")
	ErrInvalidEmail = errors.New("invalid email")
	ErrUserExist    = errors.New("user already exists")
	ErrUnautorized  = errors.New("unautorized")
	ErrEmailAndPwd  = errors.New("invalid email and password")
)

type AuthService interface {
	CreateUser(*models.User) error
	Authentication(email, password string) (models.User, error)
	GetUserByToken(token string) (models.User, error)
}

type Auth struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *Auth {
	return &Auth{
		repo: repo,
	}
}

func (s *Auth) CreateUser(user *models.User) error {
	var err error
	user.Password, err = generatePassword(user.Password)
	if err != nil {
		return err
	}
	if err = validate(*user); err != nil {
		return err
	}
	var img string
	if user.Role == "trainer" {
		img, err = ConvertToBinary("./ui/img/ava.png")
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	err = s.repo.CreateUser(user, img)
	if err != nil {
		if strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint") {
			return ErrUserExist
		} else {
			return err
		}
	}
	return err
}

func (s *Auth) Authentication(email, password string) (models.User, error) {
	user, err := s.repo.Authentication(email)
	if err != nil {
		log.Println(err.Error())
		return models.User{}, ErrUnautorized
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Println(err.Error())
		return models.User{}, ErrInvalidPwd
	}
	token, err := uuid.NewV4()
	if err != nil {
		log.Printf("service: %s\n", err)
		return models.User{}, err
	}
	user.Token = token.String()
	user.TokenDuration = time.Now().Add(24 * time.Hour)
	if err := s.repo.UpdateToken(email, user.Token, user.TokenDuration); err != nil {
		fmt.Println("service: " + err.Error())
		return models.User{}, err
	}
	// user.TokenDuration = time.Now().Add(12 * time.Hour)
	return user, nil
}

func (s *Auth) GetUserByToken(token string) (models.User, error) {
	return s.repo.GetUserByToken(token)
}

func generatePassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}
