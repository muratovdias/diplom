package service

type AuthService interface {
	CreateClient()
	CreateTrainer()
}

type Auth struct {
}

func NewAuthService() *Auth {
	return &Auth{}
}

func (s *Auth) CreateClient()  {}
func (s *Auth) CreateTrainer() {}
