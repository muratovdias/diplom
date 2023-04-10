package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	AuthRepository
	Trainer
	Client
}

func NewRepository(db *sqlx.DB, tx *sqlx.Tx) *Repository {
	return &Repository{
		AuthRepository: NewAuthRepository(db),
		Trainer:        NewTrainerRepo(db),
		Client:         NewClientRepo(db),
	}
}
