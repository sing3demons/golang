package services

import "github.com/sing3demons/goredis/repositories"

type catalogService struct {
	productRepo repositories.ProductRepository
}

func NewCatalogService(productRepo repositories.ProductRepository) CatalogService {
	return &catalogService{productRepo}
}

func (s *catalogService) GetProducts() (products []product, err error) {
	productDB, err := s.productRepo.GetProducts()

	if err != nil {
		return nil, err
	}
	
	for _, p := range productDB {
		products = append(products, product{
			ID:       p.ID,
			Name:     p.Name,
			Quantity: p.Quantity,
		})
	}
	return products, nil
}
