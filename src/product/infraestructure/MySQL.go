package infraestructure

import (
	"database/sql"
	"fmt"
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
func (mysql *MySQL) DeleteProduct(id int32) (error){
	

	query := "DELETE FROM product WHERE idproduct = ?"
	//Exec para delete
	result, err := mysql.db.Exec(query, id)
	if err != nil {
		return err
	}

	// Verificar si realmente se eliminó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró el producto con ID %d", id)
	}
	fmt.Println("Producto eliminado correctamente")
	return nil


 }
 func (mysql *MySQL) GetAll() ([]domain.Product, error) {
	fmt.Println("Lista de productos")
	data, err := mysql.db.Query("SELECT * FROM products")

	if err != nil {
		return nil, err
	}
	defer data.Close()

	var products []domain.Product
	// Itera sobre todas las filas devueltas por la consulta
	for data.Next() {
		var product domain.Product
		err := data.Scan(&product.ID,&product.Name, &product.Quantity,&product.BarCode)
		if err != nil {
			return nil, err
		}
		products = append(products, product)

	}
	if err := data.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
