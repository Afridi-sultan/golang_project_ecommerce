package product

import "eccomerce/domain"

type service struct {
	ProductRepo ProductRepo
}

func NewService(productRepo ProductRepo) Service {
	return &service{
		ProductRepo: productRepo,
	}
}

func (s *service) Create(p domain.Products) (*domain.Products, error) {
	return s.ProductRepo.Create(p)
}


func (s *service) Get(productId int) (*domain.Products, error) {
	return s.ProductRepo.Get(productId)
}	

func (s *service) List() ([]*domain.Products, error) {
	return s.ProductRepo.List()
}	

func (s *service) Delete(id int) error {
	return s.ProductRepo.Delete(id)
}

func (s *service) Update(p domain.Products) (*domain.Products, error) {
	return s.ProductRepo.Update(p)
}