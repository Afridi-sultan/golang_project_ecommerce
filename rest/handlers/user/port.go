package user

import "eccomerce/domain"

type Service interface {
	Create(usr *domain.User) (*domain.User, error)
	GetUser(email string, password string) (*domain.User, error)
}