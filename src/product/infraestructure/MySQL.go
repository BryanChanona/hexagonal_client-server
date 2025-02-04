package infraestructure

import (
	"database/sql"
	"practice/src/product/domain"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}


func (mysql *MySQL) SaveProduct(product domain.Product) (err error) {
	sentenciaPreparada, err := mysql.db.Prepare("INSERT INTO products (id, name, quantity,barcode) VALUES (?,?,?,?)")

	if err != nil {
		return err
	}

	defer sentenciaPreparada.Exec()

	_, err = sentenciaPreparada.Exec(product.ID, product.Name, product.Quantity,product.BarCode)

	if err != nil {
		return err
	}

	return nil

}