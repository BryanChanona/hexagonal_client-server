package main

import (
	"practice/src/product/infraestructure/dependencies"
	"practice/src/product/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	dependencies.Init()
	r := gin.Default()
	routes.ProductRouter(r)
	r.Run()


}