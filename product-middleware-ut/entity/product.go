package entity

import (
	"dts-fga-swswg/product-middleware-ut/dto"
	"time"
)

type Product struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	ImageUrl  string    `json:"image_url"`
	Price     int       `json:"price"`
	UserId    int       `json:"userId"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *Product) EntityToProductResponseDto() dto.ProductResponse {
	return dto.ProductResponse{
		Id:        m.Id,
		Title:     m.Title,
		ImageUrl:  m.ImageUrl,
		Price:     m.Price,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
