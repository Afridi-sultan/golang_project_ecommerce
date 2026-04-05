package user
import (
	"eccomerce/domain"
	userHandler "eccomerce/rest/handlers/user"

)

type Service interface {
	userHandler.Service
}

type UserInterface interface {
	Create(usr *domain.User) (*domain.User, error)
	GetUser(email, pass string) (*domain.User, error)
	// List() ([]*domain.User, error)
	// Delete(id int) error
	// Update(usr domain.User) (*domain.User, error)
}