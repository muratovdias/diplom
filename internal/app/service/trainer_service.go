package service

import (
	"fmt"

	"github.com/muratovdias/diplom/internal/app/repository"
	"github.com/muratovdias/diplom/internal/models"
)

type Trainer interface {
	SetSchedule(id int, schedule map[string][]string) error
	// GetAllTrainers() ([]models.Trainers, error)
	GetFullTrainerInfo(id int) (models.TrainerInfo, error)
	ViewAllTrainings(id int) ([]models.Training, error)
	UpdateTrainerInfo(trainer models.TrainerInfo) error
}

type TrainerService struct {
	repo repository.Trainer
}

func NewTrainerService(repo repository.Trainer) *TrainerService {
	return &TrainerService{
		repo: repo,
	}
}

func (t *TrainerService) SetSchedule(id int, schedule map[string][]string) error {

	return t.repo.SetSchedule(id, schedule)
}

// func (t *TrainerService) GetAllTrainers() ([]models.Trainers, error) {
// 	trainers, err := t.repo.GetAllTrainers()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil, err
// 	}
// 	// for _, trainer := range trainers {
// 	// 	trainer.Img = template.URL(base64.StdEncoding.EncodeToString([]byte(trainer.Img)))
// 	// }
// 	return trainers, nil
// }

func (t *TrainerService) GetFullTrainerInfo(id int) (models.TrainerInfo, error) {
	trainer, err := t.repo.GetFullTrainerInfo(id)
	if err != nil {
		fmt.Println("service: " + err.Error())
		return trainer, err
	}
	return trainer, nil
}

func (t *TrainerService) ViewAllTrainings(id int) ([]models.Training, error) {
	return t.repo.ViewAllTrainings(id)
}

func (t *TrainerService) UpdateTrainerInfo(trainer models.TrainerInfo) error {
	if trainer.Bio == "" {
		trainer.Bio = "NO"
	}
	if trainer.Phone == "" {
		trainer.Phone = "NO"
	}
	if trainer.Speciality == "" {
		trainer.Speciality = "trainer"
	}
	if trainer.Twitter == "" {
		trainer.Twitter = "NO"
	}
	if trainer.Facebook == "" {
		trainer.Facebook = "NO"
	}
	if trainer.Instagram == "" {
		trainer.Instagram = "NO"
	}
	return t.repo.UpdateTrainerInfo(trainer)
}
