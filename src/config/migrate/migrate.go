package main

import (
	"github.com/FelipeGadelha/ds-delivery-go/src/config"
	"github.com/FelipeGadelha/ds-delivery-go/src/domain/model"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(&model.Product{})
}
