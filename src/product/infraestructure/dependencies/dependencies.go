package dependencies

import (
	"fmt"
	"practice/src/helpers"
	"practice/src/product/application"
	"practice/src/product/infraestructure"
	"practice/src/product/infraestructure/controller"
)

var (
	mySQL infraestructure.MySQL
)


func Init() {
	db, err := helpers.ConnectDB()

	if err != nil {
		fmt.Println("Server error")
		return
	}

	mySQL =*infraestructure.NewMySQL(db)
}

func GetCreateProductController () *controller.CreateProductController{
	caseCreateProduct := application.NewCreateProduct(&mySQL)
	return controller.NewCreateProductController(caseCreateProduct)

}
func GetDeleteProduct ()*controller.DeleteProductController{
	caseDeleteProduct := application.NewDeleteProduct(&mySQL)
	return controller.NewDeleteProductController(caseDeleteProduct)
}
func GetCaseViewAllProducts () *controller.ViewAllProductsController {
	caseViewAllProducts := application.NewViewAllProducts(&mySQL)
	return controller.NewViewProductsController(caseViewAllProducts)
}