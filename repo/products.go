package repo

import (
	"database/sql"
	"eccomerce/domain"
	"eccomerce/product"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db *sqlx.DB
}

// constructor
func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}

}

// Methods
func (r *productRepo) Create(p domain.Products) (*domain.Products, error) {
	query := `
	INSERT INTO products (title, description, price, img_url)
	VALUES ($1, $2, $3, $4)
	RETURNING id;
	`
	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl)
	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *productRepo) Get(id int) (*domain.Products, error) {
	var prdt domain.Products

	query := `
	SELECT id, title, description, price, img_url
	FROM products
	WHERE id = $1;
	`
	err := r.db.Get(&prdt, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &prdt, nil
}

func (r *productRepo) List() ([]*domain.Products, error) {
	var productList []*domain.Products

	query := `
	SELECT id, title, description, price, img_url
	FROM products
	
	`
	err := r.db.Select(&productList, query)
	if err != nil {
		return nil, err
	}

	return productList, nil
}

func (r *productRepo) Delete(id int) error {
	query := `
	DELETE FROM products WHERE id = $1
	`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err

	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("product not found")
	}

	return nil
}

func (r *productRepo) Update(p domain.Products) (*domain.Products, error) {
	query := `
	UPDATE products
	SET 
		title = $1,
		description = $2,
		price = $3,
		img_url = $4
	WHERE id = $5;
	`

	result, err := r.db.Exec(query, p.Title, p.Description, p.Price, p.ImgUrl, p.ID)
	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, fmt.Errorf("product not found")
	}

	return &p, nil
}
