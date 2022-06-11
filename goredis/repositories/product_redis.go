package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type productRepositoryRedis struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewProductRepositoryRedis(db *gorm.DB, redisClient *redis.Client) ProductRepository {
	db.AutoMigrate(&product{})
	mackData(db)
	return &productRepositoryRedis{db: db, redisClient: redisClient}
}

func (r *productRepositoryRedis) GetProducts() (products []product, err error) {
	key := "repository::products"
	expiration := time.Second * 10

	productJson, err := r.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(productJson), &products)
		if err == nil {
			fmt.Println("redis")
			return products, nil
		}
	}

	if err := r.db.Order("quantity desc").Limit(30).Find(&products).Error; err != nil {
		return nil, err
	}

	value, err := json.Marshal(products)
	if err != nil {
		return nil, err
	}

	err = r.redisClient.Set(context.Background(), key, string(value), expiration).Err()
	if err != nil {
		return nil, err
	}

	fmt.Println("database")
	return products, nil
}
