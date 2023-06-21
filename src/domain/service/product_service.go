package service

import (
	"github.com/FelipeGadelha/ds-delivery-go/src/domain/model"
	"github.com/FelipeGadelha/ds-delivery-go/src/domain/repository"
)

type ProductService struct {
	productRepository repository.IProductRepository
}

func NewProductService(ProductRepository repository.IProductRepository) ProductService {
	return ProductService{
		productRepository: ProductRepository,
	}
}

func (ps *ProductService) FindAll() (*[]model.Product, *error) {
	products := ps.productRepository.FindAll()
	return products, nil
}

func (ps *ProductService) FindById(id string) (*model.Product, *error) {
	result, err := ps.productRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ps *ProductService) Create(product model.Product) (*model.Product, *error) {
	result, err := ps.productRepository.Save(product)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ps *ProductService) Update(id string, update model.Product) (*model.Product, *error) {
	product, err := ps.FindById(id)
	if err != nil {
		return nil, err
	}
	product.Update(update)
	// ps.productRepository.
	return product, nil
}

func (ps *ProductService) Remove(id string) {
	ps.productRepository.Remove(id)
}
