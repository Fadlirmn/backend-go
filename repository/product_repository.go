package repository

import (
	"backend-api-belajar/model"
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type ProductRepository interface {
	FindAllProduct() []model.Product
	SaveProduct(product model.Product)
	UpdateProduct(id int, product model.Product) error
	DeleteProduct(id int) error
}

type productRepo struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepo{db: db}
}

func (r *productRepo) FindAllProduct() []model.Product {
	rows, err := r.db.Query("SELECT id, name_product, item, type FROM products")

	if err != nil {
		log.Println("error query", err)
		return nil
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var u model.Product
		rows.Scan(&u.ProductId, &u.NameProduct, &u.Item, &u.Type)
		if err != nil {
			log.Println("error Scan:", err)
			continue
		}
		products = append(products, u)
	}
	return products
}

func (r *productRepo) SaveProduct(product model.Product) {
	_, err := r.db.Exec("INSERT INTO products(name_product, item, type) VALUES($1,$2,$3)", product.NameProduct, product.Item, product.Type)
	if err != nil {
		log.Println("gagal Menambahkan Product", err)
	}
}

func (r *productRepo) UpdateProduct(id int, product model.Product) error {
	_, err := r.db.Exec("UPDATE products SET name_product=$1, item=$2, type=$3 WHERE id=$4", product.NameProduct, product.Item, product.Type, id)
	return err
}

func (r *productRepo) DeleteProduct(id int) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id=$1", id)
	return err
}
