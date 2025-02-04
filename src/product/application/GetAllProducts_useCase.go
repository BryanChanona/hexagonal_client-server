package application

import "practice/src/product/domain"

type ViewAllProducts struct {
	db domain.Iproduct
}

func NewViewAllProducts(db domain.Iproduct) *ViewAllProducts {
	return &ViewAllProducts{db: db}
}

func (vp *ViewAllProducts) Execute() ([]domain.Product, error) {
	return vp.db.GetAll()
}
