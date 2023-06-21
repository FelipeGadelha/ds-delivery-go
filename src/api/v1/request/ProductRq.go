package request

import model "github.com/FelipeGadelha/ds-delivery-go/src/domain/model"

type ProductRq struct {
	Name        string  `json:"name" binding:"required"`
	Price       float32 `json:"price" binding:"required,min=1"`
	Description string  `json:"description" binding:"required"`
	ImageUri    string  `json:"imageUri" binding:"required"`
}

func (productRq *ProductRq) ToDomain() model.Product {
	return model.Product{
		Name:        productRq.Name,
		Price:       productRq.Price,
		Description: productRq.Description,
		ImageUri:    productRq.ImageUri,
	}
}
