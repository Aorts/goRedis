package repositories

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type productRepositoryDB struct {
	db *gorm.DB
}

func NewProductRepositoryDB(db *gorm.DB) ProductRepository {
	db.AutoMigrate(&product{})
	mockData(db)
	return productRepositoryDB{db}
}

func mockData(db *gorm.DB) error {

	var count int64
	db.Model(&product{}).Count(&count)

	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)

	products := []product{}
	for i := 0; i < 5000; i++ {
		products = append(products, product{
			Name:     fmt.Sprintf("ProductNo%v", i+1),
			Quantity: rand.Intn(100),
		})
	}
	return db.Create(&products).Error
}

func (r productRepositoryDB) GetProducts() (products []product, err error) {
	err = r.db.Order("quantity desc").Limit(30).Find(&products).Error
	return products, err
}
