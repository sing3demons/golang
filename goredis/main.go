package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/sing3demons/goredis/repositories"
	"github.com/sing3demons/goredis/services"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() *gorm.DB {
	dsn := "root:P@ssw0rd@tcp(127.0.0.1:3306)/goredismariadb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func main() {
	db := initDatabase()
	redis := initRedis()

	// productRepo := repositories.NewProductRepositoryRedis(db, redis)
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewCatalogServiceRedis(productRepo, redis)
	// productService := services.NewCatalogService(productRepo)

	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		products, err := productService.GetProducts()
		if err != nil {
			c.JSON(err.Error())
		}
		return c.JSON(products)
	})

	app.Listen(":8080")
}
