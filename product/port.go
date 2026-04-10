package product

import (
	"eccomerce/domain"
	productHandler "eccomerce/rest/handlers/product"
)

type ProductRepo interface {
	Create(p domain.Products) (*domain.Products, error)
	Get(productId int) (*domain.Products, error)
	List(page, limit int64) ([]*domain.Products, error)
	Delete(id int) error
	Update(p domain.Products) (*domain.Products, error)
	Count() (int64, error)
}

type Service interface {
	productHandler.Service
}