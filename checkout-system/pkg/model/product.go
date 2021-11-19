package model

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Product struct {
	SKU      string
	Name     string
	Price    int64
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
	if p.discount == 0 {
		return fmt.Sprintf("\nProduct: [%s]\nSKU: [%s]\nPrice: [%d]\nTotal: [%d]", p.Name, p.SKU, p.Price, p.total)
	}
	return fmt.Sprintf("\nProduct: [%s]\nSKU: [%s]\nPrice: [%d]\nDiscount: [-%d]\nTotal: [%d]", p.Name, p.SKU, p.Price, p.discount, p.total)
}
