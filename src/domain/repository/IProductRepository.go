package repository

import "github.com/FelipeGadelha/ds-delivery-go/src/domain/model"

type IProductRepository interface {
	FindAll() *[]model.Product
	FindById(id string) (*model.Product, *error)
	Save(product model.Product) (*model.Product, *error)
	Remove(id string)
}
