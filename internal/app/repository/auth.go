package repository

import (
	"fmt"
	"html/template"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/muratovdias/diplom/internal/models"
)

type AuthRepository interface {
	CreateUser(user *models.User, img string) error
	Authentication(string) (models.User, error)
	UpdateToken(email, token string, token_duration time.Time) error
	GetUserByToken(string) (models.User, error)
}

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (s *AuthRepo) CreateUser(user *models.User, img string) error {
	row := s.db.QueryRow(`INSERT INTO users 
	(full_name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING user_id`,
		user.FullName, user.Email, user.Password, user.Role)
	var id int
	if err := row.Scan(&id); err != nil {
		fmt.Println(err.Error())
		return err
	}
	if user.Role == "trainer" {
		_, err := s.db.Exec(`INSERT INTO trainer (user_id, image) VALUES ($1, $2)`, id, template.URL(img))
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		// fmt.Println("trainer id ", id, template.URL(img))
	}
	fmt.Println("user created")
	return nil
}

func (s *AuthRepo) Authentication(email string) (models.User, error) {
	var user models.User
	err := s.db.Get(&user, "SELECT role, password FROM users WHERE email=$1", email)
	if err != nil {
		fmt.Println("user", err.Error())
		return user, fmt.Errorf("repo: Authentication: %w", err)
	}
	return user, nil
}

func (r *AuthRepo) UpdateToken(email, token string, token_duration time.Time) error {
	query := `UPDATE users SET token=$1, token_duration=$2 WHERE email=$3`
	_, err := r.db.Exec(query, token, token_duration, email)
	if err != nil {
		return fmt.Errorf("repo: UpdateToken: %w", err)
	}
	return nil
}

func (r *AuthRepo) GetUserByToken(token string) (models.User, error) {
	var user models.User
	query := "SELECT user_id, full_name, email, role, token_duration FROM users WHERE token=$1"
	row := r.db.QueryRow(query, token)
	err := row.Scan(&user.ID, &user.FullName, &user.Email, &user.Role, &user.TokenDuration)
	if err != nil {
		return models.User{}, fmt.Errorf("repo: get user by token %w", err)
	}
	// fmt.Println("by token: ", user)
	return user, nil
}
