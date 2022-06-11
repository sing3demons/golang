package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sing3demons/goredis/repositories"
)

type catalogServiceRedis struct {
	productRepo repositories.ProductRepository
	redisClient *redis.Client
}

func NewCatalogServiceRedis(productRepo repositories.ProductRepository, redisClient *redis.Client) CatalogService {
	return &catalogServiceRedis{
		productRepo: productRepo, redisClient: redisClient,
	}
}

func (s *catalogServiceRedis) GetProducts() (products []product, err error) {
	key := "service::GetProducts"
	expiration := time.Second * 10

	if productJson, err := s.redisClient.Get(context.Background(), key).Result(); err == nil {
		if json.Unmarshal([]byte(productJson), &products) == nil {
			fmt.Println("redis")
			return products, nil
		}
	}

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

	if value, err := json.Marshal(products); err == nil {
		s.redisClient.Set(context.Background(), key, string(value), expiration)
	}

	fmt.Println("database")
	return products, nil
}
