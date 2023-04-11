package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/muratovdias/diplom/internal/models"
)

type Trainer interface {
	SetSchedule(id int, schedule map[string][]string) error
	GetFullTrainerInfo(id int) (models.TrainerInfo, error)
	ViewAllTrainings(id int) ([]models.Training, error)
	UpdateTrainerInfo(trainer models.TrainerInfo) error
	CancelSchedule(id int, times []string) error
}

type TrainerRepo struct {
	db *sqlx.DB
}

func NewTrainerRepo(db *sqlx.DB) *TrainerRepo {
	return &TrainerRepo{
		db: db,
	}
}

func (t *TrainerRepo) SetSchedule(id int, schedule map[string][]string) error {
	_, err := t.db.Exec("SET TIME ZONE 'Asia/Almaty'")
	if err != nil {
		panic(err)
	}
	stmt, err := t.db.Prepare(`INSERT INTO trainer_schedule(user_id, date) 
								VALUES($1, $2)`)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// query := `INSERT INTO trainer_schedule(trainer_id, date) VALUES($1, $2)`
	for dateStr, times := range schedule {
		_, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		for _, timeStr := range times {
			datetimeStr := fmt.Sprintf("%s %s", dateStr, timeStr)
			datetime, err := time.Parse("2006-01-02 15:04", datetimeStr)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			_, err = stmt.Exec(id, datetime)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
		}
	}
	return nil
}

func (t *TrainerRepo) GetFullTrainerInfo(id int) (models.TrainerInfo, error) {
	var info models.TrainerInfo
	query := `SELECT user_id, full_name, email, speciality, phone, adress, image, bio, twitter, instagram, facebook
			FROM users 
			INNER JOIN trainer
			USING (user_id)
			WHERE user_id=$1`
	row := t.db.QueryRow(query, id)
	if err := row.Scan(&info.ID, &info.FullName, &info.Email, &info.Speciality, &info.Phone, &info.Adress, &info.Img, &info.Bio, &info.Twitter, &info.Instagram, &info.Facebook); err != nil {
		return info, fmt.Errorf("repo: Gett Full Trainer Info: %w", err)
	}
	return info, nil
}

func (t *TrainerRepo) ViewAllTrainings(id int) ([]models.Training, error) {
	var trainings []models.Training
	query := `SELECT full_name, note, date, role
			FROM client_schedule c 
			INNER JOIN users u
			ON c.user_id = u.user_id
			WHERE c.trainer_id=$1`
	rows, err := t.db.Query(query, id)
	if err != nil {
		fmt.Println("repo: View All Trainings: ", err.Error())
		return trainings, err
	}
	i := 1
	for rows.Next() {
		var training models.Training
		var date time.Time
		err := rows.Scan(&training.Name, &training.Note, &date, &training.Role)
		if err != nil {
			fmt.Println(err.Error())
			return trainings, err
		}
		training.Date = date.Format("2006-01-02 15:04:05")
		training.Row = i
		i++
		trainings = append(trainings, training)
	}
	// fmt.Println(trainings)
	return trainings, nil
}

func (t *TrainerRepo) UpdateTrainerInfo(trainer models.TrainerInfo) error {
	query := `UPDATE trainer
			SET speciality = $1,
				phone = $2,
				adress = $3,
				image = $4,
				bio = $5,
				twitter = $6,
				instagram = $7,
				facebook = $8
			WHERE user_id=$9`
	_, err := t.db.Exec(query, trainer.Speciality, trainer.Phone,
		trainer.Adress, trainer.Img, trainer.Bio, trainer.Twitter,
		trainer.Instagram, trainer.Facebook, trainer.ID)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if trainer.FullName != "" && strings.TrimSpace(trainer.FullName) != "" {
		query := `UPDATE users
				SET full_name = $1
				WHERE user_id = $2`
		_, err := t.db.Exec(query, trainer.FullName, trainer.ID)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	fmt.Println("updated")
	return nil
}

func (t *TrainerRepo) CancelSchedule(id int, times []string) error {
	query := `DELETE FROM trainer_schedule
			WHERE user_id=$1 AND date=$2`
	for _, time := range times {
		_, err := t.db.Exec(query, id, time)
		if err != nil {
			fmt.Printf("trainer repo: CancleSchedule: %s", err.Error())
			return err
		}
	}
	return nil
}
