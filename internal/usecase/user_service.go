package usecase

import "zura.org/oauth2-api/internal/domain"

type UserService struct {
	Repo domain.UserRepository
}

func (s *UserService) GetOrCreateUser(email domain.Email, name domain.Name) (*domain.User, error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, err
	}
	newUser := domain.User{Email: email, Name: name}
	err = s.Repo.Create(newUser)
	return newUser, err
}
