package router

import "github.com/gin-gonic/gin"

type Routes struct {
	Router *gin.Engine
}

const ContextPath = "/api/ds-delivery-go"

func (r *Routes) loadRoutes() {
	v1 := r.Router.Group(ContextPath + "/v1")
	initRoutesGroupV1(v1)

	// v2 := r.Router.Group(ContextPath + "/v2")
	// initRoutesGroupV1(v2)
}

func initRoutesGroupV1(rg *gin.RouterGroup) {
	dependencies := loadDependencies()
	rg.GET("/products", dependencies.ProductController.FindAll)
	rg.GET("/products/:id", dependencies.ProductController.FindById)
	rg.POST("/products", dependencies.ProductController.Create)
	rg.DELETE("/products/:id", dependencies.ProductController.Remove)
}

// func initRoutesGroupV2(rg *gin.RouterGroup) {
// 	dependencies := loadDependencies()
// 	rg.GET("/products", dependencies.ProductController.FindAll)
// }
