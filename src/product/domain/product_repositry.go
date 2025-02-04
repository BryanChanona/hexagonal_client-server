package domain

type Ibook interface {
	SaveProduct(book Product) (err error)
}
