package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	app := gin.Default()

	r := Routes{Router: app}
	r.loadRoutes()

	return app
}
