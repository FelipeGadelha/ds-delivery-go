package main

import (
	"github.com/FelipeGadelha/ds-delivery-go/src/api/router"
	"github.com/FelipeGadelha/ds-delivery-go/src/config"
)

func init() {
	config.LoadEnvVariables()
}

const ContextPath = "/api/ds-delivery-go/"

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {
	app := router.SetupRouter()
	app.Run()
}
