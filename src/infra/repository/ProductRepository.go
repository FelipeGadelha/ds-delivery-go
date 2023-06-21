package repository

import (
	"fmt"

	"github.com/FelipeGadelha/ds-delivery-go/src/domain/model"
	"github.com/FelipeGadelha/ds-delivery-go/src/domain/repository"
	"gorm.io/gorm"
)

type ProductRepository struct {
	database *gorm.DB
}

func NewProductRepository(database *gorm.DB) repository.IProductRepository {
	return &ProductRepository{database: database}
}

func (pr *ProductRepository) FindAll() *[]model.Product {
	fmt.Println("Testando -----------------")
	var products []model.Product
	pr.database.Find(&products)
	fmt.Println("Testando 2 -----------------")
	return &products
}

func (pr *ProductRepository) FindById(id string) (*model.Product, *error) {
	var product model.Product
	result := pr.database.Find(&product, id)

	if result.Error != nil {
		return nil, &result.Error
	}
	// user.ID             // returns inserted data's primary key
	// result.Error        // returns error
	// result.RowsAffected // returns inserted records count
	return &product, nil
}

func (pr *ProductRepository) Save(product model.Product) (*model.Product, *error) {
	result := pr.database.Create(&product)

	if result.Error != nil {
		return nil, &result.Error
	}
	return &product, nil
}

func (pr *ProductRepository) Remove(id string) {
	pr.database.Delete(&model.Product{}, id) // soft delete

	// result.Error != nil {
	// 	return &result.Error
	// }
}
