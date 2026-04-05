package user

import "eccomerce/domain"

type service struct {
	userRepo UserInterface
}

func NewService(userRepo UserInterface) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) Create(usr *domain.User) (*domain.User, error) {
	user, err := s.userRepo.Create(usr)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return user, nil
}

func (s *service) GetUser(email, pass string) (*domain.User, error) {
	user, err := s.userRepo.GetUser(email, pass)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return user, nil
}
