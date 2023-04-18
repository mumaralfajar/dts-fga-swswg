package service

import (
	"dts-fga-swswg/product-middleware-ut/entity"
	"dts-fga-swswg/product-middleware-ut/pkg/errs"
	"dts-fga-swswg/product-middleware-ut/repository/product_repository"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProductService_GetProductById_Success(t *testing.T) {
	productRepo := product_repository.NewProductRepoMock()

	productService := NewProductService(productRepo)

	currentTime := time.Now()

	product := entity.Product{
		Id:        1,
		Title:     "Test Product",
		ImageUrl:  "https://test-product.png",
		Price:     2000,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	product_repository.GetProductById = func(productId int) (*entity.Product, errs.MessageErr) {
		return &product, nil
	}

	response, err := productService.GetProductById(1)

	assert.Nil(t, err)

	assert.NotNil(t, response)

	assert.Equal(t, "Test Product", response.Title)

	assert.Equal(t, 1, response.Id)
}

func TestProductService_GetProductById_NotFoundError(t *testing.T) {
	productRepo := product_repository.NewProductRepoMock()

	productService := NewProductService(productRepo)

	product_repository.GetProductById = func(productId int) (*entity.Product, errs.MessageErr) {
		return nil, errs.NewNotFoundError("product data not found")
	}

	response, err := productService.GetProductById(1)

	assert.Nil(t, response)

	assert.NotNil(t, err)

	assert.Equal(t, http.StatusNotFound, err.Status())

	assert.Equal(t, "product data not found", err.Message())

	assert.Equal(t, "NOT_FOUND", err.Error())
}

func TestProductService_GetProducts_Success(t *testing.T) {
	productRepo := product_repository.NewProductRepoMock()

	productService := NewProductService(productRepo)

	currentTime := time.Now()

	products := []*entity.Product{
		{
			Id:        1,
			Title:     "Test Product",
			ImageUrl:  "http://test-product.png",
			Price:     3000,
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		},
	}

	product_repository.GetProducts = func() ([]*entity.Product, errs.MessageErr) {
		return products, nil
	}

	response, err := productService.GetProducts()

	assert.Nil(t, err)

	assert.NotNil(t, response)

	assert.Equal(t, 1, len(response.Data))

	assert.Equal(t, "Test Product", response.Data[0].Title)
}

func TestProductService_GetProducts_NotFound(t *testing.T) {
	productRepo := product_repository.NewProductRepoMock()

	productService := NewProductService(productRepo)

	product_repository.GetProducts = func() ([]*entity.Product, errs.MessageErr) {
		return []*entity.Product{}, nil
	}

	response, err := productService.GetProducts()

	assert.Nil(t, err)

	assert.NotNil(t, response)

	assert.Equal(t, 0, len(response.Data))

}
