package product

import "eccomerce/domain"

type Service interface {
	Create(p domain.Products) (*domain.Products, error)
	Get(productId int) (*domain.Products, error)
	List() ([]*domain.Products, error)
	Delete(id int) error
	Update(p domain.Products) (*domain.Products, error)
}
