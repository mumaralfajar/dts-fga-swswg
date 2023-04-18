package product_pg

import (
	"database/sql"
	"dts-fga-swswg/product-middleware-ut/entity"
	"dts-fga-swswg/product-middleware-ut/pkg/errs"
	"dts-fga-swswg/product-middleware-ut/repository/product_repository"
	"errors"
	"fmt"
)

const (
	getProductByIdQuery = `
		SELECT id, title, userId, imageUrl, price, createdAt, updatedAt from "products"
		WHERE id = $1;
	`

	updateProductByIdQuery = `
		UPDATE "products"
		SET title = $2,
		imageUrl = $3,
		price = $4
		WHERE id = $1;
	`
)

type productPG struct {
	db *sql.DB
}

func NewProductPG(db *sql.DB) product_repository.ProductRepository {
	return &productPG{
		db: db,
	}
}

func (m *productPG) GetProducts() ([]*entity.Product, errs.MessageErr) {
	return nil, nil
}

func (m *productPG) UpdateProductById(payload entity.Product) errs.MessageErr {
	_, err := m.db.Exec(updateProductByIdQuery, payload.Id, payload.Title, payload.ImageUrl, payload.Price)

	if err != nil {

		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (m *productPG) GetProductById(productId int) (*entity.Product, errs.MessageErr) {
	row := m.db.QueryRow(getProductByIdQuery, productId)

	var product entity.Product

	err := row.Scan(&product.Id, &product.Title, &product.UserId, &product.ImageUrl, &product.Price, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("product not found")
		}

		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &product, nil
}

func (m *productPG) CreateProduct(productPayload *entity.Product) (*entity.Product, errs.MessageErr) {
	createProductQuery := `
		INSERT INTO "product"
		(
			title,
			imageUrl,
			price,
			userId
		)
		VALUES($1, $2, $3, $4)
		RETURNING id,title, imageUrl, price, userId;
	`
	row := m.db.QueryRow(createProductQuery, productPayload.Title, productPayload.ImageUrl, productPayload.Price, productPayload.UserId)

	var product entity.Product

	err := row.Scan(&product.Id, &product.Title, &product.ImageUrl, &product.Price, &product.UserId)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &product, nil

}
