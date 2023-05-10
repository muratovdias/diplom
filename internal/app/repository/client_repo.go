package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
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

type ClientRepo struct {
	db *sqlx.DB
}

func NewClientRepo(db *sqlx.DB) *ClientRepo {
	return &ClientRepo{
		db: db,
	}
}

func (c *ClientRepo) ViewSchedule(id int) (map[string][]models.Time, error) {
	_, err := c.db.Exec("SET TIME ZONE 'Asia/Almaty'")
	if err != nil {
		panic(err)
	}
	schedule := make(map[string][]models.Time)
	query := `SELECT date, available 
			FROM trainer_schedule 
			WHERE user_id=$1 AND date::date >= CURRENT_DATE::date
			ORDER BY date`

	rows, err := c.db.Query(query, id)
	if err != nil {
		fmt.Println("view schedule ", err.Error())
		return schedule, err
	}
	for rows.Next() {
		var temp models.Time
		var date time.Time
		var available bool
		if err := rows.Scan(&date, &available); err != nil {
			fmt.Println("view schedule: scan", err.Error())
		}
		// fmt.Println(date, time.Now(), date.Before(time.Now()))
		if date.Before(time.Now()) || !available {
			temp.Flag = true
		}
		dateStr := date.Format("2006-01-02 15:04:05")[0:10]
		temp.Time = date.Format("2006-01-02 15:04:05")[11:16]
		schedule[dateStr] = append(schedule[dateStr], temp)
	}
	// fmt.Println(schedule)
	return schedule, nil
}

func (c *ClientRepo) GetAllTrainers() ([]models.Trainers, error) {
	query := `SELECT user_id, full_name, speciality, image
			FROM users INNER JOIN trainer
			USING (user_id)`

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("repo: GetAllTrainers: %w", err)
	}
	var res []models.Trainers
	for rows.Next() {
		var temp models.Trainers
		if err := rows.Scan(&temp.ID, &temp.FullName, &temp.Speciality, &temp.Img); err != nil {
			return nil, fmt.Errorf("repo: GetAllTrainers: %w", err)
		}
		res = append(res, temp)
	}
	// fmt.Println(res)
	return res, nil
}

func (c *ClientRepo) SetTrainings(id, trainerID int, note string, dates []string) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Printf("Error rolling back transaction: %v", err)
		}
	}()
	query := `INSERT INTO client_schedule(user_id, trainer_id, note, date) 
			VALUES($1, $2, $3, $4)`

	for _, time := range dates {
		_, err := tx.Exec("SET TIME ZONE 'Asia/Almaty'")
		if err != nil {
			fmt.Println(err)
			return err
		}
		_, err = tx.Exec(query, id, trainerID, note, time)
		if err != nil {
			fmt.Println(err)
			return err
		}
		_, err = tx.Exec(`UPDATE trainer_schedule 
		SET available=false WHERE user_id = $1 AND date=$2`, trainerID, time)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	tx.Commit()
	return nil
}

func (c *ClientRepo) ViewTrainings(id int) ([]models.Training, error) {
	_, err := c.db.Exec("SET TIME ZONE 'Asia/Almaty'")
	if err != nil {
		panic(err)
	}
	var trainings []models.Training
	query := `SELECT full_name, note, date, role
			FROM client_schedule c INNER JOIN users u
			ON c.trainer_id = u.user_id
			WHERE c.user_id=$1`

	rows, err := c.db.Query(query, id)
	if err != nil {
		fmt.Println("repo: client: ViewTrainings: ", err.Error())
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
	return trainings, nil
}

func (c *ClientRepo) CancelTraining(id int, date string) error {
	tx, err := c.db.Begin()
	if err != nil {
		log.Printf("repo: client: CancelTraining: %v", err.Error())
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Printf("Error rolling back transaction: %v", err)
		}
	}()
	var trainer_id int
	query1 := `SELECT trainer_id 
				FROM client_schedule
				WHERE user_id=$1 AND date=$2`

	row := tx.QueryRow(query1, id, date)
	if err := row.Scan(&trainer_id); err != nil {
		log.Printf("repo: client: CancelTraining: %v", err.Error())
		return err
	}
	query2 := `UPDATE client_schedule
				SET canceled = true
				WHERE user_id=$1 AND date=$2`

	_, err = tx.Exec(query2, id, date)
	if err != nil {
		log.Printf("repo: client: CancelTraining: %v", err.Error())
		return err
	}
	query3 := `UPDATE trainer_schedule 
				SET available=true
				WHERE user_id=$1 AND date=$2`

	_, err = tx.Exec(query3, trainer_id, date)
	if err != nil {
		log.Printf("repo: client: CancelTraining: %v", err.Error())
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("repo: client: CancelTraining: %v", err.Error())
		return err
	}
	return nil
}

func (c *ClientRepo) MyStats(id int) (models.Stats, error) {
	var stats models.Stats
	query1 := `SELECT COUNT(*) FROM client_schedule WHERE user_id = $1`
	query2 := `SELECT COUNT(canceled) FROM client_schedule WHERE user_id = $1 AND cancled = true`
	query3 := `SELECT COUNT(completed) FROM client_schedule WHERE user_id = $1 AND completed = true`

	row := c.db.QueryRow(query1, id)
	if err := row.Scan(&stats.All); err != nil {
		return stats, fmt.Errorf("repo: MyStats: %w", err)
	}

	row = c.db.QueryRow(query2, id)
	if err := row.Scan(&stats.Canceled); err != nil {
		return stats, fmt.Errorf("repo: MyStats: %w", err)
	}

	row = c.db.QueryRow(query3, id)
	if err := row.Scan(&stats.Completed); err != nil {
		return stats, fmt.Errorf("repo: MyStats: %w", err)
	}
	return stats, nil
}
