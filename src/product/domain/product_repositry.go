package domain

type Iproduct interface {
	SaveProduct(book Product) (err error)
	GetAll()([]Product,error)
	DeleteProduct(id int32) error
}
