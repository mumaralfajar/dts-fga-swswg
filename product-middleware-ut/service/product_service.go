package service

import (
	"dts-fga-swswg/product-middleware-ut/dto"
	"dts-fga-swswg/product-middleware-ut/entity"
	"dts-fga-swswg/product-middleware-ut/pkg/errs"
	"dts-fga-swswg/product-middleware-ut/pkg/helpers"
	"dts-fga-swswg/product-middleware-ut/repository/product_repository"
	"net/http"
)

type ProductService interface {
	CreateProduct(userId int, payload dto.NewProductRequest) (*dto.NewProductResponse, errs.MessageErr)
	UpdateProductById(productId int, productRequest dto.NewProductRequest) (*dto.NewProductResponse, errs.MessageErr)
	GetProductById(productId int) (*dto.ProductResponse, errs.MessageErr)
	GetProducts() (*dto.GetProductsResponse, errs.MessageErr)
}

type productService struct {
	productRepo product_repository.ProductRepository
}

func NewProductService(productRepo product_repository.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (m *productService) GetProducts() (*dto.GetProductsResponse, errs.MessageErr) {
	products, err := m.productRepo.GetProducts()

	if err != nil {
		return nil, err
	}

	productResponse := []dto.ProductResponse{}

	for _, eachProduct := range products {
		productResponse = append(productResponse, eachProduct.EntityToProductResponseDto())
	}

	response := dto.GetProductsResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "product data have been sent successfully",
		Data:       productResponse,
	}

	return &response, nil
}

func (m *productService) GetProductById(productId int) (*dto.ProductResponse, errs.MessageErr) {
	result, err := m.productRepo.GetProductById(productId)

	if err != nil {
		return nil, err
	}

	response := result.EntityToProductResponseDto()

	return &response, nil
}

func (m *productService) UpdateProductById(productId int, productRequest dto.NewProductRequest) (*dto.NewProductResponse, errs.MessageErr) {

	err := helpers.ValidateStruct(productRequest)

	if err != nil {
		return nil, err
	}

	payload := entity.Product{
		Id:       productId,
		Title:    productRequest.Title,
		ImageUrl: productRequest.ImageUrl,
		Price:    productRequest.Price,
	}

	err = m.productRepo.UpdateProductById(payload)

	if err != nil {
		return nil, err
	}

	response := dto.NewProductResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "product data successfully updated",
	}

	return &response, nil
}

func (m *productService) CreateProduct(userId int, payload dto.NewProductRequest) (*dto.NewProductResponse, errs.MessageErr) {
	productRequest := &entity.Product{
		Title:    payload.Title,
		Price:    payload.Price,
		ImageUrl: payload.ImageUrl,
		UserId:   userId,
	}

	_, err := m.productRepo.CreateProduct(productRequest)

	if err != nil {
		return nil, err
	}

	response := dto.NewProductResponse{
		StatusCode: http.StatusCreated,
		Result:     "success",
		Message:    "new product data successfully created",
	}

	return &response, err
}
