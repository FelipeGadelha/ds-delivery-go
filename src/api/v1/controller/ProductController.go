package controller

import (
	"net/http"

	"github.com/FelipeGadelha/ds-delivery-go/src/api/v1/request"
	"github.com/FelipeGadelha/ds-delivery-go/src/domain/service"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{
		ProductService: productService,
	}
}

func (pc *ProductController) FindAll(c *gin.Context) {
	products, err := pc.ProductService.FindAll()

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) FindById(c *gin.Context) {
	id := c.Param("id")
	product, err := pc.ProductService.FindById(id)
	errorHandler(err, c)
	c.JSON(http.StatusOK, product)
}

func (pc *ProductController) Create(c *gin.Context) {
	var request request.ProductRq
	var err error
	err = c.Bind(&request)
	errorHandler(&err, c)
	var product = request.ToDomain()
	created, e := pc.ProductService.Create(product)
	errorHandler(e, c)
	c.JSON(http.StatusOK, created)
}

func (pc *ProductController) Remove(c *gin.Context) {
	id := c.Param("id")
	pc.ProductService.Remove(id)
}

func errorHandler(err *error, c *gin.Context) {
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
}
