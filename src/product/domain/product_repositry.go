package domain

type Iproduct interface {
	SaveProduct(book Product) (err error)
}
