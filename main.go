package main

import (
	"fmt"
	"redis/repositories"

	"github.com/go-redis/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	/*app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		time.Sleep(time.Millisecond * 10)
		return c.SendString("hello A")
	})
	app.Listen(":8080")*/
	db := initDatabase()
	redisClient := initRedis()
	//productRep := repositories.NewProductRepositoryRedis(db, redisClient)

	productRepo := repositories.NewProductRepositoryRedis(db, redisClient)

	products, err := productRepo.GetProducts()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(products)

}

func initDatabase() *gorm.DB {
	connStr := postgres.Open("postgres://ts:ts@localhost:5432/postgres?sslmode=disable")
	db, err := gorm.Open(connStr, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
