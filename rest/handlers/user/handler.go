package user

import (
	
	"eccomerce/repo"
)

type Handler struct {
	userRepo repo.UserInterface
	
}

func NewHandler(userRepo repo.UserInterface) *Handler {
	return &Handler{
		userRepo: userRepo,
	}
}
