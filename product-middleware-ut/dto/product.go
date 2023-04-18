package dto

import "time"

type NewProductRequest struct {
	Title    string `json:"product" valid:"required~product title cannot be empty" example:"Buku"`
	ImageUrl string `json:"imageUrl" valid:"required~image url cannot be empty" example:"http://imageurl.com"`
	Price    int    `json:"price" valid:"required~price cannot be empty" example:"20000"`
}

type NewProductResponse struct {
	Result     string `json:"result"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

type ProductResponse struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	ImageUrl  string    `json:"imageUrl"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GetProductsResponse struct {
	Result     string            `json:"result"`
	Message    string            `json:"message"`
	StatusCode int               `json:"statusCode"`
	Data       []ProductResponse `json:"data"`
}
