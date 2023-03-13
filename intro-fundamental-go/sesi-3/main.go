package main

import "fmt"

type MyString string

func (m MyString) PriceFormatter() string {
	return fmt.Sprintf("Rp %s", m)
}

type OrderDetail struct {
	Id         string
	Amount     int
	TotalPrice int
}

func (p Product) PriceFormatter() string {
	return fmt.Sprintf("Rp %d", p.Price)
}

type Product struct {
	Id          string
	Title       string
	Price       int
	Orders      []OrderDetail
	formatPrice func(int) string
}

func main() {
	var p1 Product = Product{
		Id:    "prod-1",
		Title: "Jas Hujan Kece",
		Price: 3000,
		formatPrice: func(price int) string {
			return fmt.Sprintf("Rp %d", price)
		},
	}

	var p1Price = p1.PriceFormatter()

	fmt.Println("p1Price:", p1Price)

	var orders []OrderDetail = []OrderDetail{
		{
			Id:         "o-1",
			Amount:     2,
			TotalPrice: 6000,
		},
		{
			Id:         "o-2",
			Amount:     3,
			TotalPrice: 9000,
		},
	}

	p1.Orders = append(p1.Orders, orders...)

	// for _, eachOrder := range orders {
	// 	p1.Orders = append(p1.Orders, eachOrder)
	// }

	// fmt.Printf("%+v \n", p1)

	// fmt.Println("p1 price:", p1)

	// fmt.Printf("p1:%+v \n", p1)

	var p2 Product

	p2.Id = "prod-2"

	p2.Title = "Plastik Merah"

	p2.Price = 100

	// fmt.Printf("p2:%+v \n", p2)

	var products []Product = []Product{
		{
			Id:    "prod-1",
			Title: "Jas Hujan Kece",
			Price: 3000,
		},
		{
			Id:    "prod-2",
			Title: "Kemeja",
			Price: 9000,
		},
	}

	_ = products

	// for _, eachProduct := range products {
	// 	fmt.Println("id:", eachProduct.Id)
	// 	fmt.Println("title:", eachProduct.Title)
	// 	fmt.Println("price:", eachProduct.Price)
	// }

	var productDetail map[string]Product = map[string]Product{
		"product-1": Product{
			Id:    "p-1",
			Title: "Jaket Supreme",
			Price: 2000000,
		},
		"product-2": Product{
			Id:    "p-2",
			Title: "Sweat Pants Tommy Hilfiger",
			Price: 6000000,
		},
	}

	_ = productDetail

	// for key, value := range productDetail {
	// 	fmt.Printf("%s: %+v \n", key, value)
	// }

}
