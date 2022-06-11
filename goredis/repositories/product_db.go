package repositories

import (
	"gorm.io/gorm"
)

type productRepositoryDB struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	db.AutoMigrate(&product{})
	mackData(db)
	return &productRepositoryDB{db}
}

func (r *productRepositoryDB) GetProducts() (products []product, err error) {
	err = r.db.Order("quantity desc").Limit(30).Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}
