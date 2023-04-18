package product_repository

import (
	"dts-fga-swswg/product-middleware-ut/entity"
	"dts-fga-swswg/product-middleware-ut/pkg/errs"
)

type ProductRepository interface {
	CreateProduct(productPayload *entity.Product) (*entity.Product, errs.MessageErr)
	GetProductById(productId int) (*entity.Product, errs.MessageErr)
	UpdateProductById(payload entity.Product) errs.MessageErr
	GetProducts() ([]*entity.Product, errs.MessageErr)
}
