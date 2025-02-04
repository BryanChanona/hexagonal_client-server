package application

import "practice/src/product/domain"

type CreateProduct struct {
	db domain.Iproduct
}

func NewCreateProduct(db domain.Iproduct) *CreateProduct {
	return &CreateProduct{db: db}

}

func (uc *CreateProduct) Execute(product domain.Product) (err error) {

	return uc.db.SaveProduct(product)
}
