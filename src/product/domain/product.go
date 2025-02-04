package domain

type Product struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
	Quantity int `json:"quantity"`
	BarCode string `json:"barcode"`
}