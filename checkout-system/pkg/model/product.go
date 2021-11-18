package model

type Product struct {
	SKU   string
	Name  string
	Price int64
}

func (p Product) ToDiscountProduct(newPrice int64) DiscountProduct {
	return DiscountProduct{
		Product: p,
		Price:   newPrice,
	}
}

type DiscountProduct struct {
	Product
	Price int64
}
