package routes

import (
	_ "net/http"
	"practice/src/product/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)


func ProductRouter( router *gin.Engine){
	routes := router.Group("/products")
	createProductController := dependencies.GetCreateProductController().CreateProduct
	getDeleteProduct := dependencies.GetDeleteProduct().DeleteProduct
	

	
	routes.POST("", createProductController)
	routes.DELETE("/:id",getDeleteProduct)
}








