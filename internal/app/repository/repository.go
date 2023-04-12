package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	AuthRepository
	Trainer
	Client
	Comment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthRepository: NewAuthRepository(db),
		Trainer:        NewTrainerRepo(db),
		Client:         NewClientRepo(db),
		Comment:        NewCommentRepo(db),
	}
}
