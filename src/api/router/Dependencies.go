package router

import (
	"github.com/FelipeGadelha/ds-delivery-go/src/api/v1/controller"
	"github.com/FelipeGadelha/ds-delivery-go/src/config"
	"github.com/FelipeGadelha/ds-delivery-go/src/domain/service"
	"github.com/FelipeGadelha/ds-delivery-go/src/infra/repository"
)

type Dependencies struct {
	ProductController controller.ProductController
}

func loadDependencies() *Dependencies {
	// ctx := context.Background()
	database := config.DB
	productRepository := repository.NewProductRepository(database)
	productService := service.NewProductService(productRepository)
	return &Dependencies{
		ProductController: *controller.NewProductController(productService),
	}
}
