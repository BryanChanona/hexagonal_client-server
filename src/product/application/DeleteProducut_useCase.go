package application

import "practice/src/product/domain"

type DeleteProduct struct {
	db domain.Iproduct
}

func NewDeleteProduct(db domain.Iproduct) *DeleteProduct {
	return &DeleteProduct{db: db}
}

/**/
func (uc *DeleteProduct) Execute(id int32) error {
	return uc.db.DeleteProduct(id)
}
