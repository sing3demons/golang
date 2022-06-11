package services

type product struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type CatalogService interface {
	GetProducts() ([]product, error)
}
