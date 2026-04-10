package product

import "eccomerce/domain"

type Service interface {
	Create(p domain.Products) (*domain.Products, error)
	Get(productId int) (*domain.Products, error)
	List(page, limit int64) ([]*domain.Products, error)
	Count()(int64, error)
	Delete(id int) error
	Update(p domain.Products) (*domain.Products, error)
}
