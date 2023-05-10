package service

import (
	"fmt"

	"github.com/muratovdias/diplom/internal/app/repository"
	"github.com/muratovdias/diplom/internal/models"
)

type Client interface {
	ViewSchedule(id int) (map[string][]models.Time, error)
	GetAllTrainers() ([]models.Trainers, error)
	SetTrainings(id, trainerID int, note string, dates []string) error
	ViewTrainings(id int) ([]models.Training, error)
	CancelTraining(id int, date string) error
	MyStats(id int) (models.Stats, error)
}

type ClientService struct {
	repo repository.Client
}

func NewClientService(repo repository.Client) *ClientService {
	return &ClientService{
		repo: repo,
	}
}

func (c *ClientService) ViewSchedule(id int) (map[string][]models.Time, error) {
	schedule, err := c.repo.ViewSchedule(id)
	// var res models.Times
	if err != nil {
		return schedule, err
	}
	return c.repo.ViewSchedule(id)
}

func (c *ClientService) SetTrainings(id, trainerID int, note string, dates []string) error {
	return c.repo.SetTrainings(id, trainerID, note, dates)
}

func (c *ClientService) GetAllTrainers() ([]models.Trainers, error) {
	trainers, err := c.repo.GetAllTrainers()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	// for _, trainer := range trainers {
	// 	trainer.Img = template.URL(base64.StdEncoding.EncodeToString([]byte(trainer.Img)))
	// }
	return trainers, nil
}

func (c *ClientService) ViewTrainings(id int) ([]models.Training, error) {
	return c.repo.ViewTrainings(id)
}

func (c *ClientService) CancelTraining(id int, date string) error {
	return c.repo.CancelTraining(id, date)
}

func (c *ClientService) MyStats(id int) (models.Stats, error) {
	return c.repo.MyStats(id)
}
