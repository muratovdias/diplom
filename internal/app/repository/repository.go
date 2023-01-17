package repository

type Repository struct {
	AuthRepository
}

func NewRepository() *Repository {
	return &Repository{
		AuthRepository: NewAuthRepository(),
	}
}
