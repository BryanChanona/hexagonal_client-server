package controller

import (
	"net/http"
	"practice/src/product/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	deleteProduct *application.DeleteProduct
}

func NewDeleteProductController(uc *application.DeleteProduct) *DeleteProductController {
	return &DeleteProductController{deleteProduct: uc}
}

func (c *DeleteProductController) DeleteProduct(ctx *gin.Context)  {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = c.deleteProduct.Execute(int32(id)) // Elimina `productDelete` si `Execute` solo devuelve error.
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Producto eliminado correctamente"})
}
