package repository

import (
	"fmt"

	"github.com/kamalshkeir/klog"
	"github.com/kamalshkeir/korm"
	"github.com/muratovdias/diplom/internal/models"
)

var err error

type AuthRepository interface {
	CreateClient(*models.Client) (int, error)
	CreateTrainer(*models.Trainer) (int, error)
}

type AuthRepo struct {
	// db *gorm.DB
}

func NewAuthRepository() *AuthRepo {
	return &AuthRepo{}
}

func (s *AuthRepo) CreateClient(client *models.Client) (int, error) {
	err = korm.AutoMigrate[models.Client]("client")
	if klog.CheckError(err) {
		fmt.Println("error create tables")
		return 0, err
	}
	id, err := korm.Model[models.Client]().Insert(client)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	return id, nil
}
func (s *AuthRepo) CreateTrainer(trainer *models.Trainer) (int, error) {
	err = korm.AutoMigrate[models.Trainer]("trainer")
	if klog.CheckError(err) {
		fmt.Println("error create tables")
		return 0, err
	}
	id, err := korm.Model[models.Trainer]().Insert(trainer)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	return id, nil
}

func (s *AuthRepo) Authentication(role, email, password string) error {

	return nil
}
