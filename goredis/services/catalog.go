package services


type product struct {
	ID       uint
	Name     string
	Quantity int
}


type CatalogService interface {
	GetProducts() ([]product, error)
}
