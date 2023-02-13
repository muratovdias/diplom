package repository

import (
	"errors"
	"fmt"

	"github.com/kamalshkeir/klog"
	"github.com/kamalshkeir/korm"
	"github.com/muratovdias/diplom/internal/models"
)

var (
	err     error
	trainer models.User
	client  models.User
)

type AuthRepository interface {
	CreateClient(*models.User) (int, error)
	CreateTrainer(*models.User) (int, error)
	Authentication(string, string) (models.User, error)
}

type AuthRepo struct {
	// db *gorm.DB
}

func NewAuthRepository() *AuthRepo {
	return &AuthRepo{}
}

func (s *AuthRepo) CreateClient(client *models.User) (int, error) {
	err = korm.AutoMigrate[models.User]("client")
	if klog.CheckError(err) {
		fmt.Println("error create clinet table", err.Error())
		return 0, err
	}
	id, err := korm.Model[models.User]("client").Insert(client)
	if err != nil {
		// fmt.Println(err.Error())
		return 0, err
	}
	return id, nil
}
func (s *AuthRepo) CreateTrainer(trainer *models.User) (int, error) {
	err = korm.AutoMigrate[models.User]("trainer")
	if klog.CheckError(err) {
		fmt.Println("error create trainer table", err.Error())
		return 0, err
	}
	id, err := korm.Model[models.User]("trainer").Insert(trainer)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *AuthRepo) Authentication(role, email string) (models.User, error) {
	switch role {
	case "trainer":
		trainer, err = korm.Model[models.User]("trainer").Where("email=?", email).Select("pwd").One()
		if err != nil {
			fmt.Println("trainer", err.Error())
			return models.User{}, err
		}
		return trainer, nil
	case "client":
		client, err = korm.Model[models.User]("client").Select("pwd").Where("email = ?", email).One()
		if err != nil {
			fmt.Println("client", err.Error())
			return client, err
		}
		return client, nil
	default:
		return models.User{}, errors.New("no such role")
	}
}
