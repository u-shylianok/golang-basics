package model

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Product struct {
	SKU      string `json:"SKU"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	discount int64
	total    int64
}

func NewProduct(sku, name string, price int64) Product {
	return Product{
		SKU:   sku,
		Name:  name,
		Price: price,
	}
}

func (p *Product) SetDiscount(newPrice int64) {
	discount := p.Price - newPrice
	if discount <= 0 || discount > p.Price {
		logrus.Info("trying to set an unreal discount - skip")
		return
	}
	if discount < p.discount {
		logrus.Info("trying to set a lower discount - skip")
		return
	}
	p.discount = discount
	p.total = newPrice
}

func (p *Product) GetTotalPrice() int64 {
	if p.discount == 0 {
		return p.Price
	}
	return p.total
}

func (p Product) String() string {
	price := dollarString(p.Price)
	total := dollarString(p.GetTotalPrice())
	if p.discount == 0 {
		return fmt.Sprintf("\nProduct: %s\nSKU: [%s]\nPrice: %s\nTotal: %s", p.Name, p.SKU, price, total)
	}

	discount := dollarString(p.discount)
	return fmt.Sprintf("\nProduct: %s\nSKU: %s\nPrice: %s\nDiscount: -%s\nTotal: %s", p.Name, p.SKU, price, discount, total)
}

func dollarString(price int64) string {
	return fmt.Sprintf("%.2f$", float64(price)/100)
}
