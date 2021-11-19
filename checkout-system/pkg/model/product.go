package model

type Product struct {
	SKU           string
	Name          string
	Price         int64
	DiscountPrice int64
}

func NewProduct(sku, name string, price int64) Product {
	return Product{
		SKU:           sku,
		Name:          name,
		Price:         price,
		DiscountPrice: price,
	}
}
