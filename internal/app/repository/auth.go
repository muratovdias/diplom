package repository

type AuthRepository interface {
	CreateClient()
	CreateTrainer()
}

type AuthRepo struct {
	// db *gorm.DB
}

func NewAuthRepository() *AuthRepo {
	return &AuthRepo{}
}

func (s *AuthRepo) CreateClient()  {}
func (s *AuthRepo) CreateTrainer() {}
