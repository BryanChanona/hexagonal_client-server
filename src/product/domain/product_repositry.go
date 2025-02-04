package domain

type Iproduct interface {
	SaveProduct(book Product) (err error)
	
	DeleteProduct(id int32) error
	
}
